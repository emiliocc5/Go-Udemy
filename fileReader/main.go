package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type consoleWriter struct{}

func main() {
	f := os.Args[1]
	file, err := os.Open(f)
	if err != nil {
		log.Fatal(err)
	}
	cw := consoleWriter{}
	_, err = io.Copy(cw, file)
	if err != nil {
		log.Fatal(err)
	}
}

func (consoleWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	return len(bs), nil
}
