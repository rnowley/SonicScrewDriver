package csharp

import (
	"testing"
)

func TestGetCSharpRunCommand(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.DestinationDirectory = "./buildtest/"
	configuration.OutputFilename = "test"
	configuration.RunArguments = []string{"argument1", "argument2"}

	// Act
	commandToTest := GetCSharpRunCommand(configuration)

	// Assert
	const expectedExecutableName = "./buildtest/test.exe"

	if commandToTest.ExecutableName != expectedExecutableName {
		t.Error(
			"For", "TestGetCSharpRunCommand",
			"expected", expectedExecutableName, "got",
			commandToTest.ExecutableName,
		)
	}
}
