package scala

import (
	"strings"
	"testing"
)

func TestGetScalaRunTestCommand(t *testing.T) {
	// Arrange
	var configuration ScalaProject
	configuration.ClassPath = []string{"./lib/a.jar", "./lib/b.jar"}

	var testConfiguration ScalaTestProject
	testConfiguration.ClassPath = []string{"./lib/testA.jar", "./lib/testB.jar"}
	testConfiguration.MainClass = "TestMain.Class"
	testConfiguration.RunArguments = []string{"arg1", "arg2"}

	configuration.TestProject = testConfiguration

	// Act
	commandToTest := GetScalaRunTestCommand(configuration)

	// Assert
	const expectedClassPath = ". ./lib/a.jar ./lib/b.jar " +
		"./lib/testA.jar ./lib/testB.jar"
	actualClassPath := strings.Join(commandToTest.ClassPath, " ")

	if actualClassPath != expectedClassPath {
		t.Error(
			"For", "TestGetScalaRunTestCommand",
			"expected", expectedClassPath, "got",
			actualClassPath,
		)
	}

	const expectedMainClass = "TestMain.Class"

	if commandToTest.MainClass != expectedMainClass {
		t.Error(
			"For", "TestGetScalaRunTestCommand",
			"expected", expectedMainClass, "got",
			commandToTest.MainClass,
		)
	}

	const expectedRunArguments = "arg1 arg2"
	actualRunArguments := strings.Join(commandToTest.RunArguments, " ")

	if actualRunArguments != expectedRunArguments {
		t.Error(
			"For", "TestGetScalaRunTestCommand",
			"expected", expectedRunArguments, "got",
			actualRunArguments,
		)
	}
}
