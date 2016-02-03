package project

type Command interface {
	GetCommandName() string
	GetDestinationDirectory() string
	GenerateArgumentList() []string
}
