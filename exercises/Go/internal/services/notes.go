package services

import (
	"context"

	"check-in/api/internal/database"
	"check-in/api/internal/dtos"
	"check-in/api/internal/models"
)

type NotesService struct {
	db database.DB
}

func (service NotesService) GetTotalCount(ctx context.Context) (*int64, error) {
	query := `
		SELECT COUNT(*)
		FROM notes
	`

	var total *int64

	err := service.db.QueryRow(ctx, query).Scan(&total)
	if err != nil {
		return nil, handleError(err)
	}

	return total, nil
}

func (service NotesService) GetAllPaginated(
	ctx context.Context,
	limit int64,
	offset int64,
) ([]*models.Note, error) {
	query := `
		SELECT id, title, contents
		FROM notes
		ORDER BY title ASC
		LIMIT $1 OFFSET $2
	`

	rows, err := service.db.Query(ctx, query, limit, offset)
	if err != nil {
		return nil, handleError(err)
	}

	notes := []*models.Note{}

	for rows.Next() {
		var note models.Note

		err = rows.Scan(
			&note.ID,
			&note.Title,
			&note.Contents,
		)

		if err != nil {
			return nil, handleError(err)
		}

		notes = append(notes, &note)
	}

	if err = rows.Err(); err != nil {
		return nil, handleError(err)
	}

	return notes, nil
}

func (service NotesService) GetAll(
	ctx context.Context,
) ([]*models.Note, error) {
	query := `
		SELECT id, title, contents
		FROM notes
	`

	rows, err := service.db.Query(ctx, query)
	if err != nil {
		return nil, handleError(err)
	}

	notes := []*models.Note{}

	for rows.Next() {
		var note models.Note

		err = rows.Scan(
			&note.ID,
			&note.Title,
			&note.Contents,
		)

		if err != nil {
			return nil, handleError(err)
		}

		notes = append(notes, &note)
	}

	if err = rows.Err(); err != nil {
		return nil, handleError(err)
	}

	return notes, nil
}

func (service NotesService) GetByID(
	ctx context.Context,
	id string,
) (*models.Note, error) {
	query := `
		SELECT title, contents
		FROM notes
		WHERE id = $1
	`

	note := models.Note{
		ID: id,
	}

	err := service.db.QueryRow(
		ctx,
		query,
		id).Scan(&note.Title, &note.Contents)

	if err != nil {
		return nil, handleError(err)
	}

	return &note, nil
}

func (service NotesService) Create(
	ctx context.Context,
	title string,
	contents string,
) (*models.Note, error) {
	query := `
		INSERT INTO notes (title, contents)
		VALUES ($1, $2)
		RETURNING id
	`

	note := models.Note{
		Title:    title,
		Contents: contents,
	}

	err := service.db.QueryRow(ctx, query, title, contents).Scan(&note.ID)

	if err != nil {
		return nil, handleError(err)
	}

	return &note, nil
}

func (service NotesService) Update(
	ctx context.Context,
	note *models.Note,
	updateNoteDto dtos.UpdateNoteDto,
) error {
	if updateNoteDto.Title != nil {
		note.Title = *updateNoteDto.Title
	}

	if updateNoteDto.Contents != nil {
		note.Contents = *updateNoteDto.Contents
	}

	query := `
		UPDATE notes
		SET title = $2, contents = $3
		WHERE id = $1
	`

	result, err := service.db.Exec(ctx, query, note.ID, note.Title, note.Contents)
	if err != nil {
		return handleError(err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}

func (service NotesService) Delete(ctx context.Context, id string) error {
	query := `
		DELETE FROM notes
		WHERE id = $1
	`

	result, err := service.db.Exec(ctx, query, id)
	if err != nil {
		return handleError(err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}

	return nil
}
