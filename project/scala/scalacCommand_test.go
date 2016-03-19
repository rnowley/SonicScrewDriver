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

	// Act
	argumentList := command.GenerateArgumentList()
	argumentString := strings.Join(argumentList, " ")

	// Assert
	const expected = "-d ./testbuild/ -deprecation"

	if argumentString != expected {
		t.Error(
			"For", "TestGenerateArgumentListForInstanceWithAllFieldsSet",
			"expected", expected, "got",
			argumentString,
		)
	}

}
