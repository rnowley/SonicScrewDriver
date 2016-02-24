package java

import (
	"strings"
	"testing"
)

func TestGetJavaBuildCommandDeprecatedFalse(t *testing.T) {
	// Arrange

	var configuration JavaProject
	configuration.Name = "Test Java Project"
	configuration.Version = "1.0.0"
	configuration.Description = "A project for unit testing."
	configuration.Language = "java"
	configuration.DestinationDirectory = "./build/"
	configuration.ClassPath = []string{"./lib/a.jar", "./lib/b.jar"}
	configuration.SourceFiles = []string{"a.java", "b.java", "c.java"}
	configuration.SourceVersion = "1.7"
	configuration.JarFile = "test.jar"
	configuration.RunArguments = []string{"arg1", "arg2"}
	configuration.DebuggingInformation = []string{"all"}

	// Act
	var commandToTest = GetJavaBuildCommand(configuration, false)

	// Assert
	const expectedDestinationDirectory = "./build/"

	if commandToTest.DestinationDirectory != expectedDestinationDirectory {
		t.Error(
			"For", "GetArgumentList",
			"expected", expectedDestinationDirectory, "got",
			commandToTest.DestinationDirectory,
		)
	}

	const expectedSourceDirectory = "./src/"

	if commandToTest.SourceDirectory != expectedSourceDirectory {
		t.Error(
			"For", "GetArgumentList",
			"expected", expectedSourceDirectory, "got",
			commandToTest.SourceDirectory,
		)
	}

	const expectedClassPath = "./lib/a.jar ./lib/b.jar"
	actualClassPath := strings.Join(commandToTest.ClassPath, " ")

	if actualClassPath != expectedClassPath {
		t.Error(
			"For", "GetArgumentList",
			"expected", expectedClassPath, "got",
			actualClassPath,
		)
	}

	const expectedDeprecation = false

	if commandToTest.Deprecation != expectedDeprecation {
		t.Error(
			"For", "GetArgumentList",
			"expected", expectedDeprecation, "got",
			commandToTest.Deprecation,
		)
	}

	const expectedSourceFileList = "./src/a.java ./src/b.java ./src/c.java"
	actualSourceFileList := strings.Join(commandToTest.SourceFiles, " ")

	if actualSourceFileList != expectedSourceFileList {
		t.Error(
			"For", "commandToTest.SourceFiles",
			"expected", expectedSourceFileList, "got",
			actualSourceFileList,
		)
	}

	const expectedSourceVersion = "1.7"

	if commandToTest.SourceVersion != expectedSourceVersion {
		t.Error(
			"For", "commandToTest,SourceVersion",
			"expected", expectedSourceVersion, "got",
			commandToTest.SourceVersion,
		)
	}

	const expectedDebuggingInformation = "-g"

	if commandToTest.DebuggingInformation != expectedDebuggingInformation {
		t.Error(
			"For", "commandToTest.DebuggingInformation",
			"expected", expectedDebuggingInformation, "got",
			commandToTest.DebuggingInformation,
		)
	}

}
