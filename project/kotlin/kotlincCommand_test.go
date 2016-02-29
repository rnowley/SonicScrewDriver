package kotlin

import (
	"strings"
	"testing"
)

func TestNewDefaultKotlincCommand(t *testing.T) {
	// Arrange

	// Act

	commandToTest := NewDefaultKotlincCommand()

	// Assert
	const expectedCommandName = "kotlinc"

	if commandToTest.CommandName != expectedCommandName {
		t.Error(
			"For", "TestNewDefaultKotlincCommand",
			"expected", expectedCommandName, "got",
			commandToTest.CommandName,
		)
	}

	const expectedSourceDirectory = "./src/"

	if commandToTest.SourceDirectory != expectedSourceDirectory {
		t.Error(
			"For", "TestNewDefaultKotlincCommand",
			"expected", expectedSourceDirectory, "got",
			commandToTest.SourceDirectory,
		)
	}

	const expectedDestinationDirectory = "./build/"

	if commandToTest.DestinationDirectory != expectedDestinationDirectory {
		t.Error(
			"For", "TestNewDefaultKotlincCommand",
			"expected", expectedDestinationDirectory, "got",
			commandToTest.DestinationDirectory,
		)
	}

	if commandToTest.ClassPath == nil {
		t.Error(
			"For", "TestNewDefaultKotlincCommand",
			"expected", "!= nil", "got",
			commandToTest.ClassPath,
		)
	}

	const expectedOutputFilename = "out.jar"

	if commandToTest.OutputFilename != expectedOutputFilename {
		t.Error(
			"For", "TestNewDefaultKotlincCommand",
			"expected", expectedOutputFilename, "got",
			commandToTest.OutputFilename,
		)
	}

	if commandToTest.SourceFiles == nil {
		t.Error(
			"For", "TestNewDefaultKotlincCommand",
			"expected", "!= nil", "got",
			commandToTest.SourceFiles,
		)
	}
}

func TestKotlincGetCommandName(t *testing.T) {
	// Arrange

	// Act
	commandToTest := NewDefaultKotlincCommand()
	actualCommandName := commandToTest.GetCommandName()

	// Assert
	const expectedCommandName = "kotlinc"

	if actualCommandName != expectedCommandName {
		t.Error(
			"For", "TestKotlincGetCommandName",
			"expected", expectedCommandName, "got",
			actualCommandName,
		)
	}

}

func TestGenerateArgumentListForDefaultInstance(t *testing.T) {
	// Arrange

	// Act
	commandToTest := NewDefaultKotlincCommand()

	argumentList := commandToTest.GenerateArgumentList()
	actualArgumentString := strings.Join(argumentList, " ")

	// Assert

	const expectedArgumentString = "-d ./build/out.jar"

	if actualArgumentString != expectedArgumentString {
		t.Error(
			"For", "TestGenerateArgumentListForDefaultInstance",
			"expected", "-d ./build/", "got",
			actualArgumentString,
		)
	}
}

func TestGenerateArgumentListForInstanceWithAllFieldsSet(t *testing.T) {
	// Arrange

	command := NewDefaultKotlincCommand()
	command.BuildTarget = "executable"
	command.ClassPath = []string{".", "./lib/a.jar", "./lib/x/b.jar"}
	command.DestinationDirectory = "./testbuild/"
	command.OutputFilename = "testout.jar"
	command.SourceFiles = []string{"a.kt", "b.kt", "c.kt"}

	// Act

	argumentList := command.GenerateArgumentList()
	actualArgumentString := strings.Join(argumentList, " ")

	// Assert

	const expectedArgumentString = "-d ./testbuild/testout.jar " +
		"-cp .:./lib/a.jar:./lib/x/b.jar -include-runtime " +
		"a.kt b.kt c.kt"

	if actualArgumentString != expectedArgumentString {
		t.Error(
			"For", "TestGenerateArgumentListForInstanceWithAllFieldsSet",
			"expected", expectedArgumentString, "got",
			actualArgumentString,
		)
	}
}
