package main

import (
	"fmt"
	"os"
	"path"
)

var total int64

func main() {

	if len(os.Args[1:]) != 1 {
		fmt.Fprintln(os.Stderr, "Usage: du.go [file | dir]")
		return
	}

	filename := os.Args[1]

	duRescure(filename)

	fmt.Printf("total disk use: %d\n", total)
}

func duRescure(filename string) {

	file, err := os.Open(filename)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	file.Close() // defer file.Close()

	ff, err := file.Readdir(-1)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	for _, fileInfo := range ff {
		realPath := path.Join(filename, fileInfo.Name())
		if fileInfo.IsDir() {
			duRescure(realPath)
		} else {

			size := fileInfo.Size()

			total += size

			fmt.Printf("%s : %d\n", realPath, size)
		}

	}

}
