package kotlin

import (
	"strings"
	"testing"
)

func TestGetKotlinRunCommand(t *testing.T) {

	// Arrange
	var configuration KotlinProject
	configuration.ClassPath = []string{"./lib/a.jar", "./lib/b.jar"}
	configuration.DestinationDirectory = "./testbuild/"
	configuration.JarFile = "test.jar"
	configuration.RunArguments = []string{"arg1", "arg2"}

	// Act
	commandToTest := GetKotlinRunCommand(configuration)

	// Assert
	const expectedClassPath = ". ./lib/a.jar ./lib/b.jar"
	actualClassPath := strings.Join(commandToTest.ClassPath, " ")

	if actualClassPath != expectedClassPath {
		t.Error(
			"For", "TestGetKotlinRunCommand",
			"expected", expectedClassPath, "got",
			actualClassPath,
		)
	}

	const expectedJarFile = "./testbuild/test.jar"

	if commandToTest.JarFile != expectedJarFile {
		t.Error(
			"For", "TestGetKotlinRunCommand",
			"expected", expectedJarFile, "got",
			commandToTest.JarFile,
		)
	}

	const expectedRunArguments = "arg1 arg2"
	actualRunArguments := strings.Join(commandToTest.RunArguments, " ")

	if actualRunArguments != expectedRunArguments {
		t.Error(
			"For", "TestGetKotlinRunCommand",
			"expected", expectedRunArguments, "got",
			actualRunArguments,
		)
	}

}
