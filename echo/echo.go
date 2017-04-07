package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// for _, arg := range(os.Args[1:]) {
	// 	fmt.Print(arg + " ")
	// }

	// fmt.Println();

	fmt.Println(strings.Join(os.Args[1:], " "))
}
