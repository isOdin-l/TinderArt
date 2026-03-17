package repository

import (
	"context"
	"errors"

	"github.com/isOdin-l/TinderArt/services/swipe/internal/database/sqlc"
	"github.com/isOdin-l/TinderArt/services/swipe/internal/entities"
)

type Repository struct {
	q *sqlc.Queries
}

func NewRepository(db sqlc.DBTX) *Repository {
	return &Repository{q: sqlc.New(db)}
}

func (r *Repository) CreateSwipe(ctx context.Context, swipe *entities.Swipe) error {
	insertModel := sqlc.InsertSwipeParams{
		ID:        ToPgUUID(&swipe.Id),
		UserID1:   ToPgUUID(&swipe.UserId),
		UserID2:   ToPgUUID(&swipe.TargetId),
		Desicion1: ToPgBool(swipe.Decision1),
	}

	swipeId, errDB := r.q.InsertSwipe(ctx, insertModel)
	if errDB != nil {
		return errDB
	}
	if swipeId.String() != swipe.Id.String() {
		return errors.New("Already exist")
	}

	return nil
}

func (r *Repository) UpdateSwipe(ctx context.Context, swipe *entities.Swipe) (bool, bool, error) {
	updateParams := sqlc.UpdateSwipeParams{
		UserID1:   ToPgUUID(&swipe.UserId),
		UserID2:   ToPgUUID(&swipe.TargetId),
		Desicion2: ToPgBool(swipe.Decision2),
	}

	updateRow, errUpdate := r.q.UpdateSwipe(ctx, updateParams)
	if errUpdate != nil {
		return false, false, errUpdate
	}

	return updateRow.Desicion1.Bool, updateRow.Desicion2.Bool, nil
}
