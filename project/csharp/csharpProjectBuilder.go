package csharp

import (
	"bytes"
	"fmt"
	"github.com/rnowley/SonicScrewDriver/utilities"
	"os"
	"os/exec"
	"path/filepath"
)

type CSharpProjectBuilder struct {
	command CSharpCommand
	project CSharpProject
}

func New(command CSharpCommand, project CSharpProject) CSharpProjectBuilder {
	return CSharpProjectBuilder{command, project}
}

func (builder CSharpProjectBuilder) ExecutePreBuildTasks() error {
	err := builder.ensureDestinationDirectoryExists()
	return err
}

func (builder CSharpProjectBuilder) BuildProject() error {
	binary, lookErr := exec.LookPath(builder.command.GetCommandName())
	fmt.Println(binary)

	if lookErr != nil {
		return lookErr
	}
	args := []string{binary, "-out:./build/nancySelfHost.exe", "", "-target:exe", "-lib:./lib/Nancy.1.4.1/lib/net40/,./lib/Nancy.Hosting.Self.1.4.1/lib/net40/",
		"-r:Nancy,Nancy.Hosting.Self", "./src/Program.cs", "", "", ""}
	//args := builder.command.GenerateArgumentList()
	//env := os.Environ()
	fmt.Println(args)
	cmd := exec.Command(binary, "")
	//cmd := exec.Command("mcs", "-out:./build/nancySelfHost.exe", "-target:exe", "-lib:./lib/Nancy.1.4.1/lib/net40/,./lib/Nancy.Hosting.Self.1.4.1/lib/net40/", "-r:Nancy,Nancy.Hosting.Self", "./src/Program.cs")
	cmd.Args = args
	var out bytes.Buffer
	cmd.Stdout = &out
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	fmt.Printf("**%s\n", out.String())
	if err != nil {
		return err
	}

	return nil
}

func (builder CSharpProjectBuilder) ExecutePostBuildTasks() error {
	fmt.Println("Post build tasks")
	builder.copyReferences()
	return nil
}

func (builder CSharpProjectBuilder) copyReferences() error {
	referenceCount := len(builder.project.References)
	fmt.Println(referenceCount)

	if referenceCount == 0 {
		return nil
	}

	destinationDirectory := builder.command.DestinationDirectory

	for i := 0; i < referenceCount; i++ {

		if builder.project.References[i].Path == "" {
			continue
		}

		path := builder.project.References[i].Path
		referenceName := builder.project.References[i].Name
		fileExtension := ".dll"

		utilities.CopyFile(fmt.Sprintf("%s%s%s", destinationDirectory, referenceName, fileExtension),
			fmt.Sprintf("%s%s%s", path, referenceName, fileExtension),
		)
	}
	return nil
}

func (builder CSharpProjectBuilder) copyResources() error {
	resourceCount := len(builder.project.Resources)

	if resourceCount == 0 {
		return nil
	}

	for i := 0; i < resourceCount; i++ {
		destinationDirectory := filepath.Dir(builder.project.Resources[i].Destination)

		utilities.EnsurePathExists(fmt.Sprintf("%s%s", builder.command.GetDestinationDirectory(), destinationDirectory))
		utilities.CopyFile(fmt.Sprintf("%s%s", builder.command.SourceDirectory, builder.project.Resources[i].Source),
			fmt.Sprintf("%s%s", builder.command.DestinationDirectory, builder.project.Resources[i].Destination))
	}

	return nil
}

func (builder CSharpProjectBuilder) ensureDestinationDirectoryExists() error {
	destinationDirectory := builder.command.GetDestinationDirectory()
	_, err := os.Stat(destinationDirectory)

	if err != nil {
		err = os.MkdirAll(destinationDirectory, 0777)

		if err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println("Created directory.")
		return nil
	}

	fmt.Println("File already exists, nothing to do.")
	return err
}
