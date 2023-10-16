/* eslint-disable @typescript-eslint/no-extraneous-class */
import { Module } from "@nestjs/common"
import { APP_GUARD } from "@nestjs/core"
import { MikroOrmModule } from "@mikro-orm/nestjs"
import { ThrottlerGuard, ThrottlerModule } from "@nestjs/throttler"
import sharedConfig from "./shared-config"
import { NotesModule } from "./notes/notes.module"

const modules = [
  MikroOrmModule.forRoot({
    ...sharedConfig,
    driverOptions: {
      ...(process.env.NODE_ENV === "production" && {
        connection: { ssl: { ca: process.env.CA_CERT } }
      })
    }
  }),
  NotesModule
]

if (process.env.THROTTLE_DISABLED !== "true") {
  modules.push(
    ThrottlerModule.forRoot({
      ttl: 10, // the number of seconds that each request will last in storage
      limit: 30 // the maximum number of requests within the TTL limit
    })
  )
}

@Module({
  imports: modules,
  providers: [
    {
      provide: APP_GUARD,
      useClass: ThrottlerGuard
    }
  ]
})
export class AppModule {}
