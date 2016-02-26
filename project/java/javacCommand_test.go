package java

import (
	"strings"
	"testing"
)

func TestNewDefaultJavacCommand(t *testing.T) {
	const expectedCommand = "javac"
	const expectedSourceDirectory = "./src/"
	const expectedDestinationDirectory = "./build/"

	command := NewDefaultJavacCommand()

	if command.CommandName != expectedCommand {
		t.Error(
			"For", "command.CommandName",
			"expected", expectedCommand, "got",
			command.CommandName,
		)
	}

	if command.SourceDirectory != expectedSourceDirectory {
		t.Error(
			"For", "command.SourceDirectory",
			"expected", expectedSourceDirectory, "got",
			command.SourceDirectory,
		)
	}

	if command.DestinationDirectory != expectedDestinationDirectory {
		t.Error(
			"For", "command.DestinationDirectory",
			"expected", expectedDestinationDirectory, "got",
			command.DestinationDirectory,
		)
	}

	if command.ClassPath == nil {
		t.Error(
			"For", "command.ClassPath",
			"expected", "!= nil", "got",
			command.ClassPath,
		)
	}

}

func TestGetCommandName(t *testing.T) {
	command := NewDefaultJavacCommand()

	commandName := command.GetCommandName()

	if commandName != "javac" {
		t.Error(
			"For", "command.GetCommandName",
			"expected", "javac", "got",
			commandName,
		)
	}

}

func TestGenerateArgumentListForDefaultInstance(t *testing.T) {
	command := NewDefaultJavacCommand()

	argumentList := command.GenerateArgumentList()
	argumentString := strings.Join(argumentList, " ")

	if argumentString != "-d ./build/" {
		t.Error(
			"For", "command.GetCommandName",
			"expected", "-d ./build/", "got",
			argumentString,
		)
	}
}

func TestGenerateArgumentListForInstanceWithAllFieldsSet(t *testing.T) {
	const expected = "-d ./build/ -deprecation " +
		"a.java b.java c.java " +
		"-cp .:./lib/a.jar:./lib/x/b.jar " +
		"-source 1.7"

	command := NewDefaultJavacCommand()
	command.SourceFiles = []string{"a.java", "b.java", "c.java"}
	command.ClassPath = []string{".", "./lib/a.jar", "./lib/x/b.jar"}
	command.Deprecation = true
	command.SourceVersion = "1.7"

	argumentList := command.GenerateArgumentList()
	argumentString := strings.Join(argumentList, " ")

	if argumentString != expected {
		t.Error(
			"For", "TestGenerateArgumentListForInstanceWithAllFieldsSet",
			"expected", expected, "got",
			argumentString,
		)
	}
}
