package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
)

var invertMatch = flag.Bool("v", false, "Selected lines are those not matching any of the specified patterns.")
var caseInsensitive = flag.Bool("i", false, "Perform case insensitive matching.")
var showLineNumber = flag.Bool("n", false, "Each output line is preceded by its relative line number in the file, starting at line 1.")

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		flag.PrintDefaults()
		return
	}

	pattern := args[0]

	if *caseInsensitive {
		pattern = "(?i)" + pattern
	}

	re, _ := regexp.Compile(pattern)

	if len(args) > 1 {

		for _, filename := range args[1:] {
			file, err := os.Open(filename)

			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue
			}

			grep(file, filename, re)

			file.Close()
		}

	} else {

		grep(os.Stdin, "", re)
	}

}

func grep(reader io.Reader, filename string, re *regexp.Regexp) (err error) {
	linenum := 0
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		linenum++
		line := scanner.Text()

		match := re.MatchString(line)

		if *invertMatch {
			match = !match
		}

		if match {
			if *showLineNumber {
				if filename != "" {
					fmt.Printf("%s:%d:", filename, linenum)
				} else {
					fmt.Printf("%d:", linenum)
				}
			}

			fmt.Println(line)
		}
	}

	return scanner.Err()
}
