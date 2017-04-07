package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	for _, fileName := range os.Args[1:] {

		err := cat(fileName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

	}
}

func cat(fileName string) (err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer file.Close()

	_, err = io.Copy(os.Stdout, file)
	return err
}
