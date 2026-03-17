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

func ToPgBool(v bool) pgtype.Bool {
	return pgtype.Bool{
		Bool:  v,
		Valid: true,
	}
}
