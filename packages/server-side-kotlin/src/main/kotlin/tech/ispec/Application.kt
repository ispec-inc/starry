package tech.ispec

import io.ktor.server.application.*
import io.ktor.server.request.*
import io.ktor.server.response.*
import io.ktor.server.routing.*

fun main(args: Array<String>) {
  io.ktor.server.netty.EngineMain.main(args)
}

fun Application.module() {
  configureGraphQL()
  configureSecurity()
  configureDatabases()
  configureFrameworks()
  configureHTTP()
}
