package scala

import (
	"strings"
	"testing"
)

func TestNewDefaultScalacCommand(t *testing.T) {
	// Arrange

	// Act
	command := NewDefaultScalacCommand()

	// Assert
	const expectedCommand = "scalac"
	const expectedSourceDirectory = "./src/"
	const expectedDestinationDirectory = "./build/"

	if command.CommandName != expectedCommand {
		t.Error(
			"For", "TestNewDefaultScalacCommand",
			"expected", expectedCommand, "got",
			command.CommandName,
		)
	}

	if command.SourceDirectory != expectedSourceDirectory {
		t.Error(
			"For", "TestNewDefaultScalacCommand",
			"expected", expectedSourceDirectory, "got",
			command.SourceDirectory,
		)
	}

	if command.DestinationDirectory != expectedDestinationDirectory {
		t.Error(
			"For", "TestNewDefaultScalacCommand",
			"expected", expectedDestinationDirectory, "got",
			command.DestinationDirectory,
		)
	}

	if command.ClassPath == nil {
		t.Error(
			"For", "TestNewDefaultScalacCommand",
			"expected", "!= nil", "got",
			command.ClassPath,
		)
	}

}

func TestGetCommandName(t *testing.T) {
	// Arrange
	command := NewDefaultScalacCommand()

	// Act
	actualCommandName := command.GetCommandName()

	// Assert
	const expectedCommandName = "scalac"

	if actualCommandName != expectedCommandName {
		t.Error(
			"For", "TestGetCommandName",
			"expected", expectedCommandName, "got",
			actualCommandName,
		)
	}

}

func TestGetDestinationDirectory(t *testing.T) {
	// Arrange
	command := NewDefaultScalacCommand()

	// Act
	actualDestinationDirectory := command.GetDestinationDirectory()

	// Assert
	const expectedDestinationDirectory = "./build/"

	if actualDestinationDirectory != expectedDestinationDirectory {
		t.Error(
			"For", "TestGetDestinationDirectory",
			"expected", expectedDestinationDirectory, "got",
			actualDestinationDirectory,
		)
	}

}

func TestGenerateArgumentListForDefaultInstance(t *testing.T) {
	// Arrange
	command := NewDefaultScalacCommand()

	// Act
	argumentList := command.GenerateArgumentList()
	actualArgumentString := strings.Join(argumentList, " ")

	// Assert
	const expectedArgumentString = "-d ./build/"

	if actualArgumentString != expectedArgumentString {
		t.Error(
			"For", "TestGenerateArgumentListForDefaultInstance",
			"expected", expectedArgumentString, "got",
			actualArgumentString,
		)
	}

}

func TestGenerateArgumentListForInstanceWithAllFieldsSet(t *testing.T) {
	// Arrange
	command := NewDefaultScalacCommand()
	command.DestinationDirectory = "./testbuild/"
	command.Deprecation = true
	command.Verbose = true
	command.Encoding = "utf-8"
	command.Target = "jvm-1.7"
	command.Optimise = true
	command.ClassPath = []string{"./lib/a/a.jar", "./lib/b/b.jar"}
	command.DebuggingInformation = "-g:source,line"
	command.NoWarnings = true
	command.SourceFiles = []string{"./src/Main.scala", "./src/Greeter.scala"}

	// Act
	argumentList := command.GenerateArgumentList()
	argumentString := strings.Join(argumentList, " ")

	// Assert
	const expected = "-d ./testbuild/ -deprecation -verbose -encoding utf-8 " +
		"-target jvm-1.7 -optimise -classpath ./lib/a/a.jar:./lib/b/b.jar " +
		"-g:source,line -nowarn ./src/Main.scala ./src/Greeter.scala"

	if argumentString != expected {
		t.Error(
			"For", "TestGenerateArgumentListForInstanceWithAllFieldsSet",
			"expected", expected, "got",
			argumentString,
		)
	}

}

func TestExtractSourceFileListNoSourceFiles(t *testing.T) {
	// Arrange

	var configuration ScalaProject
	configuration.SourceFiles = []string{}

	// Act
	fileList := ExtractSourceFileList(configuration, "./src/")

	// Assert
	const expectedSourceFileList = ""
	actualSourceFileList := strings.Join(fileList, " ")

	if actualSourceFileList != expectedSourceFileList {
		t.Error(
			"For", "ExtractSourceFileList",
			"expected", expectedSourceFileList, "got",
			actualSourceFileList,
		)
	}

}

func TestExtractSourceFileListOneSourceFile(t *testing.T) {
	// Arrange

	var configuration ScalaProject
	configuration.SourceFiles = []string{"a.scala"}

	// Act
	fileList := ExtractSourceFileList(configuration, "./src/")

	// Assert
	const expectedSourceFileList = "./src/a.scala"
	actualSourceFileList := strings.Join(fileList, " ")

	if actualSourceFileList != expectedSourceFileList {
		t.Error(
			"For", "ExtractSourceFileList",
			"expected", expectedSourceFileList, "got",
			actualSourceFileList,
		)
	}

}

func TestExtractSourceFileListThreeSourceFiles(t *testing.T) {
	// Arrange

	var configuration ScalaProject
	configuration.SourceFiles = []string{"a.scala", "b.scala", "c.scala"}

	// Act
	fileList := ExtractSourceFileList(configuration, "./src/")

	// Assert
	const expectedSourceFileList = "./src/a.scala ./src/b.scala ./src/c.scala"
	actualSourceFileList := strings.Join(fileList, " ")

	if actualSourceFileList != expectedSourceFileList {
		t.Error(
			"For", "ExtractSourceFileList",
			"expected", expectedSourceFileList, "got",
			actualSourceFileList,
		)
	}

}
