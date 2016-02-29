package csharp

import (
	"strings"
	"testing"
)

func TestGetCSharpTestBuildCommand(t *testing.T) {
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

	var testConfiguration CSharpTests
	testConfiguration.TestRunner = "Test.Runner.Console"
	testConfiguration.DestinationDirectory = "./testbuild/"
	testConfiguration.References = []Reference{Reference{Name: "Library.a",
		Path: "./lib/Library.a/"}, Reference{Name: "Library.b",
		Path: "./lib/Library.b/"}}
	testConfiguration.OutputFilename = "unitTests"
	testConfiguration.SourceFiles = []string{"atest.cs", "btest.cs", "ctest.cs"}
	testConfiguration.PackageList = []string{"package3", "package4"}
	testConfiguration.RunArguments = []string{"arg3", "arg4"}
	testConfiguration.LibraryPath = []string{"./lib/Library.a/", "./lib/Library.b/"}

	configuration.TestProject = testConfiguration

	// Act
	commandToTest := GetCSharpTestBuildCommand(configuration)

	// Assert

	const expectedSourceDirectory = "./testsrc/"

	if commandToTest.SourceDirectory != expectedSourceDirectory {
		t.Error(
			"For", "TestGetCSharpTestBuildCommand",
			"expected", expectedSourceDirectory, "got",
			commandToTest.SourceDirectory,
		)
	}

	const expectedDestinationDirectory = "./testbuild/"

	if commandToTest.DestinationDirectory != expectedDestinationDirectory {
		t.Error(
			"For", "TestGetCSharpTestBuildCommand",
			"expected", expectedDestinationDirectory, "got",
			commandToTest.DestinationDirectory,
		)
	}

	const expectedOutputFilename = "-out:./testbuild/unitTests"

	if commandToTest.OutputFilename != expectedOutputFilename {
		t.Error(
			"For", "TestGetCSharpTestBuildCommand",
			"expected", expectedOutputFilename, "got",
			commandToTest.OutputFilename,
		)
	}

	const expectedBuildTarget = "-target:library"

	if commandToTest.BuildTarget != expectedBuildTarget {
		t.Error(
			"For", "TestGetCSharpTestBuildCommand",
			"expected", expectedBuildTarget, "got",
			commandToTest.BuildTarget,
		)
	}

	const expectedSourceFiles = "./testsrc/atest.cs ./testsrc/btest.cs ./testsrc/ctest.cs"
	actualSourceFiles := strings.Join(commandToTest.SourceFiles, " ")

	if actualSourceFiles != expectedSourceFiles {
		t.Error(
			"For", "TestGetCSharpTestBuildCommand",
			"expected", expectedSourceFiles, "got",
			actualSourceFiles,
		)
	}

	const expectedReferences = "-r:Library.x,Library.y,Library.a,Library.b"

	if commandToTest.References != expectedReferences {
		t.Error(
			"For", "TestGetCSharpTestBuildCommand",
			"expected", expectedReferences, "got",
			commandToTest.References,
		)
	}

	const expectedLibraryPath = "-lib:./lib/Library.a/,./lib/Library.b/"

	if commandToTest.LibraryPath != expectedLibraryPath {
		t.Error(
			"For", "TestGetCSharpTestBuildCommand",
			"expected", expectedLibraryPath, "got",
			commandToTest.LibraryPath,
		)
	}

	const expectedPackageList = "-pkg:package3,package4"

	if commandToTest.PackageList != expectedPackageList {
		t.Error(
			"For", "TestGetCSharpTestBuildCommand",
			"expected", expectedPackageList, "got",
			commandToTest.PackageList,
		)
	}

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

	if actualReferencePaths[2].Name != "Library.a" ||
		actualReferencePaths[2].Path != "./lib/Library.a/" {
		t.Error(
			"For", "GetCSharpBuildCommand",
			"expected", "Library.a ./lib/Library.a/", "got",
			actualReferencePaths[2].Name+","+actualReferencePaths[0].Path,
		)
	}

	if actualReferencePaths[3].Name != "Library.b" ||
		actualReferencePaths[3].Path != "./lib/Library.b/" {
		t.Error(
			"For", "GetCSharpBuildCommand",
			"expected", "Library.b ./lib/Library.b/", "got",
			actualReferencePaths[3].Name+","+actualReferencePaths[1].Path,
		)
	}
}

func TestCSharpExtractTestSourceFileListNoFiles(t *testing.T) {
	// Arrange
	var configuration CSharpProject

	// Act
	fileList := ExtractTestSourceFileList(configuration, "./src/")

	// Assert
	const expectedSourceFiles = ""
	actualSourceFiles := strings.Join(fileList, " ")

	if actualSourceFiles != expectedSourceFiles {
		t.Error(
			"For", "TestCSharpExtractTestSourceFileListNoFiles",
			"expected", expectedSourceFiles, "got",
			actualSourceFiles,
		)
	}
}

