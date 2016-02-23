package java

import "testing"

func TestGetJavaBuildCommandDeprecatedFalse(t *testing.T) {
	var configuration JavaProject
	configuration.Name = "Test Java Project"
	configuration.Version = "1.0.0"
	configuration.Description = "A project for unit testing."
	configuration.Language = "java"
	configuration.DestinationDirectory = "./build/"
	configuration.ClassPath = []string{"./lib/a.jar", "./lib/b.jar"}
}
