package mappers

import (
	"github.com/isOdin-l/TinderArt/pkg/db_models"
	"github.com/isOdin-l/TinderArt/services/profile/internal/entities"
	"github.com/jackc/pgx/v5/pgtype"
)

// DB ---> Entity
func FromMatchesDBToEntity(req []pgtype.UUID) *entities.Profile {
	matcheUid := db_models.FromPgUUIDTextArray(req)
	if matcheUid == nil {
		return nil
	}

	return &entities.Profile{
		Matches: *matcheUid,
	}
}
