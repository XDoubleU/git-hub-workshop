import {
  BadRequestException,
  Body,
  Controller,
  DefaultValuePipe,
  Delete,
  Get,
  NotFoundException,
  Param,
  ParseUUIDPipe,
  Patch,
  Post,
  Query
} from "@nestjs/common"
import { NotesService } from "./notes.service"
import {
  type GetAllPaginatedNoteDto,
  type CreateNoteDto,
  type UpdateNoteDto
} from "types-custom"
import { type NoteEntity } from "../entities/note"

@Controller("notes")
export class NotesController {
  private readonly notesService: NotesService

  public constructor(notesService: NotesService) {
    this.notesService = notesService
  }

  @Get()
  public async getAllPaged(
    @Query("page", new DefaultValuePipe(undefined)) queryPage?: number
  ): Promise<GetAllPaginatedNoteDto> {
    const pageSize = 4
    const current = queryPage ?? 1

    if (current <= 0) {
      throw new BadRequestException("Page should be greater than 0")
    }

    const amountOfNotes = await this.notesService.getTotalCount()
    const notes = await this.notesService.getAllPaged(current, pageSize)

    return {
      data: notes,
      pagination: {
        current,
        total: Math.ceil(amountOfNotes / pageSize)
      }
    }
  }

  @Post()
  public async create(
    @Body() createNoteDto: CreateNoteDto
  ): Promise<NoteEntity> {
    return await this.notesService.create(
      createNoteDto.title,
      createNoteDto.contents
    )
  }

  @Patch(":id")
  public async update(
    @Param("id", ParseUUIDPipe) id: string,
    @Body() updateNoteDto: UpdateNoteDto
  ): Promise<NoteEntity> {
    const note = await this.notesService.getById(id)
    if (!note) {
      throw new NotFoundException("Note not found")
    }

    return await this.notesService.update(
      note,
      updateNoteDto.title,
      updateNoteDto.contents
    )
  }

  @Delete(":id")
  public async delete(
    @Param("id", ParseUUIDPipe) id: string
  ): Promise<NoteEntity> {
    const note = await this.notesService.getById(id)
    if (!note) {
      throw new NotFoundException("Note not found")
    }

    return await this.notesService.delete(note)
  }
}
