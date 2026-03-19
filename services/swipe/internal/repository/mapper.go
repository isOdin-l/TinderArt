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

func ToPgUUIDNullable(v *pgtype.Bool) *bool {
	if !v.Valid {
		return nil
	}

	return &v.Bool
}

func ToPgBollNullable(v *bool) pgtype.Bool {
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

func ToPgBool(v bool) pgtype.Bool {
	return pgtype.Bool{
		Bool:  v,
		Valid: true,
	}
}
