package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/rnowley/SonicScrewDriver/project"
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
		err = BuildProject(file, arguments.Mode, arguments)
		return
	case "build-test":
		err = BuildUnitTests(file, arguments.Mode, arguments)
		return
		case "build-all":
		err = BuildAll(file, arguments.Mode, arguments)
	default:
		fmt.Printf("Invalid mode: %s.", arguments.Mode)
	}

}

// BuildProject builds a project of the type determined by the project file, mode and
// the provided arguments.
func BuildProject(file []byte, mode string, arguments project.Arguments) error {
	projectBuilder, _ := project.GetProjectBuilder(file, mode, arguments)

	projectBuilder.ExecutePreBuildTasks()
	fmt.Println("Build project")
	projectBuilder.BuildProject()
	fmt.Println("Post build")
	projectBuilder.ExecutePostBuildTasks()
	return nil
}

// BuildUnitTests builds the unit test project. This operation depends on the build project
// having been executed previously.
func BuildUnitTests(file []byte, mode string, arguments project.Arguments) error {

	// -----------
	// Build the unit test project
	// -----------
	fmt.Println("Building unit tests")

	BuildProject(file, "build-test", arguments)

	return nil
}

// BuildUnitTests builds the unit test project. This operation depends on the build project
// having been executed previously.
func BuildAll(file []byte, mode string, arguments project.Arguments) error {
	// -----------
	// Build the project
	// -----------
	fmt.Println("Building project")

	BuildProject(file, "build", arguments)

	// -----------
	// Build the unit test project
	// -----------
	fmt.Println("Building unit tests")

	BuildProject(file, "build-test", arguments)

	return nil
}

// ParseArguments parses the arguments passed in through the command line.
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
