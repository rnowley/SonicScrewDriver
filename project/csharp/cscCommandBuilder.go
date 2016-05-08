package csharp

import (
	"fmt"
	"strings"
)

// BuildCommand is a function for building up a csc command that can be used for building
// a .netcore CSharp project. This command is built up using the project configuration.
func GetCscBuildCommand(configuration CSharpProject) CscCommand {
	command := NewDefaultCscCommand()

	if configuration.SourceDirectory != "" {
		command.SourceDirectory = configuration.SourceDirectory
	}

	if configuration.DestinationDirectory != "" {
		command.DestinationDirectory = configuration.DestinationDirectory
	}

	if configuration.OutputFilename != "" {
		command.OutputFilename = fmt.Sprintf("/out:%s%s", command.DestinationDirectory, configuration.OutputFilename)
	}

	command.SourceFiles = ExtractSourceFileList(configuration, command.SourceDirectory)
	command.BuildTarget = ExtractCscBuildTarget(configuration)
	command.References = ExtractCscReferences(configuration)
	command.LibraryPath = ExtractCscLibraryPath(configuration)
	command.ReferencePaths = ExtractCscReferencePaths(configuration)

	return command
}

func ExtractCscBuildTarget(configuration CSharpProject) string {

	switch configuration.BuildTarget {
	case "exe", "library", "module", "winexe", "winmdobj":
		return fmt.Sprintf("/target:%s", configuration.BuildTarget)
	default:
		return "/target:exe"
	}

}

// ExtractLibraryPath extracts all of the library paths provided in the
// configuration file and returns them as an argument item for the compiler.
func ExtractCscLibraryPath(configuration CSharpProject) string {
	fileCount := len(configuration.LibraryPath)

	if fileCount == 0 {
		return ""
	}

	return "/lib:" + strings.Join(configuration.LibraryPath, ",")
}

// ExtractReferences extracts all of the references provided in the
// configuration file and returns them as an argument item for the compiler.
func ExtractCscReferences(configuration CSharpProject) string {
	fileCount := len(configuration.References)

	if fileCount == 0 {
		return ""
	}

	fileList := make([]string, fileCount)

	for i := 0; i < fileCount; i++ {
		fileList[i] = configuration.References[i].Name
	}

	return "/r:" + strings.Join(fileList, ",")
}

func ExtractCscReferencePaths(configuration CSharpProject) []Reference {
	referenceCount := len(configuration.References)
	referenceList := make([]Reference, 0)

	if referenceCount == 0 {
		return referenceList
	}

	for i := 0; i < referenceCount; i++ {
		referenceList = append(referenceList, configuration.References[i])
	}

	return referenceList
}
