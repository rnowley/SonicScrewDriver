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

func TestCSharpExtractSourceFileListNoFiles(t *testing.T) {
	// Arrange
	var configuration CSharpProject

	// Act
	fileList := ExtractSourceFileList(configuration, "./src/")

	// Assert
	const expectedSourceFiles = ""
	actualSourceFiles := strings.Join(fileList, " ")

	if actualSourceFiles != expectedSourceFiles {
		t.Error(
			"For", "TestCSharpExtractSourceFileListNoFiles",
			"expected", expectedSourceFiles, "got",
			actualSourceFiles,
		)
	}
}

func TestCSharpExtractSourceFileListOneFile(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.SourceFiles = []string{"dir1/a.cs"}

	// Act
	fileList := ExtractSourceFileList(configuration, "./src/")

	// Assert
	const expectedSourceFiles = "./src/dir1/a.cs"
	actualSourceFiles := strings.Join(fileList, " ")

	if actualSourceFiles != expectedSourceFiles {
		t.Error(
			"For", "TestCSharpExtractSourceFileListNoFiles",
			"expected", expectedSourceFiles, "got",
			actualSourceFiles,
		)
	}
}

func TestCSharpExtractSourceFileListThreeFiles(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.SourceFiles = []string{"dir1/a.cs", "dir1/b.cs", "dir2/c.cs"}

	// Act
	fileList := ExtractSourceFileList(configuration, "./src/")

	// Assert
	const expectedSourceFiles = "./src/dir1/a.cs ./src/dir1/b.cs ./src/dir2/c.cs"
	actualSourceFiles := strings.Join(fileList, " ")

	if actualSourceFiles != expectedSourceFiles {
		t.Error(
			"For", "TestCSharpExtractSourceFileListNoFiles",
			"expected", expectedSourceFiles, "got",
			actualSourceFiles,
		)
	}
}

func TestCSharpExtractBuildTargetInvalid(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.BuildTarget = "invalid"

	// Act
	actualBuildTarget := ExtractBuildTarget(configuration)

	// Assert
	const expectedBuildTarget = "-target:exe"

	if actualBuildTarget != expectedBuildTarget {
		t.Error(
			"For", "TestCSharpExtractBuildTargetInvalid",
			"expected", expectedBuildTarget, "got",
			actualBuildTarget,
		)
	}
}

func TestCSharpExtractBuildTargetExe(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.BuildTarget = "exe"

	// Act
	actualBuildTarget := ExtractBuildTarget(configuration)

	// Assert
	const expectedBuildTarget = "-target:exe"

	if actualBuildTarget != expectedBuildTarget {
		t.Error(
			"For", "TestCSharpExtractBuildTargetExe",
			"expected", expectedBuildTarget, "got",
			actualBuildTarget,
		)
	}
}

func TestCSharpExtractBuildTargetLibrary(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.BuildTarget = "library"

	// Act
	actualBuildTarget := ExtractBuildTarget(configuration)

	// Assert
	const expectedBuildTarget = "-target:library"

	if actualBuildTarget != expectedBuildTarget {
		t.Error(
			"For", "TestCSharpExtractBuildTargetLibrary",
			"expected", expectedBuildTarget, "got",
			actualBuildTarget,
		)
	}
}

func TestCSharpExtractBuildTargetModule(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.BuildTarget = "module"

	// Act
	actualBuildTarget := ExtractBuildTarget(configuration)

	// Assert
	const expectedBuildTarget = "-target:module"

	if actualBuildTarget != expectedBuildTarget {
		t.Error(
			"For", "TestCSharpExtractBuildTargetModule",
			"expected", expectedBuildTarget, "got",
			actualBuildTarget,
		)
	}
}

func TestCSharpExtractBuildTargetWinexe(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.BuildTarget = "winexe"

	// Act
	actualBuildTarget := ExtractBuildTarget(configuration)

	// Assert
	const expectedBuildTarget = "-target:winexe"

	if actualBuildTarget != expectedBuildTarget {
		t.Error(
			"For", "TestCSharpExtractBuildTargetWinexe",
			"expected", expectedBuildTarget, "got",
			actualBuildTarget,
		)
	}
}

func TestCSharpExtractLibraryPathNoLibraryPaths(t *testing.T) {
	// Arrange
	var configuration CSharpProject

	// Act
	actualLibraryPaths := ExtractLibraryPath(configuration)

	// Assert
	const expectedLibraryPaths = ""

	if actualLibraryPaths != expectedLibraryPaths {
		t.Error(
			"For", "TestCSharpExtractLibraryPathNoLibraryPaths",
			"expected", expectedLibraryPaths, "got",
			actualLibraryPaths,
		)
	}
}

func TestCSharpExtractLibraryPathOneLibrary(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.LibraryPath = []string{"./lib/lib1/"}

	// Act
	actualLibraryPaths := ExtractLibraryPath(configuration)

	// Assert
	const expectedLibraryPaths = "-lib:./lib/lib1/"

	if actualLibraryPaths != expectedLibraryPaths {
		t.Error(
			"For", "TestCSharpExtractLibraryPathOneLibrary",
			"expected", expectedLibraryPaths, "got",
			actualLibraryPaths,
		)
	}
}

func TestCSharpExtractLibraryPathThreeLibraries(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.LibraryPath = []string{"./lib/lib1/", "./lib/lib2/", "./lib/lib3/"}

	// Act
	actualLibraryPaths := ExtractLibraryPath(configuration)

	// Assert
	const expectedLibraryPaths = "-lib:./lib/lib1/,./lib/lib2/,./lib/lib3/"

	if actualLibraryPaths != expectedLibraryPaths {
		t.Error(
			"For", "TestCSharpExtractLibraryPathThreeLibraries",
			"expected", expectedLibraryPaths, "got",
			actualLibraryPaths,
		)
	}
}

func TestCSharpExtractPackageListNoPackages(t *testing.T) {
	// Arrange
	var configuration CSharpProject

	// Act
	actualPackageList := ExtractPackageList(configuration)

	// Assert
	const expectedPackageList = ""

	if actualPackageList != expectedPackageList {
		t.Error(
			"For", "TestCSharpExtractPackageListNoPackages",
			"expected", expectedPackageList, "got",
			actualPackageList,
		)
	}
}

func TestCSharpExtractPackageListOnePackage(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.PackageList = []string{"package1"}

	// Act
	actualPackageList := ExtractPackageList(configuration)

	// Assert
	const expectedPackageList = "-pkg:package1"

	if actualPackageList != expectedPackageList {
		t.Error(
			"For", "TestCSharpExtractPackageListOnePackage",
			"expected", expectedPackageList, "got",
			actualPackageList,
		)
	}
}

func TestCSharpExtractPackageListThreePackages(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.PackageList = []string{"package1", "package2", "package3"}

	// Act
	actualPackageList := ExtractPackageList(configuration)

	// Assert
	const expectedPackageList = "-pkg:package1,package2,package3"

	if actualPackageList != expectedPackageList {
		t.Error(
			"For", "TestCSharpExtractPackageListThreePackages",
			"expected", expectedPackageList, "got",
			actualPackageList,
		)
	}
}
