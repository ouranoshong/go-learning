package main

import (
	"archive/tar"
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

	trfile, err := gzip.NewReader(gzfile)

	if err != nil {
		return
	}

	defer trfile.Close()

	file := tar.NewReader(trfile)

	_, err = io.Copy(os.Stdout, file)
	return err
}
