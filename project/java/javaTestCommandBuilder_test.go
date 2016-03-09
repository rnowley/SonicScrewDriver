package java

import (
	"strings"
	"testing"
)

func TestGetJavaTestBuildCommandDeprecationFalse(t *testing.T) {
	// Arrange

	var configuration JavaProject
	configuration.Name = "Test Java Project"
	configuration.Version = "1.0.0"
	configuration.Description = "A project for unit testing."
	configuration.Language = "java"
	configuration.DestinationDirectory = "./build/"
	configuration.ClassPath = []string{"./lib/a.jar", "./lib/b.jar"}
	configuration.SourceFiles = []string{"a.java", "b.java", "c.java"}
	configuration.SourceVersion = "1.7"
	configuration.JarFile = "test.jar"
	configuration.RunArguments = []string{"arg1", "arg2"}
	configuration.DebuggingInformation = []string{"all"}
	configuration.Target = "1.8"

	var testConfiguration JavaTests
	testConfiguration.SourceFiles = []string{"testA.java", "testB.java"}
	testConfiguration.TestFrameWork = "junit"
	testConfiguration.SourceDirectory = "./src/test/"
	testConfiguration.ClassPath = []string{"./lib/c.jar", "./lib/d.jar"}
	testConfiguration.DestinationDirectory = "./testbuild/"
	testConfiguration.MainClass = "Test.Runner.Console"
	testConfiguration.RunArguments = []string{"mainTest"}

	configuration.TestProject = testConfiguration

	// Act
	commandToTest := GetJavaTestBuildCommand(configuration, false)

	// Assert
	const expectedDestinationDirectory = "./testbuild/"

	if commandToTest.DestinationDirectory != expectedDestinationDirectory {
		t.Error(
			"For", "TestGetJavaTestBuildCommandDeprecationFalse",
			"expected", expectedDestinationDirectory, "got",
			commandToTest.DestinationDirectory,
		)
	}

	const expectedClassPath = ". ./lib/a.jar ./lib/b.jar " +
		"./lib/c.jar ./lib/d.jar"
	actualClassPath := strings.Join(commandToTest.ClassPath, " ")

	if actualClassPath != expectedClassPath {
		t.Error(
			"For", "TestGetJavaTestBuildCommandDeprecationFalse",
			"expected", expectedClassPath, "got",
			actualClassPath,
		)
	}

	const expectedDeprecationValue = false

	if commandToTest.Deprecation != expectedDeprecationValue {
		t.Error(
			"For", "TestGetJavaTestBuildCommandDeprecationFalse",
			"expected", expectedDeprecationValue, "got",
			commandToTest.Deprecation,
		)
	}

	const expectedTarget = "1.8"

	if commandToTest.Target != expectedTarget {
		t.Error(
			"For", "TestGetJavaTestBuildCommandDeprecationFalse",
			"expected", expectedTarget, "got",
			commandToTest.Target,
		)
	}

	const expectedSourceFileList = "./src/test/testA.java ./src/test/testB.java"
	actualSourceFileList := strings.Join(commandToTest.SourceFiles, " ")

	if actualSourceFileList != expectedSourceFileList {
		t.Error(
			"For", "TestGetJavaTestBuildCommandDeprecationFalse",
			"expected", expectedSourceFileList, "got",
			actualSourceFileList,
		)
	}
}

func TestGetJavaTestBuildCommandDeprecationTrue(t *testing.T) {
	// Arrange

	var configuration JavaProject
	configuration.Name = "Test Java Project"
	configuration.Version = "1.0.0"
	configuration.Description = "A project for unit testing."
	configuration.Language = "java"
	configuration.DestinationDirectory = "./build/"
	configuration.ClassPath = []string{"./lib/a.jar", "./lib/b.jar"}
	configuration.SourceFiles = []string{"a.java", "b.java", "c.java"}
	configuration.SourceVersion = "1.7"
	configuration.JarFile = "test.jar"
	configuration.RunArguments = []string{"arg1", "arg2"}
	configuration.DebuggingInformation = []string{"all"}

	var testConfiguration JavaTests
	testConfiguration.SourceFiles = []string{"testA.java", "testB.java"}
	testConfiguration.TestFrameWork = "junit"
	testConfiguration.SourceDirectory = "./src/test"
	testConfiguration.ClassPath = []string{"./lib/c.jar", "./lib/d.jar"}
	testConfiguration.DestinationDirectory = "./testbuild/"
	testConfiguration.MainClass = "Test.Runner.Console"
	testConfiguration.RunArguments = []string{"mainTest"}

	configuration.TestProject = testConfiguration

	// Act
	commandToTest := GetJavaTestBuildCommand(configuration, true)

	// Assert

	const expectedDeprecationValue = true

	if commandToTest.Deprecation != expectedDeprecationValue {
		t.Error(
			"For", "GetJavaTestBuildCommand",
			"expected", expectedDeprecationValue, "got",
			commandToTest.Deprecation,
		)
	}
}

