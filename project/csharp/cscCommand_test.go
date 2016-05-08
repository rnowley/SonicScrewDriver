package csharp

import (
	"strings"
	"testing"
)

func TestNewDefaultCscCommand(t *testing.T) {
	// Arrange

	// Act

	commandToTest := NewDefaultCscCommand()
	// Assert

	const expectedCommand = "csc"
	const expectedDebugFlag = "/debug"
	const expectedSourceDirectory = "./src/"
	const expectedDestinationDirectory = "./build/"

	if commandToTest.CommandName != expectedCommand {
		t.Error(
			"For", "TestNewDefaultCscCommand",
			"expected", expectedCommand, "got",
			commandToTest.CommandName,
		)
	}

	if commandToTest.DebugFlag != expectedDebugFlag {
		t.Error(
			"For", "TestNewDefaultCscCommand",
			"expected", expectedDebugFlag, "got",
			commandToTest.DebugFlag,
		)
	}

	if commandToTest.SourceDirectory != expectedSourceDirectory {
		t.Error(
			"For", "TestNewDefaultCscCommand",
			"expected", expectedSourceDirectory, "got",
			commandToTest.SourceDirectory,
		)
	}

	if commandToTest.DestinationDirectory != expectedDestinationDirectory {
		t.Error(
			"For", "TestNewDefaultCscCommand",
			"expected", expectedDestinationDirectory, "got",
			commandToTest.DestinationDirectory,
		)
	}
}

func TestCscGetCommandName(t *testing.T) {
	// Arrange

	// Act
	commandToTest := NewDefaultCscCommand()

	// Assert
	const expectedCommandName = "csc"

	if commandToTest.GetCommandName() != expectedCommandName {
		t.Error(
			"For", "TestCscGetCommandName",
			"expected", expectedCommandName, "got",
			commandToTest.GetCommandName(),
		)
	}
}

func TestCscGetDestinationDirectory(t *testing.T) {
	// Arrange

	// Act
	commandToTest := NewDefaultCscCommand()

	// Assert
	const expectedDestinationDirectory = "./build/"

	if commandToTest.GetDestinationDirectory() != expectedDestinationDirectory {
		t.Error(
			"For", "TestCscGetDestinationDirectory",
			"expected", expectedDestinationDirectory, "got",
			commandToTest.GetDestinationDirectory(),
		)
	}

}

func TestGenerateCscArgumentListForInstanceWithAllFieldsSet(t *testing.T) {
	// Arrange

	commandToTest := NewDefaultCscCommand()
	commandToTest.OutputFilename = "/out:./build/test"
	commandToTest.SourceFiles = []string{"a.cs", "b.cs", "c.cs"}
	commandToTest.BuildTarget = "/target:exe"
	commandToTest.LibraryPath = "/lib:lib1,lib2"
	commandToTest.References = "/r:reference1,reference2"

	// Act

	argumentList := commandToTest.GenerateArgumentList()
	actual := strings.Join(argumentList, " ")

	// Assert

	expected := "/debug /out:./build/test.exe /target:exe " +
		"/lib:lib1,lib2 /r:reference1,reference2 " +
		"a.cs b.cs c.cs"

	if actual != expected {
		t.Error(
			"For", "TestGenerateCscArgumentListForInstanceWithAllFieldsSet",
			"expected", expected, "got",
			actual,
		)
	}
}

func TestGetCscFileSuffixTargetNotProvided(t *testing.T) {
	// Arrange

	// Act

	actualSuffix := GetCscFileSuffix("")

	// Assert

	const expectedSuffix = ".exe"

	if actualSuffix != expectedSuffix {
		t.Error(
			"For", "TestGetCscFileSuffixTargetNotProvided",
			"expected", expectedSuffix, "got",
			actualSuffix,
		)
	}
}

func TestCscGetFileSuffixTargetInvalid(t *testing.T) {
	// Arrange

	// Act

	actualSuffix := GetCscFileSuffix("invalid")

	// Assert

	const expectedSuffix = ".exe"

	if actualSuffix != expectedSuffix {
		t.Error(
			"For", "TestGetCscFileSuffixTargetNotProvided",
			"expected", expectedSuffix, "got",
			actualSuffix,
		)
	}
}

func TestGetCscFileSuffixTargetExe(t *testing.T) {
	// Arrange

	// Act

	actualSuffix := GetCscFileSuffix("/target:exe")

	// Assert

	const expectedSuffix = ".exe"

	if actualSuffix != expectedSuffix {
		t.Error(
			"For", "TestGetCscFileSuffixTargetExe",
			"expected", expectedSuffix, "got",
			actualSuffix,
		)
	}
}

func TestGetCscFileSuffixTargetLibrary(t *testing.T) {
	// Arrange

	// Act

	actualSuffix := GetCscFileSuffix("/target:library")

	// Assert

	const expectedSuffix = ".dll"

	if actualSuffix != expectedSuffix {
		t.Error(
			"For", "TestGetCscFileSuffixTargetLibrary",
			"expected", expectedSuffix, "got",
			actualSuffix,
		)
	}
}

func TestGetCscFileSuffixTargetModule(t *testing.T) {
	// Arrange

	// Act

	actualSuffix := GetCscFileSuffix("/target:module")

	// Assert

	const expectedSuffix = ".netmodule"

	if actualSuffix != expectedSuffix {
		t.Error(
			"For", "TestGetCscFileSuffixTargetModule",
			"expected", expectedSuffix, "got",
			actualSuffix,
		)
	}
}

func TestGetCscFileSuffixTargetWinexe(t *testing.T) {
	// Arrange

	// Act

	actualSuffix := GetCscFileSuffix("/target:winexe")

	// Assert

	const expectedSuffix = ".exe"

	if actualSuffix != expectedSuffix {
		t.Error(
			"For", "TestGetCscFileSuffixTargetWinexe",
			"expected", expectedSuffix, "got",
			actualSuffix,
		)
	}
}

func TestGetCscFileSuffixTargetWinmdobj(t *testing.T) {
	// Arrange

	// Act

	actualSuffix := GetCscFileSuffix("/target:winmdobj")

	// Assert

	const expectedSuffix = ".winmdobj"

	if actualSuffix != expectedSuffix {
		t.Error(
			"For", "TestGetCscFileSuffixTargetWinmdobj",
			"expected", expectedSuffix, "got",
			actualSuffix,
		)
	}
}
