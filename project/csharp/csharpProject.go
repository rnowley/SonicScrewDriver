package csharp

type CSharpProject struct {
	Name                 string
	Version              string
	Description          string
	Language             string
	References           []Reference
	SourceFiles          []string
	Resources            []Resource
	BuildTarget          string
	OutputFilename       string
	SourceDirectory      string
	DestinationDirectory string
	LibraryPath          []string
	PackageList          []string
	WarningLevel         string
	WarningsAsErrors     string
}
