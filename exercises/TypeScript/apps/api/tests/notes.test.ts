/* eslint-disable max-lines-per-function */
/* eslint-disable sonarjs/no-duplicate-string */
import request from "supertest"
import {
  type CreateNoteDto,
  type GetAllPaginatedNoteDto,
  type Note,
  type UpdateNoteDto
} from "types-custom"
import Fixture, { type ErrorResponse } from "./config/fixture"
import { NoteEntity } from "../src/entities/note"
import { v4 } from "uuid"

describe("NotesController (e2e)", () => {
  const fixture: Fixture = new Fixture()

  let notes: NoteEntity[]

  const defaultPage = 1
  const defaultPageSize = 4

  beforeAll(() => {
    return fixture.beforeAll()
  })

  afterAll(() => {
    return fixture.afterAll()
  })

  beforeEach(() => {
    return fixture
      .beforeEach()
      .then(() => fixture.em.find(NoteEntity, {}))
      .then((data) => {
        notes = data
      })
  })

  afterEach(() => {
    return fixture.afterEach()
  })

  describe("/notes (GET)", () => {
    it("gets all Notes with default page (200)", async () => {
      const response = await request(fixture.app.getHttpServer())
        .get("/notes")
        .expect(200)

      const paginatedNoteResponse = response.body as GetAllPaginatedNoteDto
      expect(paginatedNoteResponse.pagination.current).toBe(defaultPage)
      expect(paginatedNoteResponse.pagination.total).toBe(
        Math.ceil((notes.length - 1) / defaultPageSize)
      )
      expect(paginatedNoteResponse.data.length).toBe(defaultPageSize)
    })

    it("gets certain page of all Notes (200)", async () => {
      const page = 2

      const response = await request(fixture.app.getHttpServer())
        .get("/notes")
        .query({ page })
        .expect(200)

      const paginatedNoteResponse = response.body as GetAllPaginatedNoteDto
      expect(paginatedNoteResponse.pagination.current).toBe(page)
      expect(paginatedNoteResponse.pagination.total).toBe(
        Math.ceil((notes.length - 1) / defaultPageSize)
      )
      expect(paginatedNoteResponse.data.length).toBe(defaultPageSize)
    })

    it("returns Page should be greater than 0 (400)", async () => {
      const page = 0

      const response = await request(fixture.app.getHttpServer())
        .get("/notes")
        .query({ page })
        .expect(400)

      const errorResponse = response.body as ErrorResponse
      expect(errorResponse.message).toBe("Page should be greater than 0")
    })
  })

  describe("/notes (POST)", () => {
    it("creates a new Note (201)", async () => {
      const data: CreateNoteDto = {
        title: "NewNote",
        contents: "Some text"
      }

      const response = await request(fixture.app.getHttpServer())
        .post("/notes")
        .send(data)
        .expect(201)

      const noteResponse = response.body as Note
      expect(noteResponse.id).toBeDefined()
      expect(noteResponse.title).toBe("NewNote")
      expect(noteResponse.contents).toBe("Some text")
    })
  })

  describe("/notes/:id (PATCH)", () => {
    it("updates a new Note (200)", async () => {
      const id = notes[1].id
      const data: UpdateNoteDto = {
        title: "NewNote2"
      }

      const response = await request(fixture.app.getHttpServer())
        .patch(`/notes/${id}`)
        .send(data)
        .expect(200)

      const noteResponse = response.body as Note
      expect(noteResponse.id).toBe(id)
      expect(noteResponse.title).toBe(data.title)
    })

    it("returns Note not found (404)", async () => {
      const data: UpdateNoteDto = {
        title: "NewSchool2"
      }

      const response = await request(fixture.app.getHttpServer())
        .patch(`/notes/${v4()}`)
        .send(data)
        .expect(404)

      const errorResponse = response.body as ErrorResponse
      expect(errorResponse.message).toBe("Note not found")
    })

    it("returns Bad request, id is uuid (400)", async () => {
      const data: UpdateNoteDto = {
        title: "NewNote"
      }

      const response = await request(fixture.app.getHttpServer())
        .patch("/notes/8000")
        .send(data)
        .expect(400)

      const errorResponse = response.body as ErrorResponse
      expect(errorResponse.message).toBe("Validation failed (uuid is expected)")
    })
  })

  describe("/notes/:id (DELETE)", () => {
    it("deletes a Note (200)", async () => {
      const id = notes[1].id

      const response = await request(fixture.app.getHttpServer())
        .delete(`/notes/${id}`)
        .expect(200)

      const noteResponse = response.body as Note
      expect(noteResponse.id).toBe(id)
    })

    it("returns Note not found (404)", async () => {
      const response = await request(fixture.app.getHttpServer())
        .delete(`/notes/${v4()}`)
        .expect(404)

      const errorResponse = response.body as ErrorResponse
      expect(errorResponse.message).toBe("Note not found")
    })

    it("returns Bad request, id is not a uuid (400)", async () => {
      const response = await request(fixture.app.getHttpServer())
        .delete("/notes/8000")
        .expect(400)

      const errorResponse = response.body as ErrorResponse
      expect(errorResponse.message).toBe("Validation failed (uuid is expected)")
    })
  })
})
