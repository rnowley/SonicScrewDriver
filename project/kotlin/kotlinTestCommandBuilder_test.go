package kotlin

import (
	"strings"
	"testing"
)

func TestGetKotlinTestBuildCommand(t *testing.T) {
	// Arrange

	var configuration KotlinProject
	configuration.Name = "Test Kotlin Project"
	configuration.Version = "1.0.0"
	configuration.Description = "A project for unit testing."
	configuration.Language = "kotlin"
	configuration.DestinationDirectory = "./build/"
	configuration.JarFile = "test.jar"
	configuration.SourceDirectory = "./src/main/"
	configuration.OutputFilename = "test.jar"
	configuration.ClassPath = []string{"./lib/a.jar", "./lib/b.jar"}
	configuration.SourceFiles = []string{"a.kt", "b.kt", "c.kt"}
	configuration.BuildTarget = "executable"
	configuration.RunArguments = []string{"arg1", "arg2"}

	var testConfiguration KotlinTests
	testConfiguration.SourceFiles = []string{"testA.kt", "testB.kt"}
	testConfiguration.DestinationDirectory = "./testbuild/"
	testConfiguration.SourceDirectory = "./src/test/"
	testConfiguration.OutputFilename = "unittests.jar"
	testConfiguration.ClassPath = []string{"./lib/c.jar", "./lib/d.jar"}
	testConfiguration.TestRunner = "Test.Runner.Console"
	testConfiguration.RunArguments = []string{"mainTest"}

	configuration.TestProject = testConfiguration

	// Act
	commandToTest := GetKotlincTestBuildCommand(configuration)

	// Assert

	const expectedBuildTarget = "executable"

	if commandToTest.BuildTarget != expectedBuildTarget {
		t.Error(
			"For", "TestGetKotlinTestBuildCommand",
			"expected", expectedBuildTarget, "got",
			commandToTest.BuildTarget,
		)
	}

	const expectedDestinationDirectory = "./testbuild/"

	if commandToTest.DestinationDirectory != expectedDestinationDirectory {
		t.Error(
			"For", "TestGetKotlinTestBuildCommand",
			"expected", expectedDestinationDirectory, "got",
			commandToTest.DestinationDirectory,
		)
	}

	const expectedOutputFilename = "unittests.jar"

	if commandToTest.OutputFilename != expectedOutputFilename {
		t.Error(
			"For", "TestGetKotlinTestBuildCommand",
			"expected", expectedOutputFilename, "got",
			commandToTest.OutputFilename,
		)
	}

	const expectedClassPath = ". ./lib/a.jar ./lib/b.jar " +
		"./lib/c.jar ./lib/d.jar"
	actualClassPath := strings.Join(commandToTest.ClassPath, " ")

	if actualClassPath != expectedClassPath {
		t.Error(
			"For", "TestGetKotlinTestBuildCommand",
			"expected", expectedClassPath, "got",
			actualClassPath,
		)
	}

	const expectedSourceFileList = "./src/test/testA.kt ./src/test/testB.kt"
	actualSourceFileList := strings.Join(commandToTest.SourceFiles, " ")

	if actualSourceFileList != expectedSourceFileList {
		t.Error(
			"For", "TestGetKotlinTestBuildCommand",
			"expected", expectedSourceFileList, "got",
			actualSourceFileList,
		)
	}

}

func TestExtractKotlinTestSourceFileListNoSourceFiles(t *testing.T) {
	// Arrange

	var configuration KotlinProject
	configuration.TestProject.SourceFiles = []string{}

	// Act
	fileList := ExtractTestSourceFileList(configuration, "./src/test/")

	// Assert
	const expectedSourceFileList = ""
	actualSourceFileList := strings.Join(fileList, " ")

	if actualSourceFileList != expectedSourceFileList {
		t.Error(
			"For", "TestExtractKotlinTestSourceFileListNoSourceFiles",
			"expected", expectedSourceFileList, "got",
			actualSourceFileList,
		)
	}

}

func TestExtractKotlinTestSourceFileListOneSourceFile(t *testing.T) {
	// Arrange

	var configuration KotlinProject
	configuration.TestProject.SourceFiles = []string{"testA.kt"}

	// Act
	fileList := ExtractTestSourceFileList(configuration, "./src/test/")

	// Assert
	const expectedSourceFileList = "./src/test/testA.kt"
	actualSourceFileList := strings.Join(fileList, " ")

	if actualSourceFileList != expectedSourceFileList {
		t.Error(
			"For", "TestExtractKotlinTestSourceFileListNoSourceFiles",
			"expected", expectedSourceFileList, "got",
			actualSourceFileList,
		)
	}

}

func TestExtractKotlinTestSourceFileListThreeSourceFiles(t *testing.T) {
	// Arrange

	var configuration KotlinProject
	configuration.TestProject.SourceFiles = []string{"testA.kt", "testB.kt", "testC.kt"}

	// Act
	fileList := ExtractTestSourceFileList(configuration, "./src/test/")

	// Assert
	const expectedSourceFileList = "./src/test/testA.kt ./src/test/testB.kt ./src/test/testC.kt"
	actualSourceFileList := strings.Join(fileList, " ")

	if actualSourceFileList != expectedSourceFileList {
		t.Error(
			"For", "TestExtractKotlinTestSourceFileListNoSourceFiles",
			"expected", expectedSourceFileList, "got",
			actualSourceFileList,
		)
	}

}
