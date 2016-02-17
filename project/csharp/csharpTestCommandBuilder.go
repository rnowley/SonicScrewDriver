package csharp

import (
	"fmt"
	"strconv"
	"strings"
)

// BuildCommand is a function for building up a mcs command that can be used for building
// a CSharp project. This command is built up using the project configuration.
func GetCSharpTestBuildCommand(configuration CSharpProject) CSharpCommand {
	command := NewDefaultCommand()

	if configuration.SourceDirectory != "" {
		command.SourceDirectory = configuration.SourceDirectory
	}

	if configuration.DestinationDirectory != "" {
		command.DestinationDirectory = configuration.TestProject.DestinationDirectory
	}

	if configuration.TestProject.OutputFilename != "" {
		command.OutputFilename = fmt.Sprintf("-out:%s%s", command.DestinationDirectory, configuration.TestProject.OutputFilename)
	}

	command.BuildTarget = ExtractTestBuildTarget(configuration)
	command.SourceFiles = ExtractTestSourceFileList(configuration, command.SourceDirectory)
	command.References = ExtractTestReferences(configuration)
	command.LibraryPath = ExtractTestLibraryPath(configuration)
	command.PackageList = ExtractTestPackageList(configuration)
	//command.WarningLevel = SetWarningLevel(configuration)
	//command.WarningsAsErrors = TreatWarningsAsErrors(configuration)
	command.ReferencesPaths = ExtractTestReferencePaths(configuration)

	return command
}

// ExtractSourceFileList is a function that reads all of the source files to be
// compiled from the configuration file and returns a slice of source files to be
// compiled using the mcs command. Each source file has had the base path appended
// to it when returned from the function.
func ExtractTestSourceFileList(configuration CSharpProject, sourceDirectory string) []string {
	fileCount := len(configuration.TestProject.SourceFiles)
	fileList := make([]string, fileCount)

	for i := 0; i < fileCount; i++ {
		fileList[i] = sourceDirectory + configuration.TestProject.SourceFiles[i]
	}

	return fileList
}

func ExtractTestBuildTarget(configuration CSharpProject) string {
	return "-target:library"
}

// ExtractLibraryPath extracts all of the library paths provided in the
// configuration file and returns them as an argument item for the compiler.
func ExtractTestLibraryPath(configuration CSharpProject) string {
	fileCount := len(configuration.TestProject.LibraryPath)

	if fileCount == 0 {
		return ""
	}

	return "-lib:" + strings.Join(configuration.TestProject.LibraryPath, ",")
}

// ExtractPackageList extracts all of the packages provided in the
// configuration file and returns them as an argument item for the compiler.
func ExtractTestPackageList(configuration CSharpProject) string {
	fileCount := len(configuration.PackageList)

	if fileCount == 0 {
		return ""
	}

	return "-pkg:" + strings.Join(configuration.PackageList, ",")
}

// ExtractReferences extracts all of the references provided in the
// configuration file and returns them as an argument item for the compiler.
func ExtractTestReferences(configuration CSharpProject) string {
	fileCount := len(configuration.References)

	/*if fileCount == 0 {
		return ""
	}*/

	fileList := make([]string, fileCount)

	for i := 0; i < fileCount; i++ {
		fileList = append(fileList, configuration.References[i].Name)
		//fileList[i] = configuration.References[i].Name
	}

	fileCount = len(configuration.TestProject.References)

	for i := 0; i < fileCount; i++ {
		fileList = append(fileList, configuration.TestProject.References[i].Name)
		//fileList[i] = configuration.References[i].Name
	}

	return "-r:" + strings.Join(fileList, ",")
}

func ExtractTestReferencePaths(configuration CSharpProject) []Reference {
	referenceCount := len(configuration.References)
	referenceList := make([]Reference, referenceCount)

	for i := 0; i < referenceCount; i++ {
		referenceList = append(referenceList, configuration.References[i])
	}

	referenceCount = len(configuration.TestProject.References)

	if referenceCount == 0 {
		return referenceList
	}

	for i := 0; i < referenceCount; i++ {
		referenceList = append(referenceList, configuration.TestProject.References[i])
	}

	fmt.Println(referenceList)

	return referenceList
}

// SetWarningLevel extracts the provided warning level to be used from the
// configuration file and returned as an argument for the compiler if one is provided.
func SetTestWarningLevel(configuration CSharpProject) string {

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

// TreatWarningsAsErrors determines if the treat warnings as errors option is enabled
// in the configuration file. If no value is provided then the default for the compiler is used.
func TestTreatWarningsAsErrors(configuration CSharpProject) string {

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
