package kotlin

import (
	"strings"
	"testing"
)

func TestGetKotlincBuildCommand(t *testing.T) {
	// Arrange

	var configuration KotlinProject
	configuration.Name = "Test Kotlin Project"
	configuration.Version = "1.0.0"
	configuration.Description = "A project for unit testing."
	configuration.Language = "kotlin"
	configuration.DestinationDirectory = "./build/"
	configuration.JarFile = "test.jar"
	configuration.OutputFilename = "out.jar"
	configuration.ClassPath = []string{"./lib/a.jar", "./lib/b.jar"}
	configuration.SourceFiles = []string{"a.kt", "b.kt", "c.kt"}
	configuration.BuildTarget = "exectuable"

	configuration.RunArguments = []string{"arg1", "arg2"}

	// Act
	var commandToTest = GetKotlincBuildCommand(configuration)

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

	const expectedOutputFilename = "out.jar"

	if commandToTest.OutputFilename != expectedOutputFilename {
		t.Error(
			"For", "GetArgumentList",
			"expected", expectedOutputFilename, "got",
			commandToTest.OutputFilename,
		)
	}

	const expectedSourceFileList = "./src/a.kt ./src/b.kt ./src/c.kt"
	actualSourceFileList := strings.Join(commandToTest.SourceFiles, " ")

	if actualSourceFileList != expectedSourceFileList {
		t.Error(
			"For", "commandToTest.SourceFiles",
			"expected", expectedSourceFileList, "got",
			actualSourceFileList,
		)
	}

}

func TestKotlincExtractSourceFileListNoSourceFiles(t *testing.T) {
	// Arrange

	var configuration KotlinProject
	configuration.SourceFiles = []string{}

	// Act
	fileList := ExtractSourceFileList(configuration, "./src/")

	// Assert
	const expectedSourceFileList = ""
	actualSourceFileList := strings.Join(fileList, " ")

	if actualSourceFileList != expectedSourceFileList {
		t.Error(
			"For", "TestKotlincExtractSourceFileListNoSourceFiles",
			"expected", expectedSourceFileList, "got",
			actualSourceFileList,
		)
	}

}

func TestKotlincExtractSourceFileListOneSourceFile(t *testing.T) {
	// Arrange

	var configuration KotlinProject
	configuration.SourceFiles = []string{"a.java"}

	// Act
	fileList := ExtractSourceFileList(configuration, "./src/")

	// Assert
	const expectedSourceFileList = "./src/a.java"
	actualSourceFileList := strings.Join(fileList, " ")

	if actualSourceFileList != expectedSourceFileList {
		t.Error(
			"For", "TestKotlincExtractSourceFileListNoSourceFiles",
			"expected", expectedSourceFileList, "got",
			actualSourceFileList,
		)
	}

}

func TestKotlincExtractSourceFileListThreeSourceFiles(t *testing.T) {
	// Arrange

	var configuration KotlinProject
	configuration.SourceFiles = []string{"a.kt", "b.kt", "c.kt"}

	// Act
	fileList := ExtractSourceFileList(configuration, "./src/")

	// Assert
	const expectedSourceFileList = "./src/a.kt ./src/b.kt ./src/c.kt"
	actualSourceFileList := strings.Join(fileList, " ")

	if actualSourceFileList != expectedSourceFileList {
		t.Error(
			"For", "TestKotlincExtractSourceFileListNoSourceFiles",
			"expected", expectedSourceFileList, "got",
			actualSourceFileList,
		)
	}

}
