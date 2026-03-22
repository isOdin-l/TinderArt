package entities

import (
	"github.com/jackc/pgx/v5/pgtype"
)

func FromPgTypeToInterface(matchesDb []pgtype.UUID) []interface{} {
	matches := make([]interface{}, len(matchesDb))
	for i, v := range matchesDb {
		matches[i] = v.String()
	}
	return matches
}
