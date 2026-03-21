package mappers

import (
	"github.com/isOdin-l/TinderArt/pkg/db_models"
	"github.com/isOdin-l/TinderArt/services/profile/internal/entities"
	"github.com/jackc/pgx/v5/pgtype"
)

// DB ---> ENTITY

// ENTITY ---> DB
func FromEntityToCreateFavArtStyles(entity *entities.Profile) *db_models.CreateFavArtStyleParams {
	styleDb := make([]pgtype.Text, len(entity.FavArtStyles))
	idsDb := make([]pgtype.UUID, len(entity.FavArtStyles))

	for idx := range len(entity.FavArtStyles) {
		styleDb[idx] = db_models.ToPgText(&entity.FavArtStyles[idx])
		idsDb[idx] = db_models.ToPgUUID(&entity.FavArtStylesIds[idx])
	}

	return &db_models.CreateFavArtStyleParams{
		Ids:       idsDb,
		ProfileID: db_models.ToPgUUID(&entity.UserId),
		Styles:    styleDb,
	}
}
