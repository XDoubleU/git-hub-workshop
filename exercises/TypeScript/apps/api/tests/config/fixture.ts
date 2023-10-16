import { ValidationPipe, type INestApplication } from "@nestjs/common"
import { Test } from "@nestjs/testing"
import { AppModule } from "../../src/app.module"
import cookieParser from "cookie-parser"
import { NoteEntity } from "../../src/entities/note"
import { MikroORM, type Transaction } from "@mikro-orm/core"
import {
  type Knex,
  type EntityManager,
  type PostgreSqlDriver
} from "@mikro-orm/postgresql"
import { TestModule } from "./test.module"
import { ContextManager } from "./test.middleware"
import helmet from "helmet"

export interface ErrorResponse {
  message: string
}

export default class Fixture {
  public app!: INestApplication
  public em!: EntityManager
  public contextManager!: ContextManager
  public mainTransaction!: Transaction<Knex.Transaction>

  public async beforeAll(): Promise<void> {
    const module = await Test.createTestingModule({
      imports: [
        // Import the AppModule without any change to config
        AppModule,
        // Add the test module to register the TransactionContextMiddleware
        TestModule
      ]
    }).compile()

    this.app = module.createNestApplication()
    this.app.use(helmet())
    this.app.use(cookieParser())
    this.app.useGlobalPipes(new ValidationPipe({ transform: true }))

    this.app.enableShutdownHooks()

    await this.app.init()

    const orm = this.app.get<MikroORM<PostgreSqlDriver>>(MikroORM)
    this.contextManager = this.app.get(ContextManager)

    this.em = orm.em.fork()

    this.mainTransaction = await this.em.getConnection().begin()
    this.em.setTransactionContext(this.mainTransaction)

    await this.seedDatabase()
  }

  public async beforeEach(): Promise<void> {
    const testTransaction = await this.em
      .getConnection()
      .begin({ ctx: this.mainTransaction })
    this.em.setTransactionContext(testTransaction)
    this.contextManager.setContext(testTransaction)
  }

  public async afterEach(): Promise<void> {
    const testTransaction = this.contextManager.resetContext()
    if (!testTransaction) {
      throw new Error("testTransaction is undefined")
    }

    await this.em.getConnection().rollback(testTransaction)
    this.em.clear()
  }

  public async afterAll(): Promise<void> {
    await this.em.getConnection().rollback(this.mainTransaction)
    await this.app.close()
  }

  // eslint-disable-next-line max-lines-per-function
  private async seedDatabase(): Promise<void> {
    const notes: NoteEntity[] = []
    for (let i = 0; i < 20; i++) {
      const newNote = new NoteEntity(`TestNote${i}`, "Some text")
      notes.push(newNote)
    }

    await this.em.persistAndFlush(notes)

    this.em.clear()
  }
}
