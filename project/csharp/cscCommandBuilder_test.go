package csharp

import (
	"strings"
	"testing"
)

func TestGetCscBuildCommand(t *testing.T) {
	// Arrange

	var configuration CSharpProject
	configuration.Name = "Test C# Project"
	configuration.Version = "1.0.0"
	configuration.Description = "A project for unit testing."
	configuration.Language = "csharp:ms"
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
	commandToTest := GetCscBuildCommand(configuration)

	// Assert

	const expectedSourceDirectory = "./testsrc/"

	if commandToTest.SourceDirectory != expectedSourceDirectory {
		t.Error(
			"For", "GetCscBuildCommand",
			"expected", expectedSourceDirectory, "got",
			commandToTest.SourceDirectory,
		)
	}

	const expectedDestinationDirectory = "./testbuild/"

	if commandToTest.DestinationDirectory != expectedDestinationDirectory {
		t.Error(
			"For", "GetCscBuildCommand",
			"expected", expectedDestinationDirectory, "got",
			commandToTest.DestinationDirectory,
		)
	}

	const expectedOutputFilename = "/out:./testbuild/test"

	if commandToTest.OutputFilename != expectedOutputFilename {
		t.Error(
			"For", "GetCscBuildCommand",
			"expected", expectedOutputFilename, "got",
			commandToTest.OutputFilename,
		)
	}

	const expectedSourceFiles = "./testsrc/a.cs ./testsrc/b.cs ./testsrc/c.cs"
	actualSourceFiles := strings.Join(commandToTest.SourceFiles, " ")

	if actualSourceFiles != expectedSourceFiles {
		t.Error(
			"For", "GetCscBuildCommand",
			"expected", expectedSourceFiles, "got",
			actualSourceFiles,
		)
	}

	const expectedBuildTarget = "/target:exe"

	if commandToTest.BuildTarget != expectedBuildTarget {
		t.Error(
			"For", "GetCscBuildCommand",
			"expected", expectedBuildTarget, "got",
			commandToTest.BuildTarget,
		)
	}

	const expectedReferences = "/r:Library.x,Library.y"

	if commandToTest.References != expectedReferences {
		t.Error(
			"For", "GetCscBuildCommand",
			"expected", expectedReferences, "got",
			commandToTest.References,
		)
	}

	const expectedLibraryPath = "/lib:./lib/lib1,./lib/lib2"

	if commandToTest.LibraryPath != expectedLibraryPath {
		t.Error(
			"For", "GetCscBuildCommand",
			"expected", expectedLibraryPath, "got",
			commandToTest.LibraryPath,
		)
	}

	actualReferencePaths := commandToTest.ReferencePaths

	if actualReferencePaths[0].Name != "Library.x" ||
		actualReferencePaths[0].Path != "./lib/Library.x/" {
		t.Error(
			"For", "GetCscBuildCommand",
			"expected", "Library.x ./lib/Library.x/", "got",
			actualReferencePaths[0].Name+","+actualReferencePaths[0].Path,
		)
	}

	if actualReferencePaths[1].Name != "Library.y" ||
		actualReferencePaths[1].Path != "./lib/Library.y/" {
		t.Error(
			"For", "GetCscBuildCommand",
			"expected", "Library.y ./lib/Library.y/", "got",
			actualReferencePaths[1].Name+","+actualReferencePaths[1].Path,
		)
	}
}

func TestCscExtractBuildTargetInvalid(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.BuildTarget = "invalid"

	// Act
	actualBuildTarget := ExtractCscBuildTarget(configuration)

	// Assert
	const expectedBuildTarget = "/target:exe"

	if actualBuildTarget != expectedBuildTarget {
		t.Error(
			"For", "TestCscExtractBuildTargetInvalid",
			"expected", expectedBuildTarget, "got",
			actualBuildTarget,
		)
	}
}

