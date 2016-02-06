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

	buildLanguage := GetBuildLanguage(file)

	fmt.Println("Build language: " + buildLanguage)

	var projectBuilder project.ProjectBuilder

	switch buildLanguage {
	case "csharp":
		project := UnmarshalCSharpProject(file)
		command := BuildCsharpCommand(project, arguments)
		projectBuilder = csharp.New(command, project)
	case "java":
		project := UnmarshalJavaProject(file)
		command := BuildJavaCommand(project, arguments)
		projectBuilder = java.New(command, project)
	}

	projectBuilder.ExecutePreBuildTasks()
	fmt.Println("Build project")
	projectBuilder.BuildProject()
	fmt.Println("Post build")
	projectBuilder.ExecutePostBuildTasks()
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

func EnsureDestinationDirectoryExists(destinationDirectory string) {

	_, err := os.Stat(destinationDirectory)

	if err != nil {
		err = os.MkdirAll(destinationDirectory, 0777)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Created directory.")
		return
	}

	fmt.Println("File already exists, nothing to do.")
}

func GetBuildLanguage(file []byte) string {
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

func BuildJavaCommand(proj java.JavaProject, arguments project.Arguments) java.JavaCommand {
	command := java.BuildCommand(proj, arguments)
	return command
}

func UnmarshalCSharpProject(projectFile []byte) csharp.CSharpProject {
	var proj csharp.CSharpProject

	if err := json.Unmarshal(projectFile, &proj); err != nil {
		panic(err)
	}

	return proj
}

func UnmarshalJavaProject(projectFile []byte) java.JavaProject {
	var proj java.JavaProject

	if err := json.Unmarshal(projectFile, &proj); err != nil {
		panic(err)
	}

	return proj
}
