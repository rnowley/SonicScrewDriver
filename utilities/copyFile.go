package copyFile

import (
    "io"
    "os"
)

func copyFile(destination, source string) error {
    s, err := os.Open(source)

    if err != nil {
        return err
    }

    defer s.Close()
    d, err := os.Create(destination)

    if err != nil {
        return err
    }

    if _, err := io.Copy(d, s); err != nil {
        d.Close()
        return err
    }

    return d.Close()
}