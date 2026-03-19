package repository

import (
	"context"

	"github.com/isOdin-l/TinderArt/pkg/db_models"
	"github.com/isOdin-l/TinderArt/services/swipe/internal/entities"
)

type QueryBuilder interface {
	InsertUpdateSwipe(ctx context.Context, arg db_models.InsertUpdateSwipeParams) (db_models.InsertUpdateSwipeRow, error)
}

type Repository struct {
	q QueryBuilder
}

func NewRepository(db db_models.DBTX) *Repository {
	return &Repository{q: db_models.New(db)}
}

func (r *Repository) CreateUpdateSwipe(ctx context.Context, swipe *entities.Swipe) error {
	insertModel := db_models.InsertUpdateSwipeParams{
		ID:        ToPgUUID(&swipe.Id),
		UserID1:   ToPgUUID(&swipe.UserId),
		UserID2:   ToPgUUID(&swipe.TargetId),
		Desicion1: ToPgBollNullable(swipe.Decision1),
		Desicion2: ToPgBollNullable(swipe.Decision2),
	}

	res, errDB := r.q.InsertUpdateSwipe(ctx, insertModel)
	if errDB != nil {
		return errDB
	}

	swipe.Decision1 = ToPgUUIDNullable(&res.Desicion1)
	swipe.Decision2 = ToPgUUIDNullable(&res.Desicion2)

	return nil
}
