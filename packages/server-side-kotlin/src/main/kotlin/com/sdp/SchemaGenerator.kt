package com.sdp

import com.expediagroup.graphql.generator.SchemaGeneratorConfig
import com.expediagroup.graphql.generator.toSchema
import com.expediagroup.graphql.server.execution.SchemaGenerator

class SchemaGenerator : SchemaGenerator {
    override fun generateSchema() = toSchema(
        config = SchemaGeneratorConfig(supportedPackages = listOf("com.sdp")),
        queries = listOf(HelloWorldQuery())
    )
}
