package com.sdp

import com.expediagroup.graphql.server.operations.Query

class HelloWorldQuery : Query {
    fun helloWorld(): String = "Hello, GraphQL World!"
}
