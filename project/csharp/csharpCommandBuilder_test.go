package csharp

import (
	"fmt"
	"strings"
	"testing"
)

func TestGetCSharpBuildCommand(t *testing.T) {
	// Arrange

	var configuration CSharpProject
	configuration.Name = "Test C# Project"
	configuration.Version = "1.0.0"
	configuration.Description = "A project for unit testing."
	configuration.Language = "csharp"
	configuration.References = []Reference{Reference{Name: "Library.x",
		Path: "./lib/Library.x/"}, Reference{Name: "Library.y",
		Path: "./lib/Library.y/"}}
	configuration.SourceFiles = []string{"a.cs", "b.cs", "c.cs"}
	configuration.Resources = []Resource{Resource{Source: "view/index.html",
		Destination: "view/index.html"}}
	configuration.BuildTarget = "exe"
	configuration.OutputFilename = "test"
	configuration.SourceDirectory = "./testsrc/"
	configuration.DestinationDirectory = "./testbuild/"
	configuration.LibraryPath = []string{"./lib/lib1", "./lib/lib2"}
	configuration.PackageList = []string{"package1", "package2"}
	configuration.RunArguments = []string{"arg1", "arg2"}

	// Act
	commandToTest := GetCSharpBuildCommand(configuration)

	// Assert

	const expectedSourceDirectory = "./testsrc/"

	if commandToTest.SourceDirectory != expectedSourceDirectory {
		t.Error(
			"For", "GetCSharpBuildCommand",
			"expected", expectedSourceDirectory, "got",
			commandToTest.SourceDirectory,
		)
	}

	const expectedDestinationDirectory = "./testbuild/"

	if commandToTest.DestinationDirectory != expectedDestinationDirectory {
		t.Error(
			"For", "GetCSharpBuildCommand",
			"expected", expectedDestinationDirectory, "got",
			commandToTest.DestinationDirectory,
		)
	}

	const expectedOutputFilename = "-out:./testbuild/test"

	if commandToTest.OutputFilename != expectedOutputFilename {
		t.Error(
			"For", "GetCSharpBuildCommand",
			"expected", expectedOutputFilename, "got",
			commandToTest.OutputFilename,
		)
	}

	const expectedSourceFiles = "./testsrc/a.cs ./testsrc/b.cs ./testsrc/c.cs"
	actualSourceFiles := strings.Join(commandToTest.SourceFiles, " ")

	if actualSourceFiles != expectedSourceFiles {
		t.Error(
			"For", "GetCSharpBuildCommand",
			"expected", expectedSourceFiles, "got",
			actualSourceFiles,
		)
	}

	const expectedBuildTarget = "-target:exe"

	if commandToTest.BuildTarget != expectedBuildTarget {
		t.Error(
			"For", "GetCSharpBuildCommand",
			"expected", expectedBuildTarget, "got",
			commandToTest.BuildTarget,
		)
	}

	const expectedReferences = "-r:Library.x,Library.y"

	if commandToTest.References != expectedReferences {
		t.Error(
			"For", "GetCSharpBuildCommand",
			"expected", expectedReferences, "got",
			commandToTest.References,
		)
	}

	const expectedLibraryPath = "-lib:./lib/lib1,./lib/lib2"

	if commandToTest.LibraryPath != expectedLibraryPath {
		t.Error(
			"For", "GetCSharpBuildCommand",
			"expected", expectedLibraryPath, "got",
			commandToTest.LibraryPath,
		)
	}

	const expectedPackageList = "-pkg:package1,package2"

	if commandToTest.PackageList != expectedPackageList {
		t.Error(
			"For", "GetCSharpBuildCommand",
			"expected", expectedPackageList, "got",
			commandToTest.PackageList,
		)
	}

	fmt.Println(commandToTest.ReferencePaths)
	actualReferencePaths := commandToTest.ReferencePaths

	if actualReferencePaths[0].Name != "Library.x" ||
		actualReferencePaths[0].Path != "./lib/Library.x/" {
		t.Error(
			"For", "GetCSharpBuildCommand",
			"expected", "Library.x ./lib/Library.x/", "got",
			actualReferencePaths[0].Name+","+actualReferencePaths[0].Path,
		)
	}

	if actualReferencePaths[1].Name != "Library.y" ||
		actualReferencePaths[1].Path != "./lib/Library.y/" {
		t.Error(
			"For", "GetCSharpBuildCommand",
			"expected", "Library.y ./lib/Library.y/", "got",
			actualReferencePaths[1].Name+","+actualReferencePaths[1].Path,
		)
	}
}
