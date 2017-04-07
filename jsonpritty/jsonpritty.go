package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var data interface{}

func main() {

	if len(os.Args[1:]) != 1 {
		fmt.Println("Usage: jsonpritty [file]")
		return
	}

	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "\t")
	encoder.Encode(data)
}
