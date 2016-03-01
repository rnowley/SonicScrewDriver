package kotlin

import (
	"strings"
	"testing"
)

func TestGetKotlinRunTestCommand(t *testing.T) {
	// Arrange
	var configuration KotlinProject
	configuration.ClassPath = []string{"./lib/a.jar", "./lib/b.jar"}
	configuration.DestinationDirectory = "./buildtest/"
	configuration.JarFile = "test.jar"
	configuration.OutputFilename = "test.jar"
	configuration.RunArguments = []string{"arg1", "arg2"}

	var testConfiguration KotlinTests
	testConfiguration.ClassPath = []string{"./lib/testA.jar", "./lib/testB.jar"}
	testConfiguration.OutputFilename = "unitTest.jar"
	testConfiguration.DestinationDirectory = "./buildtest/"
	testConfiguration.RunArguments = []string{"arg3", "arg4"}

	configuration.TestProject = testConfiguration

	// Act
	commandToTest := GetKotlinRunTestCommand(configuration)

	// Assert
	const expectedClassPath = ". ./buildtest/test.jar ./buildtest/unitTest.jar ./lib/a.jar ./lib/b.jar " +
		"./lib/testA.jar ./lib/testB.jar"
	actualClassPath := strings.Join(commandToTest.ClassPath, " ")

	if actualClassPath != expectedClassPath {
		t.Error(
			"For", "TestGetKotlinRunTestCommand",
			"expected", expectedClassPath, "got",
			actualClassPath,
		)
	}

	const expectedJarFile = ""

	if commandToTest.JarFile != expectedJarFile {
		t.Error(
			"For", "TestGetKotlinRunTestCommand",
			"expected", expectedJarFile, "got",
			commandToTest.JarFile,
		)
	}

	const expectedRunArguments = "arg3 arg4"
	actualRunArguments := strings.Join(commandToTest.RunArguments, " ")

	if actualRunArguments != expectedRunArguments {
		t.Error(
			"For", "TestGetKotlinRunTestCommand",
			"expected", expectedRunArguments, "got",
			actualRunArguments,
		)
	}

}