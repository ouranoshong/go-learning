package main

import "net/http"
import "fmt"
import "os"
import "io"

func main() {

	if len(os.Args[1:]) != 2 {
		fmt.Println("Usage: wget [url] [file]")

		return
	}

	url := os.Args[1]
	filename := os.Args[2]

	resp, err := http.Get(url)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	defer resp.Body.Close()

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0660)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	_, err = io.Copy(file, resp.Body)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
