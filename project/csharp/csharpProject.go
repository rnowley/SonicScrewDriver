package csharp

// CSharpProject is a struct for holding the information required for
// building or running a CSharp project.
type CSharpProject struct {
	// The name of the project.
	Name string

	// Holds the project version number.
	Version string

	// Holds a description of the project.
	Description string

	// Holds the name of the programming language that the project is written in
	Language string

	// Holds the list of references required to build the project.
	References []Reference

	// Holds a list of all of the files to be compiled into the project.
	SourceFiles []string

	// The list of resources that are required by the built artifiact.
	Resources []Resource

	// The desired target type of the build.
	BuildTarget string

	// The name of the file to be output by the compiler.
	OutputFilename string

	// The base directory for source files.
	SourceDirectory string

	// The directory that build artifacts are going to be placed in.
	DestinationDirectory string

	LibraryPath []string

	// The list of packages required to build the project.
	PackageList []string

	// Used to control the level of warnings provided by the compiler.
	WarningLevel string

	// Used to indicate whether or not to treat warnings as errors.
	WarningsAsErrors string

	TestProject CSharpTests
}