func TestCSharpExtractTestSourceFileListOneFile(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	var testConfiguration CSharpTests
	testConfiguration.SourceFiles = []string{"dir1/a.cs"}
	configuration.TestProject = testConfiguration

	// Act
	fileList := ExtractTestSourceFileList(configuration, "./src/")

	// Assert
	const expectedSourceFiles = "./src/dir1/a.cs"
	actualSourceFiles := strings.Join(fileList, " ")

	if actualSourceFiles != expectedSourceFiles {
		t.Error(
			"For", "TestCSharpExtractTestSourceFileListOneFile",
			"expected", expectedSourceFiles, "got",
			actualSourceFiles,
		)
	}
}

func TestCSharpExtractTestSourceFileListThreeFiles(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	var testConfiguration CSharpTests
	testConfiguration.SourceFiles = []string{"dir1/a.cs", "dir1/b.cs", "dir2/c.cs"}
	configuration.TestProject = testConfiguration

	// Act
	fileList := ExtractTestSourceFileList(configuration, "./src/")

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

func TestExtractTestBuildTargetInvalid(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.BuildTarget = "invalid"

	// Act
	actualBuildTarget := ExtractTestBuildTarget(configuration)

	// Assert
	const expectedBuildTarget = "-target:library"

	if actualBuildTarget != expectedBuildTarget {
		t.Error(
			"For", "TestCSharpExtractBuildTargetInvalid",
			"expected", expectedBuildTarget, "got",
			actualBuildTarget,
		)
	}
}

func TestExtractTestBuildTargetExe(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.BuildTarget = "exe"

	// Act
	actualBuildTarget := ExtractTestBuildTarget(configuration)

	// Assert
	const expectedBuildTarget = "-target:library"

	if actualBuildTarget != expectedBuildTarget {
		t.Error(
			"For", "TestCSharpExtractBuildTargetExe",
			"expected", expectedBuildTarget, "got",
			actualBuildTarget,
		)
	}
}

func TestExtractTestBuildTargetLibrary(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.BuildTarget = "library"

	// Act
	actualBuildTarget := ExtractTestBuildTarget(configuration)

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

func TestExtractTestBuildTargetModule(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.BuildTarget = "module"

	// Act
	actualBuildTarget := ExtractTestBuildTarget(configuration)

	// Assert
	const expectedBuildTarget = "-target:library"

	if actualBuildTarget != expectedBuildTarget {
		t.Error(
			"For", "TestCSharpExtractBuildTargetModule",
			"expected", expectedBuildTarget, "got",
			actualBuildTarget,
		)
	}
}

func TestExtractTestBuildTargetWinexe(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.BuildTarget = "winexe"

	// Act
	actualBuildTarget := ExtractTestBuildTarget(configuration)

	// Assert
	const expectedBuildTarget = "-target:library"

	if actualBuildTarget != expectedBuildTarget {
		t.Error(
			"For", "TestCSharpExtractBuildTargetWinexe",
			"expected", expectedBuildTarget, "got",
			actualBuildTarget,
		)
	}
}

func TestExtractTestLibraryPathNoLibraryPaths(t *testing.T) {
	// Arrange
	var configuration CSharpProject

	// Act
	actualLibraryPaths := ExtractTestLibraryPath(configuration)

	// Assert
	const expectedLibraryPaths = ""

	if actualLibraryPaths != expectedLibraryPaths {
		t.Error(
			"For", "TestExtractTestLibraryPathNoLibraryPaths",
			"expected", expectedLibraryPaths, "got",
			actualLibraryPaths,
		)
	}
}

func TestExtractTestLibraryPathOneLibrary(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	var testConfiguration CSharpTests
	testConfiguration.LibraryPath = []string{"./lib/lib1/"}
	configuration.TestProject = testConfiguration

	// Act
	actualLibraryPaths := ExtractTestLibraryPath(configuration)

	// Assert
	const expectedLibraryPaths = "-lib:./lib/lib1/"

	if actualLibraryPaths != expectedLibraryPaths {
		t.Error(
			"For", "TestExtractTestLibraryPathOneLibrary",
			"expected", expectedLibraryPaths, "got",
			actualLibraryPaths,
		)
	}
}

func TestExtractTestLibraryPathThreeLibraries(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	var testConfiguration CSharpTests
	testConfiguration.LibraryPath = []string{"./lib/lib1/", "./lib/lib2/", "./lib/lib3/"}
	configuration.TestProject = testConfiguration

	// Act
	actualLibraryPaths := ExtractTestLibraryPath(configuration)

	// Assert
	const expectedLibraryPaths = "-lib:./lib/lib1/,./lib/lib2/,./lib/lib3/"

	if actualLibraryPaths != expectedLibraryPaths {
		t.Error(
			"For", "TestExtractTestLibraryPathThreeLibraries",
			"expected", expectedLibraryPaths, "got",
			actualLibraryPaths,
		)
	}
}

func TestCSharpExtractTestPackageListNoPackages(t *testing.T) {
	// Arrange
	var configuration CSharpProject

	// Act
	actualPackageList := ExtractTestPackageList(configuration)

	// Assert
	const expectedPackageList = ""

	if actualPackageList != expectedPackageList {
		t.Error(
			"For", "TestCSharpExtractTestPackageListListNoPackages",
			"expected", expectedPackageList, "got",
			actualPackageList,
		)
	}
}

func TestCSharpExtractTestPackageListOnePackage(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.TestProject.PackageList = []string{"package1"}

	// Act
	actualPackageList := ExtractTestPackageList(configuration)

	// Assert
	const expectedPackageList = "-pkg:package1"

	if actualPackageList != expectedPackageList {
		t.Error(
			"For", "TestCSharpExtractTestPackageListOnePackage",
			"expected", expectedPackageList, "got",
			actualPackageList,
		)
	}
}

func TestCSharpExtractTestPackageListListThreePackages(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.TestProject.PackageList = []string{"package1", "package2", "package3"}

	// Act
	actualPackageList := ExtractTestPackageList(configuration)

	// Assert
	const expectedPackageList = "-pkg:package1,package2,package3"

	if actualPackageList != expectedPackageList {
		t.Error(
			"For", "TestCSharpExtractTestPackageListListThreePackages",
			"expected", expectedPackageList, "got",
			actualPackageList,
		)
	}
}

func TestCSharpExtractTestReferencesNoReferences(t *testing.T) {
	// Arrange
	var configuration CSharpProject

	// Act
	actualReferences := ExtractTestReferences(configuration)

	// Assert
	const expectedReferences = ""

	if actualReferences != expectedReferences {
		t.Error(
			"For", "TestCSharpExtractTestReferencesNoReferences",
			"expected", expectedReferences, "got",
			actualReferences,
		)
	}
}

func TestCSharpExtractTestReferencesOneReference(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.TestProject.References = []Reference{Reference{Name: "reference1", Path: "./lib/reference1/"}}

	// Act
	actualReferences := ExtractTestReferences(configuration)

	// Assert
	const expectedReferences = "-r:reference1"

	if actualReferences != expectedReferences {
		t.Error(
			"For", "TestCSharpExtractTestReferencesOneReference",
			"expected", expectedReferences, "got",
			actualReferences,
		)
	}
}

func TestCSharpExtractTestReferencesThreeReferences(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.TestProject.References = []Reference{Reference{Name: "reference1", Path: "./lib/reference1/"},
		Reference{Name: "reference2", Path: "./lib/reference2/"},
		Reference{Name: "reference3", Path: "./lib/reference3/"}}

	// Act
	actualReferences := ExtractTestReferences(configuration)

	// Assert
	const expectedReferences = "-r:reference1,reference2,reference3"

	if actualReferences != expectedReferences {
		t.Error(
			"For", "TestCSharpExtractTestReferencesThreeReferences",
			"expected", expectedReferences, "got",
			actualReferences,
		)
	}
}

func TestCSharpExtractTestReferencePathsNoReferences(t *testing.T) {
	// Arrange
	var configuration CSharpProject

	// Act
	actualReferencesCount := len(ExtractTestReferencePaths(configuration))

	// Assert
	const expectedReferencesCount = 0

	if actualReferencesCount != expectedReferencesCount {
		t.Error(
			"For", "TestCSharpExtractTestReferencePathsNoReferences",
			"expected", expectedReferencesCount, "got",
			actualReferencesCount,
		)
	}
}

func TestCSharpExtractTestReferencePathsOneReference(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.TestProject.References = []Reference{Reference{Name: "reference1", Path: "./lib/reference1/"}}

	// Act
	actualReferences := ExtractTestReferencePaths(configuration)

	// Assert

	if actualReferences[0].Name != "reference1" ||
		actualReferences[0].Path != "./lib/reference1/" {
		t.Error(
			"For", "TestCSharpExtractTestReferencePathsOneReference",
			"expected", "reference1, "+"./lib/reference1/", "got",
			actualReferences[0].Name+", "+actualReferences[0].Path,
		)
	}
}

func TestCSharpExtractTestReferencePathsThreeReferences(t *testing.T) {
	// Arrange
	var configuration CSharpProject
	configuration.TestProject.References = []Reference{Reference{Name: "reference1", Path: "./lib/reference1/"},
		Reference{Name: "reference2", Path: "./lib/reference2/"},
		Reference{Name: "reference3", Path: "./lib/reference3/"}}

	// Act
	actualReferences := ExtractTestReferencePaths(configuration)

	// Assert

	if actualReferences[0].Name != "reference1" ||
		actualReferences[0].Path != "./lib/reference1/" {
		t.Error(
			"For", "TestCSharpExtractTestReferencePathsThreeReferences",
			"expected", "reference1, "+"./lib/reference1/", "got",
			actualReferences[0].Name+", "+actualReferences[0].Path,
		)
	}

	if actualReferences[1].Name != "reference2" ||
		actualReferences[1].Path != "./lib/reference2/" {
		t.Error(
			"For", "TestCSharpExtractTestReferencePathsThreeReferences",
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
