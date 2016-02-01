package main

import (
	"./project"
	"encoding/json"
	"flag"
	"fmt"
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

	var proj project.ProjectConfiguration

	if err := json.Unmarshal(file, &proj); err != nil {
		panic(err)
	}

	command := project.BuildCommand(proj, arguments)
	EnsureDestinationDirectoryExists(command)
	BuildProject(command)
}

func BuildProject(command *project.Command) {
	binary, lookErr := exec.LookPath(command.CommandName)

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

func EnsureDestinationDirectoryExists(command *project.Command) {

	if command.DestinationDirectory == "" {
		command.DestinationDirectory = "./build"
	}

	_, err := os.Stat(command.DestinationDirectory)

	if err != nil {
		err = os.MkdirAll(command.DestinationDirectory, 0777)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Created directory.")
		return
	}



	fmt.Println("File already exists, nothing to do.")
}
