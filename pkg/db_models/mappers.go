package db_models

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func ToPgUUID(v *uuid.UUID) pgtype.UUID {
	if v == nil {
		return pgtype.UUID{
			Valid: false,
		}
	}

	return pgtype.UUID{
		Bytes: *v,
		Valid: true,
	}
}

func FromPgUUID(v *pgtype.UUID) uuid.UUID {
	return uuid.UUID(v.Bytes)
}

func ToPgText(v *string) pgtype.Text {
	if v == nil {
		return pgtype.Text{
			Valid: false,
		}
	}

	return pgtype.Text{
		String: *v,
		Valid:  true,
	}
}

func FromPgText(v *pgtype.Text) string {
	return v.String
}

func ToPgBoll(v *bool) pgtype.Bool {
	if v == nil {
		return pgtype.Bool{
			Bool:  false,
			Valid: false,
		}
	}

	return pgtype.Bool{
		Bool:  *v,
		Valid: true,
	}
}
func FromPgBool(v pgtype.Bool) *bool {
	return &v.Bool
}
