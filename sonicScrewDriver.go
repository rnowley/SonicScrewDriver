package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/rnowley/SonicScrewDriver/project"
	"io/ioutil"
	"os"
	"time"
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
	case project.Build:
		err = BuildProject(file, arguments.Mode, arguments)
	case project.BuildTests:
		err = BuildUnitTests(file, arguments.Mode, arguments)
	case project.BuildAll:
		err = BuildAll(file, arguments.Mode, arguments)
	case project.Run:
		err = RunProject(file, arguments.Mode, arguments)
	case project.RunTests:
		err = RunUnitTests(file, arguments.Mode, arguments)
	case project.BuildDocs:
		err = BuildDocumentation(file, arguments.Mode, arguments)
	default:
		fmt.Printf("Invalid mode: %s.", arguments.Mode)
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Exit(0)
}

// BuildProject builds a project of the type determined by the project file, mode and
// the provided arguments.
func BuildProject(file []byte, mode string, arguments project.Arguments) error {
	start := time.Now()
	projectBuilder, _ := project.GetProjectBuilder(file, mode, arguments)

	err := projectBuilder.ExecutePreBuildTasks(arguments.Verbose)

	if err != nil {
		return err
	}

	if arguments.Verbose {
		fmt.Println("Build project")
		fmt.Println("==========")
	}

	err = projectBuilder.BuildProject(arguments.Verbose)

	if err != nil {
		return err
	}

	err = projectBuilder.ExecutePostBuildTasks(arguments.Verbose)
	duration := time.Since(start)

	fmt.Printf("The build took: %02.f:%02.f:%02.f\n", duration.Hours(), duration.Minutes(), duration.Seconds())

	if err != nil {
		return err
	}

	return nil
}

func BuildDocumentation(file []byte, mode string, arguments project.Arguments) error {
	start := time.Now()
	documentationBuilder, _ := project.GetDocumentationBuilder(file, mode, arguments)

	// -----------
	// Generate documentation for the project.
	// -----------
	if arguments.Verbose {
		fmt.Println("Generating documentation")
		fmt.Println("==========")
	}

	err := documentationBuilder.BuildDocumentation(arguments.Verbose)
	duration := time.Since(start)
	fmt.Printf("The documentation took: %02.f:%02.f:%02.f to generate.\n", duration.Hours(), duration.Minutes(), duration.Seconds())

	if err != nil {
		return err
	}

	return nil
}

// BuildUnitTests builds the unit test project. This operation depends on the build project
// having been executed previously.
func BuildUnitTests(file []byte, mode string, arguments project.Arguments) error {
	start := time.Now()

	// -----------
	// Build the unit test project
	// -----------
	if arguments.Verbose {
		fmt.Println("Building unit tests")
		fmt.Println("==========")
	}

	err := BuildProject(file, project.BuildTests, arguments)

	duration := time.Since(start)
	fmt.Printf("The build took: %02.f:%02.f:%02.f\n", duration.Hours(), duration.Minutes(), duration.Seconds())

	if err != nil {
		return err
	}

	return nil
}

// BuildUnitTests builds the unit test project. This operation depends on the build project
// having been executed previously.
func BuildAll(file []byte, mode string, arguments project.Arguments) error {
	start := time.Now()

	// -----------
	// Build the project
	// -----------
	if arguments.Verbose {
		fmt.Println("Building project")
		fmt.Println("==========\n")
	}

	err := BuildProject(file, project.Build, arguments)

	if err != nil {
		return err
	}

	// -----------
	// Build the unit test project
	// -----------
	if arguments.Verbose {
		fmt.Println("==========")
		fmt.Println("Building unit tests")
		fmt.Println("==========")
	}

	err = BuildProject(file, project.BuildTests, arguments)

	duration := time.Since(start)
	fmt.Printf("The build took: %02.f:%02.f:%02.f\n", duration.Hours(), duration.Minutes(), duration.Seconds())

	if err != nil {
		return err
	}

	return nil
}

// BuildUnitTests builds the unit test project. This operation depends on the build project
// having been executed previously.
func RunUnitTests(file []byte, mode string, arguments project.Arguments) error {
	projectRunner, _ := project.GetProjectRunner(file, mode, arguments)

	fmt.Println("Running Unit Tests")
	fmt.Println("------------------")
	err := projectRunner.RunProject()

	if err != nil {
		return err
	}

	return nil
}

// BuildUnitTests builds the unit test project. This operation depends on the build project
// having been executed previously.
func RunProject(file []byte, mode string, arguments project.Arguments) error {
	projectRunner, _ := project.GetProjectRunner(file, mode, arguments)

	fmt.Println("Running project")
	fmt.Println("---------------")
	err := projectRunner.RunProject()

	if err != nil {
		return err
	}

	return nil
}

// ParseArguments parses the arguments passed in through the command line.
func ParseArguments() (project.Arguments, error) {
	var arguments project.Arguments

	flag.BoolVar(&arguments.Deprecation, "deprecation", false,
		"Indicates that we wish to show each use or overrides of deprecated class members.")
	flag.BoolVar(&arguments.Verbose, "verbose", false,
		"Indicates that we wish to show detailed information of the build process.")

	flag.Parse()

	nonFlag := flag.Args() // Holds all arguments that are not used with a flag.

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
