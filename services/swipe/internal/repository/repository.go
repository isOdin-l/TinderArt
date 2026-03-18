package repository

import (
	"context"
	"errors"

	"github.com/isOdin-l/TinderArt/pkg/db_models"
	"github.com/isOdin-l/TinderArt/services/swipe/internal/entities"
	"github.com/jackc/pgx/v5/pgtype"
)

type QueryBuilder interface {
	InsertSwipe(ctx context.Context, arg db_models.InsertSwipeParams) (pgtype.UUID, error)
	UpdateSwipe(ctx context.Context, arg db_models.UpdateSwipeParams) (db_models.UpdateSwipeRow, error)
}

type Repository struct {
	q QueryBuilder
}

func NewRepository(db db_models.DBTX) *Repository {
	return &Repository{q: db_models.New(db)}
}

func (r *Repository) CreateSwipe(ctx context.Context, swipe *entities.Swipe) error {
	insertModel := db_models.InsertSwipeParams{
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
	updateParams := db_models.UpdateSwipeParams{
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
