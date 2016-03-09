package utilities

import (
	"fmt"
	"io"
	"os"
)

// CopyFile copies a file from its destination location to the source location
// specified in its arguments.
func CopyFile(destination, source string) error {
	s, err := os.Open(source)

	if err != nil {
		fmt.Println(err)
		return err
	}

	defer s.Close()
	d, err := os.Create(destination)

	if err != nil {
		fmt.Println(err)
		return err
	}

	if _, err := io.Copy(d, s); err != nil {
		d.Close()
		fmt.Println(err)
		return err
	}

	return d.Close()
}

// EnsurePathExists is used to ensure that all folders on the specified path exist.
// If they do not then these directories are created.
func EnsurePathExists(directoryPath string) error {
	_, err := os.Stat(directoryPath)

	if err != nil {
		err = os.MkdirAll(directoryPath, 0777)

		if err != nil {
			fmt.Println(err)
			return err
		}

		return nil
	}

	return err
}
