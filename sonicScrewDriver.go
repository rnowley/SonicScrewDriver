package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/rnowley/SonicScrewDriver/project"
	"github.com/rnowley/SonicScrewDriver/project/csharp"
	"github.com/rnowley/SonicScrewDriver/project/java"
	"io/ioutil"
	"os"
)

func main() {
	arguments, err := ParseArguments()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	file, err := ioutil.ReadFile("./project.json")

	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	switch arguments.Mode {
	case "build":
		err = BuildProject(file, arguments)
		return
	case "test":
		err = RunUnitTests(file, arguments)
		return
	default:
		fmt.Printf("Invalid mode: %s.", arguments.Mode)
	}

}

func BuildProject(file []byte, arguments project.Arguments) error {
	projectLanguage := GetProjectLanguage(file)

	fmt.Println("Project language: " + projectLanguage)

	var projectBuilder project.ProjectBuilder

	switch projectLanguage {
	case "csharp":
		project := UnmarshalCSharpProject(file)
		command := BuildCsharpCommand(project, arguments)
		projectBuilder = csharp.New(command, project)
	case "java":
		project := UnmarshalJavaProject(file)
		command := BuildJavaCommand(project, arguments)
		projectBuilder = java.NewProjectBuilder(command, project)
	}

	projectBuilder.ExecutePreBuildTasks()
	fmt.Println("Build project")
	projectBuilder.BuildProject()
	fmt.Println("Post build")
	projectBuilder.ExecutePostBuildTasks()
	return nil
}

func RunUnitTests(file []byte, arguments project.Arguments) error {
	project := UnmarshalJavaProject(file)

	// -----------
	// Build the project
	// -----------
	fmt.Println("Building the project")

	BuildProject(file, arguments)

	// -----------
	// Build the unit test project
	// -----------
	fmt.Println("Building unit tests")

	command := BuildTestJavacCommand(project, arguments)
	testProjectBuilder := java.NewProjectBuilder(command, project)

	testProjectBuilder.ExecutePreBuildTasks()
	fmt.Println("Build project")
	testProjectBuilder.BuildProject()
	fmt.Println("Post build")
	testProjectBuilder.ExecutePostBuildTasks()
	fmt.Println("Here")

	// -----------
	// Run the unit test project
	// -----------
	fmt.Println("Running unit tests")
	runCommand := BuildTestRunCommand(project, arguments)
	fmt.Println(runCommand)
	testProjectRunner := java.NewProjectRunner(runCommand, project)
	testProjectRunner.RunProject()
	return nil
}

func ParseArguments() (project.Arguments, error) {
	var arguments project.Arguments

	flag.BoolVar(&arguments.Deprecation, "deprecation", false,
		"Indicates that we wish to show each use or overrides of deprecated class members.")

	flag.Parse()

	nonFlag := flag.Args()
	//noFlagCount := len(nonFlag)

	if len(nonFlag) != 1 {
		return arguments, errors.New("Invalid arguments provided.")
	}
	arguments.Mode = nonFlag[0]

	return arguments, nil
}

// GetProjectLanguage is a function for retrieving the value that
// determines the programming language that the project is written
// in.
func GetProjectLanguage(file []byte) string {
	var projectLanguage project.ProjectLanguage

	if err := json.Unmarshal(file, &projectLanguage); err != nil {
		panic(err)
	}

	return projectLanguage.Language
}

func BuildCsharpCommand(proj csharp.CSharpProject, arguments project.Arguments) csharp.CSharpCommand {
	command := csharp.BuildCommand(proj, arguments)
	return command
}

func BuildJavaCommand(proj java.JavaProject, arguments project.Arguments) java.JavacCommand {
	command := java.BuildCommand(proj, arguments)
	return command
}

func BuildTestJavacCommand(proj java.JavaProject, arguments project.Arguments) java.JavacCommand {

	command := java.BuildTestCommand(proj, arguments)
	return command
}

func BuildTestRunCommand(proj java.JavaProject, arguments project.Arguments) java.JavaCommand {
	command := java.BuildTestRunCommand(proj, arguments)
	return command
}

// UnmarshalCSharpProject is a function that takes in the JSON representation of
// a CSharp project and transforms this into a CSharpProject object.
func UnmarshalCSharpProject(projectFile []byte) csharp.CSharpProject {
	var proj csharp.CSharpProject

	if err := json.Unmarshal(projectFile, &proj); err != nil {
		panic(err)
	}

	return proj
}

// UnmarshalJavaProject is a function that takes in the JSON representation of
// a Java project and transforms this into a JavaProject object.
func UnmarshalJavaProject(projectFile []byte) java.JavaProject {
	var proj java.JavaProject

	if err := json.Unmarshal(projectFile, &proj); err != nil {
		panic(err)
	}
	fmt.Println(proj)
	return proj
}
