package csharp

import (
	"fmt"
	"github.com/rnowley/SonicScrewDriver/project"
	"strconv"
	"strings"
)

func BuildCommand(configuration CSharpProject, arguments project.Arguments) CSharpCommand {
	command := NewDefaultCommand()

	if configuration.SourceDirectory != "" {
		command.SourceDirectory = configuration.SourceDirectory
	}

	if configuration.DestinationDirectory != "" {
		command.DestinationDirectory = configuration.DestinationDirectory
	}

	if configuration.OutputFilename != "" {
		command.OutputFilename = fmt.Sprintf("-out:%s%s", command.DestinationDirectory, configuration.OutputFilename)
	}

	command.SourceFiles = ExtractSourceFileList(configuration, command.SourceDirectory)
	command.BuildTarget = ExtractBuildTarget(configuration)
	command.SourceFiles = ExtractSourceFileList(configuration, command.SourceDirectory)
	command.References = ExtractReferences(configuration)
	command.LibraryPath = ExtractLibraryPath(configuration)
	command.PackageList = ExtractPackageList(configuration)
	command.WarningLevel = SetWarningLevel(configuration)
	command.WarningsAsErrors = TreatWarningsAsErrors(configuration)

	return command
}

func ExtractSourceFileList(configuration CSharpProject, sourceDirectory string) []string {
	fileCount := len(configuration.SourceFiles)
	fileList := make([]string, fileCount)

	for i := 0; i < fileCount; i++ {
		fileList[i] = sourceDirectory + configuration.SourceFiles[i]
	}

	return fileList
}

func ExtractBuildTarget(configuration CSharpProject) string {

	switch configuration.BuildTarget {
	case "exe", "library", "module", "winexe":
		return fmt.Sprintf("-target:%s", configuration.BuildTarget)
	}

	return "-target:exe"
}

func ExtractLibraryPath(configuration CSharpProject) string {
	fileCount := len(configuration.LibraryPath)

	if fileCount == 0 {
		return ""
	}

	return "-lib:" + strings.Join(configuration.LibraryPath, ",")
}

func ExtractPackageList(configuration CSharpProject) string {
	fileCount := len(configuration.PackageList)

	if fileCount == 0 {
		return ""
	}

	return "-pkg:" + strings.Join(configuration.PackageList, ",")
}

func ExtractReferences(configuration CSharpProject) string {
	fileCount := len(configuration.References)

	if fileCount == 0 {
		return ""
	}

	fileList := make([]string, fileCount)

	for i := 0; i < fileCount; i++ {
		fileList[i] = configuration.References[i].Name
	}

	return "-r:" + strings.Join(fileList, ",")
}

func SetWarningLevel(configuration CSharpProject) string {

	if configuration.WarningLevel == "" {
		return ""
	}

	warningLevel, err := strconv.ParseInt(configuration.WarningLevel, 10, 32)

	if err != nil {
		fmt.Printf("Warning: Invalid value for warning level (%s), using the default value for the compiler.", configuration.WarningLevel)
		return ""
	}

	if warningLevel >= 0 && warningLevel <= 4 {
		return fmt.Sprintf("-warn:%d", warningLevel)
	}

	fmt.Printf("Warning: Invalid value for warning level (%s), using the default value for the compiler.", configuration.WarningLevel)

	return ""
}

func TreatWarningsAsErrors(configuration CSharpProject) string {

	if configuration.WarningsAsErrors == "" {
		return ""
	}

	warningsAsErrors, err := strconv.ParseBool(configuration.WarningsAsErrors)

	if err != nil {
		fmt.Printf("Warning: Invalid value for warning as errors (%s), using the default value for the compiler.", configuration.WarningsAsErrors)
		return ""
	}

	if warningsAsErrors {
		return "-warnaserror+"
	}

	return "-warnaserror-"
}