func TestExtractTestSourceFileListNoSourceFiles(t *testing.T) {
	// Arrange

	var configuration JavaProject
	configuration.TestProject.SourceFiles = []string{}

	// Act
	fileList := ExtractTestSourceFileList(configuration, "./src/test/")

	// Assert
	const expectedSourceFileList = ""
	actualSourceFileList := strings.Join(fileList, " ")

	if actualSourceFileList != expectedSourceFileList {
		t.Error(
			"For", "ExtractSourceFileList",
			"expected", expectedSourceFileList, "got",
			actualSourceFileList,
		)
	}

}

func TestExtractTestSourceFileListOneSourceFile(t *testing.T) {
	// Arrange

	var configuration JavaProject
	configuration.TestProject.SourceFiles = []string{"testA.java"}

	// Act
	fileList := ExtractTestSourceFileList(configuration, "./src/test/")

	// Assert
	const expectedSourceFileList = "./src/test/testA.java"
	actualSourceFileList := strings.Join(fileList, " ")

	if actualSourceFileList != expectedSourceFileList {
		t.Error(
			"For", "ExtractSourceFileList",
			"expected", expectedSourceFileList, "got",
			actualSourceFileList,
		)
	}

}

func TestExtractTestSourceFileListThreeSourceFiles(t *testing.T) {
	// Arrange

	var configuration JavaProject
	configuration.TestProject.SourceFiles = []string{"testA.java", "testB.java", "testC.java"}

	// Act
	fileList := ExtractTestSourceFileList(configuration, "./src/test/")

	// Assert
	const expectedSourceFileList = "./src/test/testA.java ./src/test/testB.java ./src/test/testC.java"
	actualSourceFileList := strings.Join(fileList, " ")

	if actualSourceFileList != expectedSourceFileList {
		t.Error(
			"For", "ExtractSourceFileList",
			"expected", expectedSourceFileList, "got",
			actualSourceFileList,
		)
	}

}

func TestExtractTestDebuggingInformationNoValueProvided(t *testing.T) {
	// Arrange

	var configuration JavaProject
	configuration.DebuggingInformation = []string{}

	// Act
	actualDebuggingInformation := ExtractTestDebuggingInformation(configuration)

	// Assert
	const expectedDebuggingInformation = ""

	if actualDebuggingInformation != expectedDebuggingInformation {
		t.Error(
			"For", "ExtractDebuggingInformation",
			"expected", expectedDebuggingInformation,
			"got", actualDebuggingInformation,
		)
	}

}

func TestExtractTestDebuggingInformationAllProvided(t *testing.T) {
	// Arrange

	var configuration JavaProject
	configuration.DebuggingInformation = []string{"all"}

	// Act
	actualDebuggingInformation := ExtractTestDebuggingInformation(configuration)

	// Assert
	const expectedDebuggingInformation = "-g"

	if actualDebuggingInformation != expectedDebuggingInformation {
		t.Error(
			"For", "ExtractDebuggingInformation",
			"expected", expectedDebuggingInformation,
			"got", actualDebuggingInformation,
		)
	}

}

func TestExtractTestDebuggingInformationNoneProvided(t *testing.T) {
	// Arrange

	var configuration JavaProject
	configuration.DebuggingInformation = []string{"none"}

	// Act
	actualDebuggingInformation := ExtractTestDebuggingInformation(configuration)

	// Assert
	const expectedDebuggingInformation = "-g:none"

	if actualDebuggingInformation != expectedDebuggingInformation {
		t.Error(
			"For", "ExtractDebuggingInformation",
			"expected", expectedDebuggingInformation,
			"got", actualDebuggingInformation,
		)
	}

}

func TestExtractTestDebuggingInformationTwoValuesProvided(t *testing.T) {
	// Arrange

	var configuration JavaProject
	configuration.DebuggingInformation = []string{"source", "lines"}

	// Act
	actualDebuggingInformation := ExtractTestDebuggingInformation(configuration)

	// Assert
	const expectedDebuggingInformation = "-g:source,lines"

	if actualDebuggingInformation != expectedDebuggingInformation {
		t.Error(
			"For", "ExtractDebuggingInformation",
			"expected", expectedDebuggingInformation,
			"got", actualDebuggingInformation,
		)
	}

}

func TestExtractTestDebuggingInformationThreeValuesProvided(t *testing.T) {
	// Arrange

	var configuration JavaProject
	configuration.DebuggingInformation = []string{"source", "lines", "vars"}

	// Act
	actualDebuggingInformation := ExtractTestDebuggingInformation(configuration)

	// Assert
	const expectedDebuggingInformation = "-g:source,lines,vars"

	if actualDebuggingInformation != expectedDebuggingInformation {
		t.Error(
			"For", "ExtractDebuggingInformation",
			"expected", expectedDebuggingInformation,
			"got", actualDebuggingInformation,
		)
	}

}
