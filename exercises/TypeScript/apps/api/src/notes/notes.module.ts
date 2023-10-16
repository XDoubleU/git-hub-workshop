/* eslint-disable @typescript-eslint/no-extraneous-class */
import { MikroOrmModule } from "@mikro-orm/nestjs"
import { Module } from "@nestjs/common"
import { NoteEntity } from "../entities/note"
import { NotesController } from "./notes.controller"
import { NotesService } from "./notes.service"

@Module({
  imports: [MikroOrmModule.forFeature([NoteEntity])],
  controllers: [NotesController],
  providers: [NotesService],
  exports: [NotesService]
})
export class NotesModule {}
