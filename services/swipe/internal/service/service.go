package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/isOdin-l/TinderArt/services/swipe/internal/entities"
)

var (
	ErrAlreadyExist = "no rows in result set"
)

type IRepository interface {
	CreateUpdateSwipe(ctx context.Context, swipe *entities.Swipe) error
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

	errDb := s.repo.CreateUpdateSwipe(ctx, swipe)
	if errDb != nil {
		return errDb
	}

	if swipe.Decision1 != nil && swipe.Decision2 != nil && *swipe.Decision1 && *swipe.Decision2 {
		if errPub := s.publishMatch(ctx, swipe); errPub != nil {
			return errPub
		}
	}
	// slog.Info(fmt.Sprintf("%s %s", swipe.Decision1, swipe.Decision2))
	return nil
}

func (s *SwipeService) publishMatch(ctx context.Context, swipe *entities.Swipe) error {
	msg := []byte(fmt.Sprintf("%s:%s", swipe.UserId, swipe.TargetId))

	return s.broker.WriteMessage(ctx, nil, msg)
}
