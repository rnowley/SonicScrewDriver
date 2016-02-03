package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/rnowley/SonicScrewDriver/project"
	"github.com/rnowley/SonicScrewDriver/project/csharp"
	"github.com/rnowley/SonicScrewDriver/project/java"
	"io/ioutil"
	"os"
	"os/exec"
	"syscall"
)

func main() {

	var arguments = ParseArguments()

	file, err := ioutil.ReadFile("./project.json")

	if err != nil {
		fmt.Printf("File error: %v\n", err)
		os.Exit(1)
	}

	buildLanguage := GetBuildLanguage(file)

	fmt.Println("Build language: " + buildLanguage)

	var command project.Command

	switch buildLanguage {
	case "csharp":
		command = BuildCsharpCommand(file, arguments)
	case "java":
		command = BuildJavaCommand(file, arguments)
	}

	EnsureDestinationDirectoryExists(command.GetDestinationDirectory())
	BuildProject(command)
}

func BuildProject(command project.Command) {
	binary, lookErr := exec.LookPath(command.GetCommandName())

	if lookErr != nil {
		panic(lookErr)
	}

	args := command.GenerateArgumentList()
	fmt.Println(args)
	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)

	if execErr != nil {
		panic(execErr)
	}

}

func ParseArguments() project.Arguments {
	var arguments project.Arguments

	flag.BoolVar(&arguments.Deprecation, "deprecation", false,
		"Indicates that we wish to show each use or overrides of deprecated class members.")

	flag.Parse()

	return arguments
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

func BuildCsharpCommand(projectFile []byte, arguments project.Arguments) project.Command {
	var command csharp.CSharpCommand

	var proj csharp.CSharpProject

	if err := json.Unmarshal(projectFile, &proj); err != nil {
		panic(err)
	}

	command = csharp.BuildCommand(proj, arguments)
	return command
}

func BuildJavaCommand(projectFile []byte, arguments project.Arguments) project.Command {
	var proj java.JavaProject

	if err := json.Unmarshal(projectFile, &proj); err != nil {
		panic(err)
	}

	command := java.BuildCommand(proj, arguments)
	return command
}
