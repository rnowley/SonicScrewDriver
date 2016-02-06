package utilities

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(destination, source string) error {
	fmt.Println(destination)
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

	fmt.Println("Finished copy")
	return d.Close()
}

func EnsurePathExists(directoryPath string) error {
	_, err := os.Stat(directoryPath)

	if err != nil {
		err = os.MkdirAll(directoryPath, 0777)

		if err != nil {
			fmt.Println(err)
			return err
		}

		fmt.Println("Created path.")
		return nil
	}

	fmt.Println(err)
	return err
}
