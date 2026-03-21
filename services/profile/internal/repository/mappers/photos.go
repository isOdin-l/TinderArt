package mappers

import (
	"github.com/isOdin-l/TinderArt/pkg/db_models"
	"github.com/isOdin-l/TinderArt/services/profile/internal/entities"
	"github.com/jackc/pgx/v5/pgtype"
)

// DB ---> ENTITY

// ENTITY ---> DB
func FromEntityToCreatePhotos(entity *entities.Profile) *db_models.CreatePhotosParams {
	urlsDb := make([]pgtype.Text, len(entity.PhotoUrls))
	idsDb := make([]pgtype.UUID, len(entity.PhotosIds))

	for idx := range len(entity.PhotoUrls) {
		urlsDb[idx] = db_models.ToPgText(&entity.PhotoUrls[idx])
		idsDb[idx] = db_models.ToPgUUID(&entity.PhotosIds[idx])
	}

	return &db_models.CreatePhotosParams{
		Ids:       idsDb,
		ProfileID: db_models.ToPgUUID(&entity.UserId),
		Urls:      urlsDb,
	}
}

func FromEntityToDeletePhotos(entity *entities.Profile) *db_models.DeletePhotosParams {
	return &db_models.DeletePhotosParams{}
}
