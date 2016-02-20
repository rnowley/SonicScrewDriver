package kotlin

// GetKotlinRunCommand creates a new KotlinCommand based on the configuration
// passed in as the argument to this function.
func GetKotlinRunCommand(configuration KotlinProject) KotlinCommand {
	command := NewDefaultKotlinCommand()

	// Add the current directory to the class path.
	command.ClassPath = append(command.ClassPath, ".")

	classPathCount := len(configuration.ClassPath)

	// Add the project's class path to the command's class path.
	for i := 0; i < classPathCount; i++ {
		command.ClassPath = append(command.ClassPath, configuration.ClassPath[i])
	}

	if configuration.JarFile != "" {
		command.JarFile = configuration.JarFile
	}

	argumentCount := len(configuration.RunArguments)

	for i := 0; i < argumentCount; i++ {
		command.RunArguments = append(command.RunArguments, configuration.RunArguments[i])
	}

	return command
}
