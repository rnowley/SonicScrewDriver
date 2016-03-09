package java

import (
	"strings"
	"testing"
)

func TestGetJavaBuildCommandDeprecatedTrue(t *testing.T) {
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

	// Act
	var commandToTest = GetJavaBuildCommand(configuration, true, false)

	// Assert
	const expectedDestinationDirectory = "./build/"

	if commandToTest.DestinationDirectory != expectedDestinationDirectory {
		t.Error(
			"For", "TestGetJavaBuildCommandDeprecatedTrue",
			"expected", expectedDestinationDirectory, "got",
			commandToTest.DestinationDirectory,
		)
	}

	const expectedSourceDirectory = "./src/"

	if commandToTest.SourceDirectory != expectedSourceDirectory {
		t.Error(
			"For", "TestGetJavaBuildCommandDeprecatedTrue",
			"expected", expectedSourceDirectory, "got",
			commandToTest.SourceDirectory,
		)
	}

	const expectedClassPath = "./lib/a.jar ./lib/b.jar"
	actualClassPath := strings.Join(commandToTest.ClassPath, " ")

	if actualClassPath != expectedClassPath {
		t.Error(
			"For", "TestGetJavaBuildCommandDeprecatedTrue",
			"expected", expectedClassPath, "got",
			actualClassPath,
		)
	}

	const expectedDeprecation = true

	if commandToTest.Deprecation != expectedDeprecation {
		t.Error(
			"For", "TestGetJavaBuildCommandDeprecatedTrue",
			"expected", expectedDeprecation, "got",
			commandToTest.Deprecation,
		)
	}

	const expectedSourceFileList = "./src/a.java ./src/b.java ./src/c.java"
	actualSourceFileList := strings.Join(commandToTest.SourceFiles, " ")

	if actualSourceFileList != expectedSourceFileList {
		t.Error(
			"For", "TestGetJavaBuildCommandDeprecatedTrue",
			"expected", expectedSourceFileList, "got",
			actualSourceFileList,
		)
	}

	const expectedSourceVersion = "1.7"

	if commandToTest.SourceVersion != expectedSourceVersion {
		t.Error(
			"For", "TestGetJavaBuildCommandDeprecatedTrue",
			"expected", expectedSourceVersion, "got",
			commandToTest.SourceVersion,
		)
	}

	const expectedDebuggingInformation = "-g"

	if commandToTest.DebuggingInformation != expectedDebuggingInformation {
		t.Error(
			"For", "TestGetJavaBuildCommandDeprecatedTrue",
			"expected", expectedDebuggingInformation, "got",
			commandToTest.DebuggingInformation,
		)
	}

	const expectedTarget = "1.8"

	if commandToTest.Target != expectedTarget {
		t.Error(
			"For", "TestGetJavaBuildCommandDeprecatedTrue",
			"expected", expectedTarget, "got",
			commandToTest.Target,
		)
	}

}

func TestGetJavaBuildCommandVerboseTrue(t *testing.T) {
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

	// Act
	var commandToTest = GetJavaBuildCommand(configuration, true, true)

	// Assert
	const expectedDestinationDirectory = "./build/"

	if commandToTest.DestinationDirectory != expectedDestinationDirectory {
		t.Error(
			"For", "TestGetJavaBuildCommandVerboseTrue",
			"expected", expectedDestinationDirectory, "got",
			commandToTest.DestinationDirectory,
		)
	}

	const expectedSourceDirectory = "./src/"

	if commandToTest.SourceDirectory != expectedSourceDirectory {
		t.Error(
			"For", "TestGetJavaBuildCommandVerboseTrue",
			"expected", expectedSourceDirectory, "got",
			commandToTest.SourceDirectory,
		)
	}

	const expectedClassPath = "./lib/a.jar ./lib/b.jar"
	actualClassPath := strings.Join(commandToTest.ClassPath, " ")

	if actualClassPath != expectedClassPath {
		t.Error(
			"For", "TestGetJavaBuildCommandVerboseTrue",
			"expected", expectedClassPath, "got",
			actualClassPath,
		)
	}

	const expectedDeprecation = true

	if commandToTest.Deprecation != expectedDeprecation {
		t.Error(
			"For", "TestGetJavaBuildCommandVerboseTrue",
			"expected", expectedDeprecation, "got",
			commandToTest.Deprecation,
		)
	}

	const expectedSourceFileList = "./src/a.java ./src/b.java ./src/c.java"
	actualSourceFileList := strings.Join(commandToTest.SourceFiles, " ")

	if actualSourceFileList != expectedSourceFileList {
		t.Error(
			"For", "TestGetJavaBuildCommandVerboseTrue",
			"expected", expectedSourceFileList, "got",
			actualSourceFileList,
		)
	}

	const expectedSourceVersion = "1.7"

	if commandToTest.SourceVersion != expectedSourceVersion {
		t.Error(
			"For", "TestGetJavaBuildCommandVerboseTrue",
			"expected", expectedSourceVersion, "got",
			commandToTest.SourceVersion,
		)
	}

	const expectedDebuggingInformation = "-g"

	if commandToTest.DebuggingInformation != expectedDebuggingInformation {
		t.Error(
			"For", "TestGetJavaBuildCommandVerboseTrue",
			"expected", expectedDebuggingInformation, "got",
			commandToTest.DebuggingInformation,
		)
	}

	const expectedVerbose = true

	if commandToTest.Verbose != expectedVerbose {
		t.Error(
			"For", "TestGetJavaBuildCommandVerboseTrue",
			"expected", expectedVerbose, "got",
			commandToTest.Verbose,
		)
	}

}

