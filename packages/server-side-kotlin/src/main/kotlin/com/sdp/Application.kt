package com.sdp

import io.ktor.server.application.*

fun main(args: Array<String>) {
  io.ktor.server.netty.EngineMain.main(args)
}

fun Application.module() {
  configureSecurity()
  configureSerialization()
  configureDatabases()
  configureFrameworks()
  configureHTTP()
  configureRouting()
  configureGraphQL()
}
