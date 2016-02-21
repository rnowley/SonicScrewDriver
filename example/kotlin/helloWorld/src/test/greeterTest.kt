package hello

import org.junit.Test
import junit.framework.TestCase
import kotlin.test.*
import org.junit.After
import org.junit.Before

// java -cp .:../lib/junit-4.12.jar:../lib/hamcrest-core-1.3.jar:./outTest.jar:./out.jar org.junit.runner.JUnitCore hello.GreeterTest

public class GreeterTest : TestCase() {

	fun testGreetMethod() {
		val greeterToTest = Greeter("test")
		val greeting = greeterToTest.greet()
		assertEquals("Hello, test", greeting)
	}
}
