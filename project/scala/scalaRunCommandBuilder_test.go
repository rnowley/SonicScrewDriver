package scala

import (
	"strings"
	"testing"
)

func TestGetScalaRunCommand(t *testing.T) {
	// Arrange
	var configuration ScalaProject
	configuration.ClassPath = []string{"./lib/a.jar", "./lib/b.jar"}
	configuration.MainClass = "TestClass"
	configuration.RunArguments = []string{"arg1", "arg2"}

	// Act
	commandToTest := GetScalaRunCommand(configuration)

	// Assert
	const expectedClassPath = ". ./lib/a.jar ./lib/b.jar"
	actualClassPath := strings.Join(commandToTest.ClassPath, " ")

	if actualClassPath != expectedClassPath {
		t.Error(
			"For", "TestGetScalaRunCommand",
			"expected", expectedClassPath, "got",
			actualClassPath,
		)
	}

	const expectedMainClass = "TestClass"

	if commandToTest.MainClass != expectedMainClass {
		t.Error(
			"For", "TestGetScalaRunCommand",
			"expected", expectedMainClass, "got",
			commandToTest.MainClass,
		)
	}

	const expectedRunArguments = "arg1 arg2"
	actualRunArguments := strings.Join(commandToTest.RunArguments, " ")

	if actualRunArguments != expectedRunArguments {
		t.Error(
			"For", "TestGetScalaRunCommand",
			"expected", expectedRunArguments, "got",
			actualRunArguments,
		)
	}
}
