object HelloWorld {

	def main(args: Array[String]): Unit = {
		val greeter = new Greeter("Hello Raymond")
		println(greeter.greet())
	}

}