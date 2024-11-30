package com.sdp

import com.expediagroup.graphql.generator.SchemaGeneratorConfig
import com.expediagroup.graphql.generator.toSchema
import com.expediagroup.graphql.generator.SchemaGenerator

class SchemaGenerator {
    fun generateSchema() = toSchema(
        config = SchemaGeneratorConfig(supportedPackages = listOf("com.sdp")),
        queries = listOf(TopLevelObject(HelloWorldQuery()))
    )
}
