package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage digest.go [md5 | hash256 | hash516] ")
		return
	}

	var hash hash.Hash

	switch strings.ToLower(os.Args[1]) {
	case "md5":
		hash = md5.New()

	case "256":
		hash = sha256.New()

	case "512":
		hash = sha512.New()
	default:
		fmt.Fprintln(os.Stderr, "Usage digest.go [md5 | hash256 | hash516] ")
		return
	}

	_, err := io.Copy(hash, os.Stdin)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	var digest []byte
	digest = hash.Sum(digest)

	fmt.Println(hex.EncodeToString(digest))
}
