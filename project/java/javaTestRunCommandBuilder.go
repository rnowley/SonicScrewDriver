package java

// http://docs.oracle.com/javase/7/docs/technotes/tools/windows/java.html
// https://github.com/junit-team/junit/wiki/Getting-started

func GetJavaRunTestCommand(configuration JavaProject) JavaCommand {
	command := NewDefaultJavaCommand()

	classPathCount := len(configuration.ClassPath)

	// Add the current directory to the command's class path.
	command.ClassPath = append(command.ClassPath, ".")

	// Add the project's class path to the command's class path.
	for i := 0; i < classPathCount; i++ {
		command.ClassPath = append(command.ClassPath, configuration.ClassPath[i])
	}

	classPathCount = len(configuration.TestProject.ClassPath)

	// Add the test project's class path to the command's class path
	for i := 0; i < classPathCount; i++ {
		command.ClassPath = append(command.ClassPath, configuration.TestProject.ClassPath[i])
	}

	command.MainClass = configuration.TestProject.MainClass

	argumentCount := len(configuration.TestProject.RunArguments)

	for i := 0; i < argumentCount; i++ {
		command.RunArguments = append(command.RunArguments, configuration.TestProject.RunArguments[i])
	}

	return command
}
