package kotlin

import "fmt"

func GetKotlinRunTestCommand(configuration KotlinProject) KotlinCommand {
	command := NewDefaultKotlinCommand()

	// Add the current directory to the command's class path.
	command.ClassPath = append(command.ClassPath, ".")

	command.JarFile = ""

	// Add the main project's jar to the class path.
	command.ClassPath = append(command.ClassPath, fmt.Sprintf("%s%s",
		configuration.DestinationDirectory, configuration.OutputFilename))

	// Add the test project's jar to the class path.
	command.ClassPath = append(command.ClassPath, fmt.Sprintf("%s%s",
		configuration.TestProject.DestinationDirectory,
		configuration.TestProject.OutputFilename))

	classPathCount := len(configuration.ClassPath)

	// Add the project's class path to the command's class path.
	for i := 0; i < classPathCount; i++ {
		command.ClassPath = append(command.ClassPath,
			configuration.ClassPath[i])
	}

	classPathCount = len(configuration.TestProject.ClassPath)

	// Add the test project's class path to the command's class path
	for i := 0; i < classPathCount; i++ {
		command.ClassPath = append(command.ClassPath,
			configuration.TestProject.ClassPath[i])
	}

	command.MainClass = configuration.TestProject.TestRunner

	argumentCount := len(configuration.TestProject.RunArguments)

	for i := 0; i < argumentCount; i++ {
		command.RunArguments = append(command.RunArguments,
			configuration.TestProject.RunArguments[i])
	}

	return command

}
