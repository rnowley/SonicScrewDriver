using System;

public class HelloWorld {

    public static void Main(string[] args) {
        var greeter = new Greeter("Hello, World!");
        Console.WriteLine(greeter.SayHello());
    }

}