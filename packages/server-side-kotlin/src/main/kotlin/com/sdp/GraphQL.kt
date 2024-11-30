package com.sdp

import com.expediagroup.graphql.server.ktor.GraphQL
import com.expediagroup.graphql.server.ktor.graphQLPostRoute
import com.expediagroup.graphql.server.ktor.graphiQLRoute
import io.ktor.server.application.*
import io.ktor.server.request.*
import io.ktor.server.response.*
import io.ktor.server.routing.*

fun Application.configureGraphQL() {
  install(GraphQL) {
    schema {
      packages = listOf("com.sdp")
      queries = listOf(HelloWorldQuery())
    }
  }

  install(Routing) {
    graphQLPostRoute()
    graphiQLRoute()
  }
}
