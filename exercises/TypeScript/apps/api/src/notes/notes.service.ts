import { EntityRepository, EntityManager } from "@mikro-orm/core"
import { InjectRepository } from "@mikro-orm/nestjs"
import { Injectable } from "@nestjs/common"
import { NoteEntity } from "../entities/note"

@Injectable()
export class NotesService {
  private readonly em: EntityManager
  private readonly notesRepository: EntityRepository<NoteEntity>

  public constructor(
    em: EntityManager,
    @InjectRepository(NoteEntity)
    notesRepository: EntityRepository<NoteEntity>
  ) {
    this.em = em
    this.notesRepository = notesRepository
  }

  public async getTotalCount(): Promise<number> {
    return this.notesRepository.count()
  }

  public async getAll(): Promise<NoteEntity[]> {
    return await this.notesRepository.findAll({
      orderBy: {
        title: "asc"
      }
    })
  }

  public async getAllPaged(
    page: number,
    pageSize: number
  ): Promise<NoteEntity[]> {
    return this.notesRepository.find(
      {},
      {
        orderBy: {
          title: "asc"
        },
        limit: pageSize,
        offset: (page - 1) * pageSize
      }
    )
  }

  public async getById(id: string): Promise<NoteEntity | null> {
    return await this.notesRepository.findOne({
      id: id
    })
  }

  public async create(title: string, contents: string): Promise<NoteEntity> {
    const note = new NoteEntity(title, contents)
    await this.em.persistAndFlush(note)
    return note
  }

  public async update(
    note: NoteEntity,
    title?: string,
    contents?: string
  ): Promise<NoteEntity> {
    note.title = title ?? note.title
    note.contents = contents ?? note.contents
    await this.em.flush()
    return note
  }

  public async delete(note: NoteEntity): Promise<NoteEntity> {
    await this.em.removeAndFlush(note)
    return note
  }
}
