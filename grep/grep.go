package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

var invertMatch = flag.Bool("v", false, "Selected lines are those not matching any of the specified patterns.")
var caseInsensitive = flag.Bool("i", false, "Perform case insensitive matching.")
var showLineNumber = flag.Bool("n", false, "Each output line is preceded by its relative line number in the file, starting at line 1.")

func main() {

	// fmt.Println(re.MatchString("AB"))

	flag.Parse()
	args := flag.Args()

	// fmt.Println("inverMatch", *invertMatch)
	// fmt.Println("caseInsensitive", *caseInsensitive)
	// fmt.Println("showLineNumber", *showLineNumber)
	// fmt.Println("args", args)

	// os.Stdin

	// if len(args) >= 1 {

	// } else {
	// 		scanner := bufio.NewScanner(os.Stdin)
	// 	linenum := 0
	// 	for scanner.Scan() {
	// 		linenum++
	// 	fmt.Println(scanner.Text())
	// }

	if (*showLineNumber != true) && (len(args) == 0) {
		fmt.Fprintln(os.Stderr, "Usage: grep.go [flag] pattern [file]")
		flag.Usage()
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

			grep(file, filename, re, len(args))

			file.Close()
		}

	} else {

		grep(os.Stdin, "", re, len(args))
	}

}

func grep(reader io.Reader, filename string, re *regexp.Regexp, arglen int) {
	linenum := 0
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		linenum++
		line := scanner.Text()

		if *showLineNumber {
			line = strconv.Itoa(linenum) + ":" + line
		}

		if *showLineNumber && (filename != "") {
			line = filename + ":" + line
		}

		switch true {
		case arglen == 0 && *showLineNumber:
			fallthrough
		case arglen >= 1 && *invertMatch && !re.MatchString(line):
			fallthrough
		case arglen >= 1 && !*invertMatch && re.MatchString(line):
			fmt.Println(line)
		default:
			break
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
