package project

import (
	"fmt"

	"github.com/rnowley/SonicScrewDriver/project/csharp"
	"github.com/rnowley/SonicScrewDriver/project/java"
	"github.com/rnowley/SonicScrewDriver/project/kotlin"
	"github.com/rnowley/SonicScrewDriver/project/scala"
)

// GetProjectBuilder is a factory function that returns an object that implements the
// ProjectBuilder interface. The type of this ProjectBuilder is determined by the
// configuration file that is passed in.
func GetProjectBuilder(configurationFile []byte, mode string, arguments Arguments) (ProjectBuilder, error) {
	projectLanguage := GetProjectLanguage(configurationFile)
	var projectBuilder ProjectBuilder

	switch projectLanguage {
	case "csharp":
		projectBuilder, _ = getCSharpProjectBuilder(configurationFile, mode)
		return projectBuilder, nil
	case "java":
		projectBuilder, _ = getJavaProjectBuilder(configurationFile, mode, arguments)
		return projectBuilder, nil
	case "kotlin":
		projectBuilder, _ = getKotlinProjectBuilder(configurationFile, mode, arguments)
		return projectBuilder, nil
	case "scala":
		projectBuilder, _ = getScalaProjectBuilder(configurationFile, mode, arguments)
		return projectBuilder, nil
	default:
		return projectBuilder, fmt.Errorf("GetProjectBuilder: the %s language is not supported", projectLanguage)
	}

}

// getCSharpProjectBuilder is called by the factory function to return a
// project builder for creating building a CSharp project.
func getCSharpProjectBuilder(configurationFile []byte, mode string) (ProjectBuilder, error) {
	var proj csharp.CSharpProject
	var projectBuilder csharp.CSharpProjectBuilder

	proj, _ = UnmarshalCSharpProject(configurationFile)

	var command csharp.CSharpCommand

	switch mode {
	case "build":
		command = csharp.GetCSharpBuildCommand(proj)
	case "build-test":
		command = csharp.GetCSharpTestBuildCommand(proj)
	default:
		return projectBuilder, fmt.Errorf("getCSharpProjectBuilder: the %s 'mode' is not supported", mode)
	}

	projectBuilder = csharp.New(command, proj)

	return projectBuilder, nil
}

// getJavaProjectBuilder is called by the factory function to return a
// project builder for creating building a Java project.
func getJavaProjectBuilder(configurationFile []byte, mode string, arguments Arguments) (ProjectBuilder, error) {
	var proj java.JavaProject
	var projectBuilder java.JavaProjectBuilder

	proj = UnmarshalJavaProject(configurationFile)

	var command java.JavacCommand

	fmt.Println("Verbose: %s", arguments.Verbose)

	switch mode {
	case "build":
		command = java.GetJavaBuildCommand(proj, arguments.Deprecation, arguments.Verbose)
	case "build-test":
		command = java.GetJavaTestBuildCommand(proj, arguments.Deprecation)
	default:
		return projectBuilder, fmt.Errorf("getJavaProjectBuilder: the %s 'mode' is not supported", mode)
	}

	projectBuilder = java.NewProjectBuilder(command, proj)

	return projectBuilder, nil
}

// getKotlinProjectBuilder is called by the factory function to return a
// project builder for creating building a Kotlin project.
func getKotlinProjectBuilder(configurationFile []byte, mode string, arguments Arguments) (ProjectBuilder, error) {
	var proj kotlin.KotlinProject
	var projectBuilder kotlin.KotlinProjectBuilder

	proj = UnmarshalKotlinProject(configurationFile)

	var command kotlin.KotlincCommand

	switch mode {
	case "build":
		command = kotlin.GetKotlincBuildCommand(proj, arguments.Verbose)
	case "build-test":
		command = kotlin.GetKotlincTestBuildCommand(proj, arguments.Verbose)
	default:
		return projectBuilder, fmt.Errorf("getKotlinProjectBuilder: the %s 'mode' is not supported", mode)
	}

	projectBuilder = kotlin.NewProjectBuilder(command, proj)

	return projectBuilder, nil
}

// getScalaProjectBuilder is called by the factory function to return a
// project builder for creating building a Scala project.
func getScalaProjectBuilder(configurationFile []byte, mode string, arguments Arguments) (ProjectBuilder, error) {
	var proj scala.ScalaProject
	var projectBuilder scala.ScalaProjectBuilder

	proj = UnmarshalScalaProject(configurationFile)

	var command scala.ScalacCommand

	switch mode {
	case "build":
		command = scala.GetScalacBuildCommand(proj, arguments.Verbose, arguments.Deprecation)
	default:
		return projectBuilder, fmt.Errorf("getScalaProjectBuilder: the %s 'mode' is not supported", mode)
	}

	projectBuilder = scala.NewScalaProjectBuilder(command, proj)

	return projectBuilder, nil
}
