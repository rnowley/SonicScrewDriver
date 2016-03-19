package java

import (
	"strings"
	"testing"
)

func TestGetJavadocBuildCommandVerboseTrue(t *testing.T) {

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
	var commandToTest = GetJavadocBuildCommand(configuration, true)

	// Assert

	if commandToTest.Verbose != true {
		t.Error(
			"For", "TestGetJavadocBuildCommandVerboseTrue",
			"expected", true, "got",
			commandToTest.Verbose,
		)
	}
}

func TestGetJavadocBuildCommandVerboseFalse(t *testing.T) {

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
	var commandToTest = GetJavadocBuildCommand(configuration, false)

	// Assert

	if commandToTest.Verbose != false {
		t.Error(
			"For", "TestGetJavadocBuildCommandVerboseFalse",
			"expected", false, "got",
			commandToTest.Verbose,
		)
	}
}

func TestGetJavadocBuildCommand(t *testing.T) {
	// Arrange

	var configuration JavaProject
	configuration.SourceFiles = []string{"./src/a", "./src/b"}

	var docConfiguration JavaDocumentation
	docConfiguration.DestinationDirectory = "./doctest/"
	docConfiguration.SourcePath = []string{"./src/a", "./src/b"}
	docConfiguration.ClassPath = []string{"./lib/y", "./lib/x"}
	docConfiguration.LinkSource = true
	docConfiguration.AccessLevel = "private"
	docConfiguration.WindowTitle = "Test Title"
	docConfiguration.LintWarnings = []string{"syntax", "accessibility"}
	docConfiguration.Verbose = true
	docConfiguration.DocTitle = "Test Title"
	docConfiguration.Header = "Test header"
	docConfiguration.Bottom = "Test bottom text"

	configuration.DocumentationProject = docConfiguration

	// Act

	commandToTest := GetJavadocBuildCommand(configuration, false)

	// Assert
	const expectedDestinationDirectory = "./doctest/"

	if commandToTest.DestinationDirectory != expectedDestinationDirectory {
		t.Error(
			"For", "TestGetJavadocBuildCommand",
			"expected", expectedDestinationDirectory, "got",
			commandToTest.DestinationDirectory,
		)
	}

	expectedSourcePath := "./src/a ./src/b"
	actualSourcePath := strings.Join(commandToTest.SourcePath, " ")

	if actualSourcePath != expectedSourcePath {
		t.Error(
			"For", "TestGetJavadocBuildCommand",
			"expected", expectedSourcePath, "got",
			actualSourcePath,
		)
	}

	const expectedClassPath = "./lib/y;./lib/x"

	if commandToTest.ClassPath != expectedClassPath {
		t.Error(
			"For", "TestGetJavadocBuildCommand",
			"expected", expectedClassPath, "got",
			commandToTest.ClassPath,
		)
	}

	const expectedLinkSource = true

	if commandToTest.LinkSource != expectedLinkSource {
		t.Error(
			"For", "TestGetJavadocBuildCommand",
			"expected", expectedLinkSource, "got",
			commandToTest.LinkSource,
		)
	}

	const expectedWindowTitle = "Test Title"

	if commandToTest.WindowTitle != expectedWindowTitle {
		t.Error(
			"For", "TestGetJavadocBuildCommand",
			"expected", expectedWindowTitle, "got",
			commandToTest.WindowTitle,
		)
	}

	const expectedDocTitle = "Test Title"

	if commandToTest.DocTitle != expectedDocTitle {
		t.Error(
			"For", "TestGetJavadocBuildCommand",
			"expected", expectedDocTitle, "got",
			commandToTest.DocTitle,
		)
	}

	const expectedHeader = "Test header"

	if commandToTest.Header != expectedHeader {
		t.Error(
			"For", "TestGetJavadocBuildCommand",
			"expected", expectedHeader, "got",
			commandToTest.Header,
		)
	}

	const expectedBottom = "Test bottom text"

	if commandToTest.Bottom != expectedBottom {
		t.Error(
			"For", "TestGetJavadocBuildCommand",
			"expected", expectedBottom, "got",
			commandToTest.Bottom,
		)
	}
}
