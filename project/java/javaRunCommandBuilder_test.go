package java

import (
	"strings"
	"testing"
)

func TestGetJavaRunCommnd(t *testing.T) {
	// Arrange
	var configuration JavaProject
	configuration.ClassPath = []string{"./lib/a.jar", "./lib/b.jar"}
	configuration.MainClass = "TestClass"
	configuration.RunArguments = []string{"arg1", "arg2"}

	// Act
	commandToTest := GetJavaRunCommand(configuration)

	// Assert
	const expectedClassPath = ". ./lib/a.jar ./lib/b.jar"
	actualClassPath := strings.Join(commandToTest.ClassPath, " ")

	if actualClassPath != expectedClassPath {
		t.Error(
			"For", "GetRunJavaCommand",
			"expected", expectedClassPath, "got",
			actualClassPath,
		)
	}

	const expectedMainClass = "TestClass"

	if commandToTest.MainClass != expectedMainClass {
		t.Error(
			"For", "GetRunJavaCommand",
			"expected", expectedMainClass, "got",
			commandToTest.MainClass,
		)
	}

	const expectedRunArguments = "arg1 arg2"
	actualRunArguments := strings.Join(commandToTest.RunArguments, " ")

	if actualRunArguments != expectedRunArguments {
		t.Error(
			"For", "GetRunJavaCommand",
			"expected", expectedRunArguments, "got",
			actualRunArguments,
		)
	}

}
