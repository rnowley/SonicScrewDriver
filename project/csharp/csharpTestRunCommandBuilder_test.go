package csharp

import (
	"strings"
	"testing"
)

func TestGetCSharpRunTestCommand(t *testing.T) {
	// Arrange
	var configuration CSharpProject

	var testConfiguration CSharpTests
	testConfiguration.TestRunner = "Test.Runner.Console"
	testConfiguration.RunArguments = []string{"arg1", "arg2"}

	configuration.TestProject = testConfiguration

	// Act
	commandToTest := GetCSharpRunTestCommand(configuration)

	// Assert
	const expectedTestRunner = "Test.Runner.Console"
	actualTestRunner := commandToTest.ExecutableName

	if actualTestRunner != expectedTestRunner {
		t.Error(
			"For", "TestGetCSharpRunTestCommand",
			"expected", expectedTestRunner, "got",
			actualTestRunner,
		)
	}

	const expectedRunArguments = "arg1 arg2"
	actualRunArguments := strings.Join(commandToTest.RunArguments, " ")

	if actualRunArguments != expectedRunArguments {
		t.Error(
			"For", "TestGetCSharpRunTestCommand",
			"expected", expectedRunArguments, "got",
			actualRunArguments,
		)
	}
}
