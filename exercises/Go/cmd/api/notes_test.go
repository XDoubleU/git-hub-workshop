package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"

	"check-in/api/internal/assert"
	"check-in/api/internal/dtos"
	"check-in/api/internal/helpers"
	"check-in/api/internal/models"
	"check-in/api/internal/tests"
)

func TestGetPaginatedNotesDefaultPage(t *testing.T) {
	testEnv, testApp := setupTest(t, mainTestEnv)
	defer tests.TeardownSingle(testEnv)

	ts := httptest.NewTLSServer(testApp.routes())
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/notes", nil)

	rs, _ := ts.Client().Do(req)

	var rsData dtos.PaginatedNotesDto
	_ = helpers.ReadJSON(rs.Body, &rsData)

	assert.Equal(t, rs.StatusCode, http.StatusOK)

	assert.Equal(t, rsData.Pagination.Current, 1)
	assert.Equal(
		t,
		rsData.Pagination.Total,
		int64(math.Ceil(float64(fixtureData.AmountOfNotes)/4)),
	)
	assert.Equal(t, len(rsData.Data), 4)

	assert.Equal(t, rsData.Data[0].ID, fixtureData.Notes[0].ID)
	assert.Equal(t, rsData.Data[0].Title, "TestNote0")
	assert.Equal(t, rsData.Data[0].Contents, "Some text")
}

func TestGetPaginatedNotesSpecificPage(t *testing.T) {
	testEnv, testApp := setupTest(t, mainTestEnv)
	defer tests.TeardownSingle(testEnv)

	ts := httptest.NewTLSServer(testApp.routes())
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/notes?page=2", nil)

	rs, _ := ts.Client().Do(req)

	var rsData dtos.PaginatedNotesDto
	_ = helpers.ReadJSON(rs.Body, &rsData)

	assert.Equal(t, rs.StatusCode, http.StatusOK)

	assert.Equal(t, rsData.Pagination.Current, 2)
	assert.Equal(
		t,
		rsData.Pagination.Total,
		int64(math.Ceil(float64(fixtureData.AmountOfNotes)/4)),
	)
	assert.Equal(t, len(rsData.Data), 4)

	assert.Equal(t, rsData.Data[0].ID, fixtureData.Notes[12].ID)
	assert.Equal(t, rsData.Data[0].Title, fixtureData.Notes[12].Title)
	assert.Equal(t, rsData.Data[0].Contents, fixtureData.Notes[12].Contents)
}

func TestGetPaginatedNotesPageZero(t *testing.T) {
	testEnv, testApp := setupTest(t, mainTestEnv)
	defer tests.TeardownSingle(testEnv)

	ts := httptest.NewTLSServer(testApp.routes())
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodGet, ts.URL+"/notes?page=0", nil)

	rs, _ := ts.Client().Do(req)

	var rsData dtos.ErrorDto
	_ = helpers.ReadJSON(rs.Body, &rsData)

	assert.Equal(t, rs.StatusCode, http.StatusBadRequest)
	assert.Equal(t, rsData.Message, "invalid page query param")
}

func TestCreateNote(t *testing.T) {
	testEnv, testApp := setupTest(t, mainTestEnv)
	defer tests.TeardownSingle(testEnv)

	ts := httptest.NewTLSServer(testApp.routes())
	defer ts.Close()

	data := dtos.CreateNoteDto{
		Title:    "NewNote",
		Contents: "Some text",
	}

	body, _ := json.Marshal(data)
	req, _ := http.NewRequest(
		http.MethodPost,
		ts.URL+"/notes",
		bytes.NewReader(body),
	)

	rs, _ := ts.Client().Do(req)

	var rsData models.Note
	_ = helpers.ReadJSON(rs.Body, &rsData)

	assert.Equal(t, rs.StatusCode, http.StatusCreated)
	assert.Equal(t, rsData.Title, "NewNote")
	assert.Equal(t, rsData.Contents, "Some text")
}

func TestCreateNoteFailValidation(t *testing.T) {
	testEnv, testApp := setupTest(t, mainTestEnv)
	defer tests.TeardownSingle(testEnv)

	ts := httptest.NewTLSServer(testApp.routes())
	defer ts.Close()

	data := dtos.CreateNoteDto{
		Title:    "",
		Contents: "",
	}

	body, _ := json.Marshal(data)
	req, _ := http.NewRequest(http.MethodPost, ts.URL+"/notes", bytes.NewReader(body))

	rs, _ := ts.Client().Do(req)

	var rsData dtos.ErrorDto
	_ = helpers.ReadJSON(rs.Body, &rsData)

	assert.Equal(t, rs.StatusCode, http.StatusUnprocessableEntity)
	assert.Equal(
		t,
		rsData.Message.(map[string]interface{})["title"],
		"must be provided",
	)
}

func TestUpdateNote(t *testing.T) {
	testEnv, testApp := setupTest(t, mainTestEnv)
	defer tests.TeardownSingle(testEnv)

	ts := httptest.NewTLSServer(testApp.routes())
	defer ts.Close()

	title := "NewNote"
	data := dtos.UpdateNoteDto{
		Title: &title,
	}

	body, _ := json.Marshal(data)
	req, _ := http.NewRequest(
		http.MethodPatch,
		ts.URL+"/notes/"+fixtureData.Notes[0].ID,
		bytes.NewReader(body),
	)

	rs, _ := ts.Client().Do(req)

	var rsData models.Note
	_ = helpers.ReadJSON(rs.Body, &rsData)

	assert.Equal(t, rs.StatusCode, http.StatusOK)
	assert.Equal(t, rsData.ID, fixtureData.Notes[0].ID)
	assert.Equal(t, rsData.Title, "NewNote")
	assert.Equal(t, rsData.Contents, fixtureData.Notes[0].Contents)
}

