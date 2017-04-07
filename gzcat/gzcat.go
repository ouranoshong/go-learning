package main

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	for _, fileName := range os.Args[1:] {

		err := zgCat(fileName)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

	}
}

func zgCat(fileName string) (err error) {
	gzfile, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer gzfile.Close()

	file, err := gzip.NewReader(gzfile)

	if err != nil {
		return
	}

	_, err = io.Copy(os.Stdout, file)
	return err
}
