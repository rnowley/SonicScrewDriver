package scala

import (
	"strings"
	"testing"
)

func TestGetScalaTestBuildCommandDeprecationFalse(t *testing.T) {
	// Arrange
	var configuration ScalaProject
	var testConfiguration ScalaTestProject
	configuration.TestProject = testConfiguration

	// Act
	commandToTest := GetScalacTestBuildCommand(configuration, false, false)

	// Assert
	const expectedDeprecation = false

	if commandToTest.Deprecation != expectedDeprecation {
		t.Error(
			"For", "TestGetScalaTestBuildCommandDeprecationFalse",
			"expected", expectedDeprecation, "got",
			commandToTest.Deprecation,
		)
	}

}

func TestGetScalaTestBuildCommandDeprecationTrue(t *testing.T) {
	// Arrange
	var configuration ScalaProject
	var testConfiguration ScalaTestProject
	configuration.TestProject = testConfiguration

	// Act
	commandToTest := GetScalacTestBuildCommand(configuration, true, false)

	// Assert
	const expectedDeprecation = true

	if commandToTest.Deprecation != expectedDeprecation {
		t.Error(
			"For", "TestGetScalaTestBuildCommandDeprecationTrue",
			"expected", expectedDeprecation, "got",
			commandToTest.Deprecation,
		)
	}

}

func TestGetScalacTestBuildCommandVerboseTrue(t *testing.T) {
	// Arrange

	var configuration ScalaProject

	// Act

	commandToTest := GetScalacTestBuildCommand(configuration, false, true)

	// Assert
	const expectedVerbose = true

	if commandToTest.Verbose != expectedVerbose {
		t.Error(
			"For", "TestGetScalacTestBuildCommandVerboseTrue",
			"expected", expectedVerbose, "got",
			commandToTest.Verbose,
		)
	}
}

func TestGetScalacTestBuildCommandVerboseFalse(t *testing.T) {
	// Arrange
	var configuration ScalaProject

	// Act
	commandToTest := GetScalacTestBuildCommand(configuration, false, false)

	// Assert
	const expectedVerbose = false

	if commandToTest.Verbose != expectedVerbose {
		t.Error(
			"For", "TestGetScalacTestBuildCommandVerboseFalse",
			"expected", expectedVerbose, "got",
			commandToTest.Verbose,
		)
	}
}

func TestGetScalacTestBuildCommand(t *testing.T) {
	// Arrange
	var configuration ScalaProject
	configuration.Language = "scala"
	configuration.DestinationDirectory = "./build/"
	configuration.ClassPath = []string{"./lib/a.jar", "./lib/b.jar"}
	configuration.SourceFiles = []string{"a.scala", "b.scala"}

	var testConfiguration ScalaTestProject
	testConfiguration.DestinationDirectory = "./buildtest/"
	testConfiguration.SourceDirectory = "./src/test/"
	testConfiguration.SourceFiles = []string{"testA.scala", "testB.scala"}
	testConfiguration.ClassPath = []string{"./lib/c.jar", "./build/"}
	testConfiguration.RunArguments = []string{"arg1", "arg2"}

	configuration.TestProject = testConfiguration

	// Act
	commandToTest := GetScalacTestBuildCommand(configuration, false, false)

	// Assert
	const expectedDestinationDirectory = "./buildtest/"

	if commandToTest.DestinationDirectory != expectedDestinationDirectory {
		t.Error(
			"For", "TestGetScalacTestBuildCommand",
			"expected", expectedDestinationDirectory, "got",
			commandToTest.DestinationDirectory,
		)
	}

	const expectedClassPath = ". ./lib/a.jar ./lib/b.jar " +
		"./lib/c.jar ./build/"
	actualClassPath := strings.Join(commandToTest.ClassPath, " ")

	if actualClassPath != expectedClassPath {
		t.Error(
			"For", "TestGetScalacTestBuildCommand",
			"expected", expectedClassPath, "got",
			actualClassPath,
		)
	}

	const expectedSourceFileList = "./src/test/testA.scala ./src/test/testB.scala"
	actualSourceFileList := strings.Join(commandToTest.SourceFiles, " ")

	if actualSourceFileList != expectedSourceFileList {
		t.Error(
			"For", "TestGetScalacTestBuildCommand",
			"expected", expectedSourceFileList, "got",
			actualSourceFileList,
		)
	}
}
