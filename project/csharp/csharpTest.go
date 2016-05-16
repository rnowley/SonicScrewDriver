package csharp

type CSharpTests struct {
	DebugFlag string
	// The source files that belong to the test project.
	SourceFiles []string

	// The name of the test framework to be used.
	TestFrameWork string

	// The destination directory that the build artifacts are to be placed.
	DestinationDirectory string

	// The name of the file for the test project.
	OutputFilename string

	References []Reference

	// The list of resources that are required by the built artifiact.
	Resources []Resource

	LibraryPath  []string
	PackageList  []string
	TestRunner   string
	RunArguments []string
}
