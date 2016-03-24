package scala

import (
	"strings"
	"testing"
)

func TestGetScalacBuildCommandDeprecationTrue(t *testing.T) {
	// Arrange

	var configuration ScalaProject

	// Act

	commandToTest := GetScalacBuildCommand(configuration, true, false)

	// Assert
	const expectedDeprecation = true

	if commandToTest.Deprecation != expectedDeprecation {
		t.Error(
			"For", "TestGetScalacBuildCommandDeprecationTrue",
			"expected", expectedDeprecation, "got",
			commandToTest.Deprecation,
		)
	}

}

func TestGetScalacBuildCommandDeprecationFalse(t *testing.T) {
	// Arrange

	var configuration ScalaProject

	// Act

	commandToTest := GetScalacBuildCommand(configuration, false, false)

	// Assert
	const expectedDeprecation = false

	if commandToTest.Deprecation != expectedDeprecation {
		t.Error(
			"For", "TestGetScalacBuildCommandDeprecationFalse",
			"expected", expectedDeprecation, "got",
			commandToTest.Deprecation,
		)
	}
}

func TestGetScalacBuildCommandVerboseTrue(t *testing.T) {
	// Arrange

	var configuration ScalaProject

	// Act

	commandToTest := GetScalacBuildCommand(configuration, false, true)

	// Assert
	const expectedVerbose = true

	if commandToTest.Verbose != expectedVerbose {
		t.Error(
			"For", "TestGetScalacBuildCommandVerboseTrue",
			"expected", expectedVerbose, "got",
			commandToTest.Verbose,
		)
	}
}

func TestGetScalacBuildCommandVerboseFalse(t *testing.T) {
	// Arrange

	var configuration ScalaProject

	// Act

	commandToTest := GetScalacBuildCommand(configuration, false, false)

	// Assert
	const expectedVerbose = false

	if commandToTest.Verbose != expectedVerbose {
		t.Error(
			"For", "TestGetScalacBuildCommandVerboseFalse",
			"expected", expectedVerbose, "got",
			commandToTest.Verbose,
		)
	}
}

func TestGetScalacBuildCommand(t *testing.T) {
	// Arrange
	const expectedClassPath = "./lib/a.jar ./lib/b.jar"
	const expectedDebuggingInformation = "vars"
	const expectedDestinationDirectory = "./testbuild/"
	const expectedEncoding = "utf-8"
	const expectedNoWarnings = true
	const expectedOptimise = true
	const expectedSourceDirectory = "./testsrc/"
	const expectedSourceFiles = "./testsrc/a.java ./testsrc/b.java ./testsrc/c.java"
	const expectedTarget = "jvm-1.7"

	var configuration ScalaProject
	configuration.ClassPath = []string{"./lib/a.jar", "./lib/b.jar"}
	configuration.DebuggingInformation = expectedDebuggingInformation
	configuration.DestinationDirectory = expectedDestinationDirectory
	configuration.Encoding = expectedEncoding
	configuration.NoWarnings = expectedNoWarnings
	configuration.Optimise = expectedOptimise
	configuration.SourceDirectory = expectedSourceDirectory
	configuration.SourceFiles = []string{"a.java", "b.java", "c.java"}
	configuration.Target = expectedTarget

	// Act

	commandToTest := GetScalacBuildCommand(configuration, false, false)

	// Assert

	if commandToTest.DebuggingInformation != expectedDebuggingInformation {
		t.Error(
			"For", "TestGetScalacBuildCommand",
			"expected", expectedDebuggingInformation, "got",
			commandToTest.DebuggingInformation,
		)
	}

	actualClassPath := strings.Join(commandToTest.ClassPath, " ")

	if actualClassPath != expectedClassPath {
		t.Error(
			"For", "TestGetScalacBuildCommand",
			"expected", expectedClassPath, "got",
			actualClassPath,
		)
	}

	if commandToTest.DestinationDirectory != expectedDestinationDirectory {
		t.Error(
			"For", "TestGetScalacBuildCommand",
			"expected", expectedDestinationDirectory, "got",
			commandToTest.DestinationDirectory,
		)
	}

	if commandToTest.Encoding != expectedEncoding {
		t.Error(
			"For", "TestGetScalacBuildCommand",
			"expected", expectedEncoding, "got",
			commandToTest.Encoding,
		)
	}

	if commandToTest.NoWarnings != expectedNoWarnings {
		t.Error(
			"For", "TestGetScalacBuildCommand",
			"expected", expectedNoWarnings, "got",
			commandToTest.NoWarnings,
		)
	}

	if commandToTest.Optimise != expectedOptimise {
		t.Error(
			"For", "TestGetScalacBuildCommand",
			"expected", expectedOptimise, "got",
			commandToTest.Optimise,
		)
	}

	if commandToTest.SourceDirectory != expectedSourceDirectory {
		t.Error(
			"For", "TestGetScalacBuildCommand",
			"expected", expectedSourceDirectory, "got",
			commandToTest.SourceDirectory,
		)
	}

	actualSourceFileList := strings.Join(commandToTest.SourceFiles, " ")

	if actualSourceFileList != expectedSourceFiles {
		t.Error(
			"For", "TestGetScalacBuildCommand",
			"expected", expectedSourceFiles, "got",
			actualSourceFileList,
		)
	}

	if commandToTest.Target != expectedTarget {
		t.Error(
			"For", "TestGetScalacBuildCommand",
			"expected", expectedTarget, "got",
			commandToTest.Target,
		)
	}

}
