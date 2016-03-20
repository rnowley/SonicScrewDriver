package scala

// ScalaProject is a struct for holding the information required for
// building or running a Scala project.
// Name is the name of the project.
// Version is used for holding the project version number.
// Description holds a description of the project.
// Language holds the name of the programming language that the project is written in.
// SourceDirectory holds the root directory for the project's source code.
// DestinationDirectory contains the directory that build artifacts are going to be placed in.
// ClassPath holds a list of items that are part of the class path.
// SourceFiles holds a list of all of the files to be compiled into the project.
// DebuggingInformation holds information about the debugging information to be generated.
// Target specifies the backend to use.
type ScalaProject struct {
	Name                 string
	Version              string
	Description          string
	Language             string
	SourceDirectory      string
	DestinationDirectory string
	ClassPath            []string
	NoWarnings           bool
	Optimise             bool
	SourceFiles          []string
	DebuggingInformation string
	Encoding             string
	Target               string
}
