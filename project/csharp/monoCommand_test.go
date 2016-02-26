package csharp

import (
	"strings"
	"testing"
)

func TestNewDefaultMonoCommand(t *testing.T) {

	// Act
	commandToTest := NewDefaultMonoCommand()

	// Assert
	const expectedCommandName = "mono"

	if commandToTest.CommandName != expectedCommandName {
		t.Error(
			"For", "TestNewDefaultMonoCommand",
			"expected", expectedCommandName, "got",
			commandToTest.CommandName,
		)
	}
}

func TestMonoGetCommandName(t *testing.T) {
	// Arrange

	// Act
	commandToTest := NewDefaultMonoCommand()

	// Assert
	const expectedCommandName = "mono"

	if commandToTest.GetCommandName() != expectedCommandName {
		t.Error(
			"For", "TestMonoGetCommandName",
			"expected", expectedCommandName, "got",
			commandToTest.GetCommandName(),
		)
	}
}

func TestMonoGenerateArgumentListForInstanceWithAllFieldsSet(t *testing.T) {
	// Arrange

	commandToTest := NewDefaultMonoCommand()
	commandToTest.ExecutableName = "test.exe"
	commandToTest.RunArguments = []string{"arg1", "arg2"}

	// Act

	argumentList := commandToTest.GenerateArgumentList()
	actual := strings.Join(argumentList, " ")

	// Assert

	expected := "test.exe arg1 arg2"

	if actual != expected {
		t.Error(
			"For", "func TestMonoGenerateArgumentListForInstanceWithAllFieldsSet",
			"expected", expected, "got",
			actual,
		)
	}
}
