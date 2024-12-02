plugins {
  alias(libs.plugins.kotlin.jvm)
  alias(libs.plugins.ktor)
  alias(libs.plugins.spotless)
}

java { toolchain { languageVersion = JavaLanguageVersion.of(21) } }

group = "com.sdp"

version = "0.0.1"

application {
  mainClass.set("io.ktor.server.netty.EngineMain")

  val isDevelopment: Boolean = project.ext.has("development")
  applicationDefaultJvmArgs = listOf("-Dio.ktor.development=$isDevelopment")
}

repositories { mavenCentral() }

spotless {
  kotlin {
    target("**/*.kt")
    ktfmt()
  }
}

tasks.test { useJUnitPlatform() }

dependencies {
  implementation(libs.graphql.ktor)
  implementation(libs.ktor.server.core)
  implementation(libs.ktor.server.auth)
  implementation(libs.ktor.serialization.kotlinx.json)
  implementation(libs.ktor.server.content.negotiation)
  implementation(libs.exposed.core)
  implementation(libs.exposed.jdbc)
  implementation(libs.h2)
  implementation(libs.koin.ktor)
  implementation(libs.koin.logger.slf4j)
  implementation(libs.ktor.server.cors)
  implementation(libs.ktor.server.netty)
  implementation(libs.logback.classic)
  implementation(libs.ktor.server.config.yaml)
  testImplementation(libs.ktor.server.test.host)
  testImplementation(libs.kotlin.test.junit)
  testImplementation(libs.kotest.junit)
  testImplementation(libs.kotest.assertions.core)
  testImplementation(kotlin("test"))
}
