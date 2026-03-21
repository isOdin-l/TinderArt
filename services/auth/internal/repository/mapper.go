package repository

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func ToPgUUID(v *uuid.UUID) pgtype.UUID {
	return pgtype.UUID{
		Bytes: *v,
		Valid: true,
	}
}

func FromPgUUID(v *pgtype.UUID) uuid.UUID {
	return uuid.UUID(v.Bytes)
}

func ToPgText(v *string) pgtype.Text {
	return pgtype.Text{
		String: *v,
		Valid:  true,
	}
}

func FromPgText(v *pgtype.Text) string {
	return v.String
}