func TestCscExtractBuildTargetExe(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.BuildTarget = "exe"

	// Act
	actualBuildTarget := ExtractCscBuildTarget(configuration)

	// Assert
	const expectedBuildTarget = "/target:exe"

	if actualBuildTarget != expectedBuildTarget {
		t.Error(
			"For", "TestCscExtractBuildTargetExe",
			"expected", expectedBuildTarget, "got",
			actualBuildTarget,
		)
	}
}

func TestCscExtractBuildTargetLibrary(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.BuildTarget = "library"

	// Act
	actualBuildTarget := ExtractCscBuildTarget(configuration)

	// Assert
	const expectedBuildTarget = "/target:library"

	if actualBuildTarget != expectedBuildTarget {
		t.Error(
			"For", "TestCscExtractBuildTargetLibrary",
			"expected", expectedBuildTarget, "got",
			actualBuildTarget,
		)
	}
}

func TestCscExtractBuildTargetModule(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.BuildTarget = "module"

	// Act
	actualBuildTarget := ExtractCscBuildTarget(configuration)

	// Assert
	const expectedBuildTarget = "/target:module"

	if actualBuildTarget != expectedBuildTarget {
		t.Error(
			"For", "TestCscExtractBuildTargetModule",
			"expected", expectedBuildTarget, "got",
			actualBuildTarget,
		)
	}
}

func TestCscExtractBuildTargetWinexe(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.BuildTarget = "winexe"

	// Act
	actualBuildTarget := ExtractCscBuildTarget(configuration)

	// Assert
	const expectedBuildTarget = "/target:winexe"

	if actualBuildTarget != expectedBuildTarget {
		t.Error(
			"For", "TestCscExtractBuildTargetWinexe",
			"expected", expectedBuildTarget, "got",
			actualBuildTarget,
		)
	}
}

func TestCscExtractBuildTargetWinmdobj(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.BuildTarget = "winmdobj"

	// Act
	actualBuildTarget := ExtractCscBuildTarget(configuration)

	// Assert
	const expectedBuildTarget = "/target:winmdobj"

	if actualBuildTarget != expectedBuildTarget {
		t.Error(
			"For", "TestCscExtractBuildTargetWinmdobj",
			"expected", expectedBuildTarget, "got",
			actualBuildTarget,
		)
	}
}

func TestCcsExtractLibraryPathNoLibraryPaths(t *testing.T) {
	// Arrange
	var configuration CSharpProject

	// Act
	actualLibraryPaths := ExtractCscLibraryPath(configuration)

	// Assert
	const expectedLibraryPaths = ""

	if actualLibraryPaths != expectedLibraryPaths {
		t.Error(
			"For", "TestCcsExtractLibraryPathNoLibraryPaths",
			"expected", expectedLibraryPaths, "got",
			actualLibraryPaths,
		)
	}
}

func TestCscExtractLibraryPathOneLibrary(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.LibraryPath = []string{"./lib/lib1/"}

	// Act
	actualLibraryPaths := ExtractCscLibraryPath(configuration)

	// Assert
	const expectedLibraryPaths = "/lib:./lib/lib1/"

	if actualLibraryPaths != expectedLibraryPaths {
		t.Error(
			"For", "TestCscExtractLibraryPathOneLibrary",
			"expected", expectedLibraryPaths, "got",
			actualLibraryPaths,
		)
	}
}

func TestCscExtractLibraryPathThreeLibraries(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.LibraryPath = []string{"./lib/lib1/", "./lib/lib2/", "./lib/lib3/"}

	// Act
	actualLibraryPaths := ExtractCscLibraryPath(configuration)

	// Assert
	const expectedLibraryPaths = "/lib:./lib/lib1/,./lib/lib2/,./lib/lib3/"

	if actualLibraryPaths != expectedLibraryPaths {
		t.Error(
			"For", "TestCscExtractLibraryPathThreeLibraries",
			"expected", expectedLibraryPaths, "got",
			actualLibraryPaths,
		)
	}
}

