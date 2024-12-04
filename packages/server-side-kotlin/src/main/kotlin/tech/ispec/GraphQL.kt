package tech.ispec

import com.expediagroup.graphql.server.ktor.GraphQL
import com.expediagroup.graphql.server.ktor.graphQLPostRoute
import com.expediagroup.graphql.server.ktor.graphiQLRoute
import io.ktor.server.application.*
import io.ktor.server.request.*
import io.ktor.server.response.*
import io.ktor.server.routing.*
import tech.ispec.controller.HelloWorldQuery

fun Application.configureGraphQL() {
  install(GraphQL) {
    schema {
      packages = listOf("tech.ispec.controller")
      queries = listOf(HelloWorldQuery())

      schemaObject = SdpSchema()
    }
  }

  install(Routing) {
    graphQLPostRoute()
    graphiQLRoute()
  }
}
