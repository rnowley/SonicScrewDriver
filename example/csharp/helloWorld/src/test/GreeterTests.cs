using System;
using NUnit.Framework;

[TestFixture]
public class GreeterTests {

	[Test]
	public void TestGreeterMethod() {
		Greeter greeterToTest = new Greeter("Hello test");
		string greeting = greeterToTest.SayHello();

		Assert.That(greeting, Is.EqualTo("Hello test"));
	}

}