func TestCscExtractReferencesNoReferences(t *testing.T) {
	// Arrange
	var configuration CSharpProject

	// Act
	actualReferences := ExtractCscReferences(configuration)

	// Assert
	const expectedReferences = ""

	if actualReferences != expectedReferences {
		t.Error(
			"For", "TestCscExtractReferencesNoReferences",
			"expected", expectedReferences, "got",
			actualReferences,
		)
	}
}

func TestCscExtractReferencesOneReference(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.References = []Reference{Reference{Name: "reference1", Path: "./lib/reference1/"}}

	// Act
	actualReferences := ExtractCscReferences(configuration)

	// Assert
	const expectedReferences = "/r:reference1"

	if actualReferences != expectedReferences {
		t.Error(
			"For", "TestCscExtractReferencesOneReference",
			"expected", expectedReferences, "got",
			actualReferences,
		)
	}
}

func TestCscExtractReferencesThreeReferences(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.References = []Reference{Reference{Name: "reference1", Path: "./lib/reference1/"},
		Reference{Name: "reference2", Path: "./lib/reference2/"},
		Reference{Name: "reference3", Path: "./lib/reference3/"}}

	// Act
	actualReferences := ExtractCscReferences(configuration)

	// Assert
	const expectedReferences = "/r:reference1,reference2,reference3"

	if actualReferences != expectedReferences {
		t.Error(
			"For", "TestCscExtractReferencesThreeReferences",
			"expected", expectedReferences, "got",
			actualReferences,
		)
	}
}

func TestCscExtractReferencePathsNoReferences(t *testing.T) {
	// Arrange
	var configuration CSharpProject

	// Act
	actualReferencesCount := len(ExtractCscReferencePaths(configuration))

	// Assert
	const expectedReferencesCount = 0

	if actualReferencesCount != expectedReferencesCount {
		t.Error(
			"For", "TestCscExtractReferencePathsNoReferences",
			"expected", expectedReferencesCount, "got",
			actualReferencesCount,
		)
	}
}

func TestCscExtractReferencePathsOneReference(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.References = []Reference{Reference{Name: "reference1", Path: "./lib/reference1/"}}

	// Act
	actualReferences := ExtractCscReferencePaths(configuration)

	// Assert

	if actualReferences[0].Name != "reference1" ||
		actualReferences[0].Path != "./lib/reference1/" {
		t.Error(
			"For", "TestCscExtractReferencePathsOneReference",
			"expected", "reference1, "+"./lib/reference1/", "got",
			actualReferences[0].Name+", "+actualReferences[0].Path,
		)
	}
}

func TestCscExtractReferencePathsThreeReferences(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.References = []Reference{Reference{Name: "reference1", Path: "./lib/reference1/"},
		Reference{Name: "reference2", Path: "./lib/reference2/"},
		Reference{Name: "reference3", Path: "./lib/reference3/"}}

	// Act
	actualReferences := ExtractCscReferencePaths(configuration)

	// Assert

	if actualReferences[0].Name != "reference1" ||
		actualReferences[0].Path != "./lib/reference1/" {
		t.Error(
			"For", "TestCscExtractReferencePathsThreeReferences",
			"expected", "reference1, "+"./lib/reference1/", "got",
			actualReferences[0].Name+", "+actualReferences[0].Path,
		)
	}

	if actualReferences[1].Name != "reference2" ||
		actualReferences[1].Path != "./lib/reference2/" {
		t.Error(
			"For", "TestCSharpExtractReferencePathsThreeReferences",
			"expected", "reference2, "+"./lib/reference2/", "got",
			actualReferences[1].Name+", "+actualReferences[1].Path,
		)
	}

	if actualReferences[2].Name != "reference3" ||
		actualReferences[2].Path != "./lib/reference3/" {
		t.Error(
			"For", "TestCSharpExtractReferencePathsThreeReferences",
			"expected", "reference3, "+"./lib/reference3/", "got",
			actualReferences[2].Name+", "+actualReferences[2].Path,
		)
	}
}
