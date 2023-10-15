package services

import (
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"

	"check-in/api/internal/database"
)

var (
	ErrRecordNotFound    = errors.New("record not found")
	ErrRecordUniqueValue = errors.New("record unique value already used")
)

type Services struct {
	Notes NotesService
}

func New(db database.DB) Services {
	notes := NotesService{db: db}

	return Services{
		Notes: notes,
	}
}

func handleError(err error) error {
	var pgxError *pgconn.PgError
	errors.As(err, &pgxError)

	switch {
	case errors.Is(err, pgx.ErrNoRows):
		return ErrRecordNotFound
	case pgxError.Code == "23503":
		return ErrRecordNotFound
	case pgxError.Code == "23505":
		return ErrRecordUniqueValue
	default:
		return err
	}
}
