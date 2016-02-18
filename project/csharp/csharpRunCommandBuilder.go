package csharp

import "fmt"

func GetCSharpRunCommand(configuration CSharpProject) MonoCommand {
	command := NewDefaultMonoCommand()
	command.ExecutableName = fmt.Sprintf("%s%s.exe", configuration.DestinationDirectory, configuration.OutputFilename)

	argumentCount := len(configuration.RunArguments)

	for i := 0; i < argumentCount; i++ {
		command.RunArguments = append(command.RunArguments, configuration.RunArguments[i])
	}

	return command
}
