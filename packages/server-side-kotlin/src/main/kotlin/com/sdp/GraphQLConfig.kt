package com.sdp

import com.expediagroup.graphql.server.execution.GraphQL
import com.expediagroup.graphql.server.execution.GraphQLRequestHandler
import com.expediagroup.graphql.server.execution.GraphQLServer
import com.expediagroup.graphql.server.execution.GraphQLRequest
import com.expediagroup.graphql.server.execution.GraphQLResponse
import io.ktor.server.application.*
import io.ktor.server.request.*
import io.ktor.server.response.*
import io.ktor.server.routing.*
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext

fun Application.configureGraphQL() {
    val graphQL = GraphQL(SchemaGenerator().generateSchema())
    val requestHandler = GraphQLRequestHandler(graphQL)
    val server = GraphQLServer(requestParser = requestHandler, contextFactory = { null }, requestHandler = requestHandler)

    routing {
        post("/graphql") {
            val request = call.receive<GraphQLRequest>()
            val response: GraphQLResponse = withContext(Dispatchers.IO) {
                server.execute(request)
            }
            call.respond(response)
        }
    }
}