func TestExtractSourceFileListNoSourceFiles(t *testing.T) {
	// Arrange

	var configuration JavaProject
	configuration.SourceFiles = []string{}

	// Act
	fileList := ExtractSourceFileList(configuration, "./src/")

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

func TestExtractSourceFileListOneSourceFile(t *testing.T) {
	// Arrange

	var configuration JavaProject
	configuration.SourceFiles = []string{"a.java"}

	// Act
	fileList := ExtractSourceFileList(configuration, "./src/")

	// Assert
	const expectedSourceFileList = "./src/a.java"
	actualSourceFileList := strings.Join(fileList, " ")

	if actualSourceFileList != expectedSourceFileList {
		t.Error(
			"For", "ExtractSourceFileList",
			"expected", expectedSourceFileList, "got",
			actualSourceFileList,
		)
	}

}

func TestExtractSourceFileListThreeSourceFiles(t *testing.T) {
	// Arrange

	var configuration JavaProject
	configuration.SourceFiles = []string{"a.java", "b.java", "c.java"}

	// Act
	fileList := ExtractSourceFileList(configuration, "./src/")

	// Assert
	const expectedSourceFileList = "./src/a.java ./src/b.java ./src/c.java"
	actualSourceFileList := strings.Join(fileList, " ")

	if actualSourceFileList != expectedSourceFileList {
		t.Error(
			"For", "ExtractSourceFileList",
			"expected", expectedSourceFileList, "got",
			actualSourceFileList,
		)
	}

}

func TestExtractDebuggingInformationNoValueProvided(t *testing.T) {
	// Arrange

	var configuration JavaProject
	configuration.DebuggingInformation = []string{}

	// Act
	actualDebuggingInformation := ExtractDebuggingInformation(configuration)

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

func TestExtractDebuggingInformationAllProvided(t *testing.T) {
	// Arrange

	var configuration JavaProject
	configuration.DebuggingInformation = []string{"all"}

	// Act
	actualDebuggingInformation := ExtractDebuggingInformation(configuration)

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

func TestExtractDebuggingInformationNoneProvided(t *testing.T) {
	// Arrange

	var configuration JavaProject
	configuration.DebuggingInformation = []string{"none"}

	// Act
	actualDebuggingInformation := ExtractDebuggingInformation(configuration)

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

func TestExtractDebuggingInformationTwoValuesProvided(t *testing.T) {
	// Arrange

	var configuration JavaProject
	configuration.DebuggingInformation = []string{"source", "lines"}

	// Act
	actualDebuggingInformation := ExtractDebuggingInformation(configuration)

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

func TestExtractDebuggingInformationThreeValuesProvided(t *testing.T) {
	// Arrange

	var configuration JavaProject
	configuration.DebuggingInformation = []string{"source", "lines", "vars"}

	// Act
	actualDebuggingInformation := ExtractDebuggingInformation(configuration)

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

func TestLintWarningsNoneProvided(t *testing.T) {
	// Arrange

	var configuration JavaProject
	configuration.LintWarnings = []string{}

	// Act
	actualLintWarnings := ExtractLintWarnings(configuration)

	// Assert
	const expectedLintWarnings = ""

	if actualLintWarnings != expectedLintWarnings {
		t.Error(
			"For", "TestLintWarningsNoneProvided",
			"expected", expectedLintWarnings,
			"got", actualLintWarnings,
		)
	}

}

func TestLintWarningsOneValueProvided(t *testing.T) {
	// Arrange

	var configuration JavaProject
	configuration.LintWarnings = []string{"all"}

	// Act
	actualLintWarnings := ExtractLintWarnings(configuration)

	// Assert
	const expectedLintWarnings = "-Xlint"

	if actualLintWarnings != expectedLintWarnings {
		t.Error(
			"For", "TestLintWarningsOneValueProvided",
			"expected", expectedLintWarnings,
			"got", actualLintWarnings,
		)
	}

}

func TestTestLintWarningsThreeValuesProvided(t *testing.T) {
	// Arrange

	var configuration JavaProject
	configuration.LintWarnings = []string{"all", "-cast", "-fallthrough"}

	// Act
	actualLintWarnings := ExtractLintWarnings(configuration)

	// Assert
	const expectedLintWarnings = "-Xlint:all,-cast,-fallthrough"

	if actualLintWarnings != expectedLintWarnings {
		t.Error(
			"For", "TestTestLintWarningsThreeValuesProvided",
			"expected", expectedLintWarnings,
			"got", actualLintWarnings,
		)
	}

}
