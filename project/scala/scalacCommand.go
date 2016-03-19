package scala

// ScalacCommand provides all of the information to generate
// a call to the Scala compiler command.
type ScalacCommand struct {
	CommandName          string
	SourceDirectory      string
	DestinationDirectory string
	ClassPath            []string
	Deprecation          bool
}

func NewDefaultScalacCommand() ScalacCommand {
	var command ScalacCommand
	command.CommandName = "scalac"
	command.SourceDirectory = "./src/"
	command.DestinationDirectory = "./build/"
	command.ClassPath = make([]string, 0, 10)
	return command
}

// GetCommandName is a method on a ScalacCommand which accesses the name of the command
// to be run.
func (command ScalacCommand) GetCommandName() string {
	return command.CommandName
}

// GetDestinationDirectory is a method on a ScalacCommand which accesses the Destination Directory command
// to be run.
func (command ScalacCommand) GetDestinationDirectory() string {
	return command.DestinationDirectory
}

// GenerateArgumentList is a method which returns a slice of strings containing
// the arguments to use when running the scalac compiler command.
func (command ScalacCommand) GenerateArgumentList() []string {
	argumentArray := make([]string, 0)
	argumentArray = append(argumentArray, "-d", command.DestinationDirectory)

	if command.Deprecation {
		argumentArray = append(argumentArray, "-deprecation")
	}

	return argumentArray
}
