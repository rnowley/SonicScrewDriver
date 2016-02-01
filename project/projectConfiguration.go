package project

type ProjectConfiguration struct {
	Name                 string
	Version              string
	Description          string
	DestinationDirectory string
	ClassPath            []string
	SourceFiles          []string
	SourceVersion        string
	DebuggingInformation []string
	//Resources   []Resource
}
