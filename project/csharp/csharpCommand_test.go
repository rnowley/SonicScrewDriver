package csharp

import (
	"strings"
	"testing"
)

func TestNewDefaultCommand(t *testing.T) {
	// Arrange

	// Act

	commandToTest := NewDefaultCommand()

	// Assert

	const expectedCommand = "mcs"
	const expectedDebugFlag = "-debug"
	const expectedSourceDirectory = "./src/"
	const expectedDestinationDirectory = "./build/"

	if commandToTest.CommandName != expectedCommand {
		t.Error(
			"For", "TestNewDefaultCommand",
			"expected", expectedCommand, "got",
			commandToTest.CommandName,
		)
	}

	if commandToTest.DebugFlag != expectedDebugFlag {
		t.Error(
			"For", "TestNewDefaultCommand",
			"expected", expectedDebugFlag, "got",
			commandToTest.DebugFlag,
		)
	}

	if commandToTest.SourceDirectory != expectedSourceDirectory {
		t.Error(
			"For", "TestNewDefaultCommand",
			"expected", expectedSourceDirectory, "got",
			commandToTest.SourceDirectory,
		)
	}

	if commandToTest.DestinationDirectory != expectedDestinationDirectory {
		t.Error(
			"For", "TestNewDefaultCommand",
			"expected", expectedDestinationDirectory, "got",
			commandToTest.DestinationDirectory,
		)
	}
}

func TestCSharpGetCommandName(t *testing.T) {
	// Arrange

	// Act
	commandToTest := NewDefaultCommand()

	// Assert
	const expectedCommandName = "mcs"

	if commandToTest.GetCommandName() != expectedCommandName {
		t.Error(
			"For", "TestCSharpGetCommandName",
			"expected", expectedCommandName, "got",
			commandToTest.GetCommandName(),
		)
	}
}

func TestGetDestinationDirectory(t *testing.T) {
	// Arrange

	// Act
	commandToTest := NewDefaultCommand()

	// Assert
	const expectedDestinationDirectory = "./build/"

	if commandToTest.GetDestinationDirectory() != expectedDestinationDirectory {
		t.Error(
			"For", "TestGetDestinationDirectory",
			"expected", expectedDestinationDirectory, "got",
			commandToTest.GetDestinationDirectory(),
		)
	}

}

func TestGenerateArgumentListForInstanceWithAllFieldsSet(t *testing.T) {
	// Arrange

	commandToTest := NewDefaultCommand()
	commandToTest.OutputFilename = "-out:./build/test"
	commandToTest.SourceFiles = []string{"a.cs", "b.cs", "c.cs"}
	commandToTest.BuildTarget = "-target:exe"
	commandToTest.LibraryPath = "-lib:lib1,lib2"
	commandToTest.References = "-r:reference1,reference2"

	// Act

	argumentList := commandToTest.GenerateArgumentList()
	actual := strings.Join(argumentList, " ")

	// Assert

	expected := "-debug -out:./build/test.exe -target:exe " +
		"-lib:lib1,lib2 -r:reference1,reference2 " +
		"a.cs b.cs c.cs"

	if actual != expected {
		t.Error(
			"For", "TestGenerateArgumentListForInstanceWithAllFieldsSet",
			"expected", expected, "got",
			actual,
		)
	}
}

func TestGetFileSuffixTargetNotProvided(t *testing.T) {
	// Arrange

	// Act

	actualSuffix := GetFileSuffix("")

	// Assert

	const expectedSuffix = ".exe"

	if actualSuffix != expectedSuffix {
		t.Error(
			"For", "TestGetFileSuffixTargetNotProvided",
			"expected", expectedSuffix, "got",
			actualSuffix,
		)
	}
}

func TestGetFileSuffixTargetInvalid(t *testing.T) {
	// Arrange

	// Act

	actualSuffix := GetFileSuffix("invalid")

	// Assert

	const expectedSuffix = ".exe"

	if actualSuffix != expectedSuffix {
		t.Error(
			"For", "TestGetFileSuffixTargetNotProvided",
			"expected", expectedSuffix, "got",
			actualSuffix,
		)
	}
}

func TestGetFileSuffixTargetExe(t *testing.T) {
	// Arrange

	// Act

	actualSuffix := GetFileSuffix("-target:exe")

	// Assert

	const expectedSuffix = ".exe"

	if actualSuffix != expectedSuffix {
		t.Error(
			"For", "TestGetFileSuffixTargetExe",
			"expected", expectedSuffix, "got",
			actualSuffix,
		)
	}
}

func TestGetFileSuffixTargetLibrary(t *testing.T) {
	// Arrange

	// Act

	actualSuffix := GetFileSuffix("-target:library")

	// Assert

	const expectedSuffix = ".dll"

	if actualSuffix != expectedSuffix {
		t.Error(
			"For", "TestGetFileSuffixTargetLibrary",
			"expected", expectedSuffix, "got",
			actualSuffix,
		)
	}
}

func TestGetFileSuffixTargetModule(t *testing.T) {
	// Arrange

	// Act

	actualSuffix := GetFileSuffix("-target:module")

	// Assert

	const expectedSuffix = ".netmodule"

	if actualSuffix != expectedSuffix {
		t.Error(
			"For", "TestGetFileSuffixTargetModule",
			"expected", expectedSuffix, "got",
			actualSuffix,
		)
	}
}

func TestGetFileSuffixTargetWinexe(t *testing.T) {
	// Arrange

	// Act

	actualSuffix := GetFileSuffix("-target:winexe")

	// Assert

	const expectedSuffix = ".exe"

	if actualSuffix != expectedSuffix {
		t.Error(
			"For", "TestGetFileSuffixTargetWinexe",
			"expected", expectedSuffix, "got",
			actualSuffix,
		)
	}
}
