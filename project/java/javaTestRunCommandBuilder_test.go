package java

import (
	"strings"
	"testing"
)

func TestGetJavaRunTestCommand(t *testing.T) {
	// Arrange
	var configuration JavaProject
	configuration.ClassPath = []string{"./lib/a.jar", "./lib/b.jar"}
	configuration.MainClass = "TestClass"
	configuration.RunArguments = []string{"arg1", "arg2"}

	var testConfiguration JavaTests
	testConfiguration.ClassPath = []string{"./lib/testA.jar", "./lib/testB.jar"}
	testConfiguration.MainClass = "TestMain.Class"
	testConfiguration.RunArguments = []string{"arg1", "arg2"}

	configuration.TestProject = testConfiguration

	// Act
	commandToTest := GetJavaRunTestCommand(configuration)

	// Assert
	const expectedClassPath = ". ./lib/a.jar ./lib/b.jar " +
		"./lib/testA.jar ./lib/testB.jar"
	actualClassPath := strings.Join(commandToTest.ClassPath, " ")

	if actualClassPath != expectedClassPath {
		t.Error(
			"For", "GetJavaRunTestCommand",
			"expected", expectedClassPath, "got",
			actualClassPath,
		)
	}

	const expectedMainClass = "TestMain.Class"

	if commandToTest.MainClass != expectedMainClass {
		t.Error(
			"For", "GetJavaRunTestCommand",
			"expected", expectedMainClass, "got",
			commandToTest.MainClass,
		)
	}

	const expectedRunArguments = "arg1 arg2"
	actualRunArguments := strings.Join(commandToTest.RunArguments, " ")

	if actualRunArguments != expectedRunArguments {
		t.Error(
			"For", "GetJavaRunTestCommand",
			"expected", expectedRunArguments, "got",
			actualRunArguments,
		)
	}

}
