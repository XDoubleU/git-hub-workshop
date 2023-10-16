import { Entity, PrimaryKey, Property } from "@mikro-orm/core"
import { v4 } from "uuid"

@Entity({ tableName: "Note" })
export class NoteEntity {
  @PrimaryKey({ type: "uuid" })
  public id = v4()

  @Property()
  public title: string

  @Property()
  public contents: string

  public constructor(title: string, contents: string) {
    this.title = title
    this.contents = contents
  }
}
