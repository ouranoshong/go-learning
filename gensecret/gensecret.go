package main

import "crypto/rand"
import "fmt"
import "os"
import "encoding/base64"
import "strconv"

func main() {
	var count int

	if len(os.Args) != 2 {
		count = 32
	} else {
		count, _ = strconv.Atoi(os.Args[1])

	}

	randSlice := make([]byte, count)
	_, err := rand.Read(randSlice)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	fmt.Println(base64.StdEncoding.EncodeToString(randSlice))
}
