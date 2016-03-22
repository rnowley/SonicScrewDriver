package scala

func GetScalaRunCommand(configuration ScalaProject) ScalaCommand {
	command := NewDefaultScalaCommand()

	classPathCount := len(configuration.ClassPath)

	// Add the current directory to the command's class path.
	command.ClassPath = append(command.ClassPath, ".")

	// Add the project's class path to the command's class path.
	for i := 0; i < classPathCount; i++ {
		command.ClassPath = append(command.ClassPath, configuration.ClassPath[i])
	}

	command.MainClass = configuration.MainClass

	argumentCount := len(configuration.RunArguments)

	for i := 0; i < argumentCount; i++ {
		command.RunArguments = append(command.RunArguments, configuration.RunArguments[i])
	}

	return command
}
