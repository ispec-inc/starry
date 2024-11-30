package com.sdp

import com.expediagroup.graphql.server.execution.GraphQL
import com.expediagroup.graphql.server.execution.GraphQLRequestHandler
import com.expediagroup.graphql.server.execution.GraphQLServer
import com.expediagroup.graphql.server.execution.GraphQLServerRequest
import com.expediagroup.graphql.server.execution.GraphQLServerResponse
import io.ktor.server.application.*
import io.ktor.server.request.*
import io.ktor.server.response.*
import io.ktor.server.routing.*
import kotlinx.coroutines.Dispatchers
import kotlinx.coroutines.withContext

fun Application.configureGraphQL() {
    val graphQL = GraphQL.newGraphQL(SchemaGenerator().generateSchema()).build()
    val requestHandler = GraphQLRequestHandler(graphQL)
    val server = GraphQLServer(requestHandler)

    routing {
        post("/graphql") {
            val request = call.receive<GraphQLServerRequest>()
            val response: GraphQLServerResponse = withContext(Dispatchers.IO) {
                server.execute(request)
            }
            call.respond(response)
        }
    }
}
