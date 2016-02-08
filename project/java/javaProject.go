package java

type JavaProject struct {
	Name                 string
	Version              string
	Description          string
	Language             string
	DestinationDirectory string
	ClassPath            []string
	SourceFiles          []string
	SourceVersion        string
	DebuggingInformation []string
	TestProject          JavaTests
}
