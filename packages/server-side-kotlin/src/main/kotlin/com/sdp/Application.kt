package com.sdp

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
  configureSerialization()
  configureDatabases()
  configureFrameworks()
  configureHTTP()
}
