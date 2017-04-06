package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	for _, fileName := range os.Args[1:] {

		file, err := os.Open(fileName)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		_, err = io.Copy(os.Stdout, file)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		err = file.Close()

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
