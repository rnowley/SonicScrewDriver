package hello

class Greeter(val name: String) {
	val greetName = name

	fun greet(): String {
		return String.format("Hello, %s", greetName)
	}

}

fun main(args: Array<String>) {
	val greeting = Greeter(args[0]).greet()
	println(greeting)
}
