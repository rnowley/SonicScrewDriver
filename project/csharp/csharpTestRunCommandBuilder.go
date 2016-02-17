package csharp

func GetCSharpRunTestCommand(configuration CSharpProject) MonoCommand {
	command := NewDefaultMonoCommand()

	command.ExecutableName = configuration.TestProject.TestRunner

	argumentCount := len(configuration.TestProject.RunArguments)

	for i := 0; i < argumentCount; i++ {
		command.RunArguments = append(command.RunArguments, configuration.TestProject.RunArguments[i])
	}

	return command
}