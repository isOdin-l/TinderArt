package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/isOdin-l/TinderArt/services/swipe/internal/entities"
)

var (
	ErrAlreadyExist = "Already exist"
)

type IRepository interface {
	CreateSwipe(ctx context.Context, swipe *entities.Swipe) error
	UpdateSwipe(ctx context.Context, swipe *entities.Swipe) (bool, bool, error)
}

type IMsgBroket interface {
	WriteMessage(ctx context.Context, key, value []byte) error
}

type ITxManager interface {
	WithTx(ctx context.Context, fn func(ctx context.Context) (any, error)) (any, error)
}

type SwipeService struct {
	repo      IRepository
	broker    IMsgBroket
	txManager ITxManager
}

func NewService(repo IRepository, broker IMsgBroket, txManager ITxManager) *SwipeService {
	return &SwipeService{repo: repo, broker: broker, txManager: txManager}
}

func (s *SwipeService) CreateSwipe(ctx context.Context, swipe *entities.Swipe) error {
	var errUid error

	swipe.Id, errUid = uuid.NewV7()
	if errUid != nil {
		return errUid
	}

	isMatch, errTx := s.txManager.WithTx(ctx, func(ctx context.Context) (any, error) {
		errCreate := s.repo.CreateSwipe(ctx, swipe)
		if errCreate != nil && errCreate.Error() != ErrAlreadyExist {
			return false, errCreate
		}

		// If user is not first, swap user and target
		swipe.UserId, swipe.TargetId = swipe.TargetId, swipe.UserId
		swipe.Decision2 = swipe.Decision1

		dec1, dec2, errUpdate := s.repo.UpdateSwipe(ctx, swipe)
		if errUpdate != nil {
			return false, errUpdate
		}

		return dec1 && dec2, nil
	})
	if errTx != nil {
		return errTx
	}

	if isMatch.(bool) {
		go func() {
			var pubErr error
			for pubErr != nil {
				pubErr = s.publishMatch(ctx, swipe)
				time.Sleep(time.Second * 5)
			}
		}()
	}

	return nil
}

func (s *SwipeService) publishMatch(ctx context.Context, swipe *entities.Swipe) error {
	msg := entities.BrokerMatchMessage{
		User1ID: swipe.UserId,
		User2ID: swipe.TargetId,
	}
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return s.broker.WriteMessage(ctx, nil, data)
}
