# SonicScrewDriver
SonicScrewDriver is a build tool for a number of languages that uses a project definition that is defined in a JSON file. The aim of this tool is to allow projects to be easily compiled without needing to use an IDE or to understand complex build tools.

The target is audience is small to medium sized projects and those users that prefer to use a text editor and the command line for developing those projects or need to because they are developing in a resource constrained environment where an IDE is not appropriate, such as a single board computer such as a Raspberry PI.

The programming languages supported by this tool are:
* Java
* Scala
* Kotlin
* C#

Currently this tool has been written to run on linux. Support for other operating systems will depend on demand.

## Install

To install this tool clone the GIT repository.
then
```
cd SonicScrewDriver
```
then run the command:
```
go get
```
to retrieve the dependencies then run
```
go install sonicScrewDriver.go
```
to install the tool.

## Usage
```
sonicScrewDriver {OPTIONS} mode

The available options are:
  -- verbose, This is used to display detailed output of the build process, it is defaulted to off.

The available modes are:
  build: Build the project.
  build-all: Build the project and the init tests.
  build-tests: Build unit tests.
  docs: Generate project documentation
  run: Run the project.
  run-tests: Run unit tests.
```
