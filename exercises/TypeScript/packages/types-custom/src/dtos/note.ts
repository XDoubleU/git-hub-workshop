import { type Note } from "../types"
import { type Pagination } from "./shared"

export interface CreateNoteDto {
  title: string
  contents: string
}

export interface GetAllPaginatedNoteDto {
  data: Note[]
  pagination: Pagination
}

export interface UpdateNoteDto {
  title?: string
  contents?: string
}
