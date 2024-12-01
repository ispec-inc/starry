package com.sdp.controller

import io.kotest.core.spec.style.FunSpec
import io.kotest.matchers.shouldBe

class HelloWorldQueryTest :
    FunSpec({ test("hello world query") { HelloWorldQuery().hello() shouldBe "Hello World!" } })
