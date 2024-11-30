package com.sdp

import com.expediagroup.graphql.server.operations.Query

class HelloWorldQuery : Query {
  fun hello(): String = "Hello World!"
}
