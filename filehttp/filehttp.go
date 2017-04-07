package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

var basedir string

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Usage: go run server.go [basedir]")
		os.Exit(1)
	}

	basedir = os.Args[1]

	fmt.Printf("serving content: %s\n", basedir)

	http.HandleFunc("/", serverHandle)

	serverHost := "localhost:7777"

	fmt.Printf("server on http://%s\n", serverHost)

	http.ListenAndServe(serverHost, nil)
}

func serverHandle(res http.ResponseWriter, req *http.Request) {

	file, err := os.Open(path.Join(basedir, req.URL.Path))

	defer file.Close()

	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}

	_, err = io.Copy(res, file)

	if err != nil {
		log.Println("serve content:", err)
	}
}