func TestUpdateNoteNotFound(t *testing.T) {
	testEnv, testApp := setupTest(t, mainTestEnv)
	defer tests.TeardownSingle(testEnv)

	ts := httptest.NewTLSServer(testApp.routes())
	defer ts.Close()

	title := "test"
	data := dtos.UpdateNoteDto{
		Title: &title,
	}

	body, _ := json.Marshal(data)
	id, _ := uuid.NewUUID()
	req, _ := http.NewRequest(
		http.MethodPatch,
		ts.URL+"/notes/"+id.String(),
		bytes.NewReader(body),
	)

	rs, _ := ts.Client().Do(req)

	var rsData dtos.ErrorDto
	_ = helpers.ReadJSON(rs.Body, &rsData)

	assert.Equal(t, rs.StatusCode, http.StatusNotFound)
	assert.Equal(
		t,
		rsData.Message.(map[string]interface{})["id"].(string),
		fmt.Sprintf("note with id '%s' doesn't exist", id.String()),
	)
}

func TestUpdateNoteNotUUID(t *testing.T) {
	testEnv, testApp := setupTest(t, mainTestEnv)
	defer tests.TeardownSingle(testEnv)

	ts := httptest.NewTLSServer(testApp.routes())
	defer ts.Close()

	title := "test"
	data := dtos.UpdateNoteDto{
		Title: &title,
	}

	body, _ := json.Marshal(data)
	req, _ := http.NewRequest(
		http.MethodPatch,
		ts.URL+"/notes/8000",
		bytes.NewReader(body),
	)

	rs, _ := ts.Client().Do(req)

	var rsData dtos.ErrorDto
	_ = helpers.ReadJSON(rs.Body, &rsData)

	assert.Equal(t, rs.StatusCode, http.StatusBadRequest)
	assert.Contains(t, rsData.Message.(string), "invalid UUID")
}

func TestUpdateNoteFailValidation(t *testing.T) {
	testEnv, testApp := setupTest(t, mainTestEnv)
	defer tests.TeardownSingle(testEnv)

	ts := httptest.NewTLSServer(testApp.routes())
	defer ts.Close()

	title := ""
	data := dtos.UpdateNoteDto{
		Title: &title,
	}

	body, _ := json.Marshal(data)
	req, _ := http.NewRequest(
		http.MethodPatch,
		ts.URL+"/notes/"+fixtureData.Notes[0].ID,
		bytes.NewReader(body),
	)

	rs, _ := ts.Client().Do(req)

	var rsData dtos.ErrorDto
	_ = helpers.ReadJSON(rs.Body, &rsData)

	assert.Equal(t, rs.StatusCode, http.StatusUnprocessableEntity)
	assert.Equal(
		t,
		rsData.Message.(map[string]interface{})["title"],
		"must be provided",
	)
}

func TestDeleteNote(t *testing.T) {
	testEnv, testApp := setupTest(t, mainTestEnv)
	defer tests.TeardownSingle(testEnv)

	ts := httptest.NewTLSServer(testApp.routes())
	defer ts.Close()

	req, _ := http.NewRequest(
		http.MethodDelete,
		ts.URL+"/notes/"+fixtureData.Notes[0].ID,
		nil,
	)

	rs, _ := ts.Client().Do(req)

	var rsData models.Note
	_ = helpers.ReadJSON(rs.Body, &rsData)

	assert.Equal(t, rs.StatusCode, http.StatusOK)
	assert.Equal(t, rsData.ID, fixtureData.Notes[0].ID)
	assert.Equal(t, rsData.Title, fixtureData.Notes[0].Title)
	assert.Equal(t, rsData.Contents, fixtureData.Notes[0].Contents)
}

func TestDeleteNoteNotFound(t *testing.T) {
	testEnv, testApp := setupTest(t, mainTestEnv)
	defer tests.TeardownSingle(testEnv)

	ts := httptest.NewTLSServer(testApp.routes())
	defer ts.Close()

	id, _ := uuid.NewUUID()
	req, _ := http.NewRequest(http.MethodDelete, ts.URL+"/notes/"+id.String(), nil)

	rs, _ := ts.Client().Do(req)

	var rsData dtos.ErrorDto
	_ = helpers.ReadJSON(rs.Body, &rsData)

	assert.Equal(t, rs.StatusCode, http.StatusNotFound)
	assert.Equal(
		t,
		rsData.Message.(map[string]interface{})["id"].(string),
		fmt.Sprintf("note with id '%s' doesn't exist", id.String()),
	)
}

func TestDeleteNoteNotUUID(t *testing.T) {
	testEnv, testApp := setupTest(t, mainTestEnv)
	defer tests.TeardownSingle(testEnv)

	ts := httptest.NewTLSServer(testApp.routes())
	defer ts.Close()

	req, _ := http.NewRequest(http.MethodDelete, ts.URL+"/notes/8000", nil)

	rs, _ := ts.Client().Do(req)

	var rsData dtos.ErrorDto
	_ = helpers.ReadJSON(rs.Body, &rsData)

	assert.Equal(t, rs.StatusCode, http.StatusBadRequest)
	assert.Contains(t, rsData.Message.(string), "invalid UUID")
}
