package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"check-in/api/internal/dtos"
	"check-in/api/internal/helpers"
	"check-in/api/internal/models"
	"check-in/api/internal/validator"
)

func (app *application) notesRoutes(router *httprouter.Router) {
	router.HandlerFunc(
		http.MethodGet,
		"/notes",
		app.getPaginatedNotesHandler,
	)
	router.HandlerFunc(
		http.MethodPost,
		"/notes",
		app.createNoteHandler,
	)
	router.HandlerFunc(
		http.MethodPatch,
		"/notes/:id",
		app.updateNoteHandler,
	)
	router.HandlerFunc(
		http.MethodDelete,
		"/notes/:id",
		app.deleteNoteHandler,
	)
}

func (app *application) getPaginatedNotesHandler(w http.ResponseWriter,
	r *http.Request) {
	var pageSize int64 = 4

	page, err := helpers.ReadIntQueryParam(r, "page", 1)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	result, err := getAllPaginated[models.Note](
		r.Context(),
		app.services.Notes,
		page,
		pageSize,
	)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, result, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createNoteHandler(w http.ResponseWriter, r *http.Request) {
	var createNoteDto dtos.CreateNoteDto

	err := helpers.ReadJSON(r.Body, &createNoteDto)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	if dtos.ValidateCreateNoteDto(v, createNoteDto); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	note, err := app.services.Notes.Create(
		r.Context(),
		createNoteDto.Title,
		createNoteDto.Contents,
	)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = helpers.WriteJSON(w, http.StatusCreated, note, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) updateNoteHandler(w http.ResponseWriter, r *http.Request) {
	var updateNoteDto dtos.UpdateNoteDto

	id, err := helpers.ReadUUIDURLParam(r, "id")
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	err = helpers.ReadJSON(r.Body, &updateNoteDto)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()

	if dtos.ValidateUpdateNoteDto(v, updateNoteDto); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	note, err := app.services.Notes.GetByID(r.Context(), id)
	if err != nil {
		app.notFoundResponse(w, r, err, "note", "id", id, "id")
		return
	}

	err = app.services.Notes.Update(r.Context(), note, updateNoteDto)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, note, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteNoteHandler(w http.ResponseWriter, r *http.Request) {
	id, err := helpers.ReadUUIDURLParam(r, "id")
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	note, err := app.services.Notes.GetByID(r.Context(), id)
	if err != nil {
		app.notFoundResponse(w, r, err, "note", "id", id, "id")
		return
	}

	err = app.services.Notes.Delete(r.Context(), note.ID)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = helpers.WriteJSON(w, http.StatusOK, note, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
