package project

// Command defines an interface for objects to implement if they want
// to act as a command.
type Command interface {

	// GetCommandName is a function for returning the name of the command.
	GetCommandName() string

	// GetDestiantionDirectory is a function for returning the
	// destination directory part of the command.
	GetDestinationDirectory() string

	// GenerateArgumentList is a function for returning the arguments
	// used by the command.
	GenerateArgumentList() []string
}
