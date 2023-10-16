import { Migration } from '@mikro-orm/migrations';

export class Migration20231016202540 extends Migration {

  async up(): Promise<void> {
    this.addSql('create table "Note" ("id" uuid not null, "title" varchar(255) not null, "contents" varchar(255) not null, constraint "Note_pkey" primary key ("id"));');
  }

  async down(): Promise<void> {
    this.addSql('drop table if exists "Note" cascade;');
  }

}
