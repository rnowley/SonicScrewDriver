package project

// GetProjectBuilder is a factory function that returns an object that implements the
// ProjectBuilder interface. The type of this ProjectBuilder is determined by the
// configurationFile that is passed in.
func GetProjectBuilder(configurationFile []byte, arguments Arguments) ProjectBuilder {
	return ProjectBuilder
}

// Use language in configuration and "build/test from arguments"

// GetProjectLanguage is a function for retrieving the value that
// determines the programming language that the project is written
// in.
func GetProjectLanguage(file []byte) string {
	var projectLanguage project.ProjectLanguage

	if err := json.Unmarshal(file, &projectLanguage); err != nil {
		panic(err)
	}

	return projectLanguage.Language
}
