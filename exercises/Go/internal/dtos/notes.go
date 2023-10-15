package dtos

import (
	"check-in/api/internal/models"
	"check-in/api/internal/validator"
)

type PaginatedNotesDto struct {
	PaginatedResultDto[models.Note]
}

type CreateNoteDto struct {
	Title    string `json:"title"`
	Contents string `json:"contents"`
}

type UpdateNoteDto struct {
	Title    *string `json:"title"`
	Contents *string `json:"contents"`
}

func ValidateCreateNoteDto(
	v *validator.Validator,
	createNoteDto CreateNoteDto,
) {
	v.Check(createNoteDto.Title != "", "title", "must be provided")
	v.Check(createNoteDto.Contents != "", "contents", "must be provided")
}

func ValidateUpdateNoteDto(
	v *validator.Validator,
	updateNoteDto UpdateNoteDto,
) {
	if updateNoteDto.Title != nil {
		v.Check(*updateNoteDto.Title != "", "title", "must be provided")
	}

	if updateNoteDto.Contents != nil {
		v.Check(*updateNoteDto.Contents != "", "contents", "must be provided")
	}
}
