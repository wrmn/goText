package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	openFile := os.Args
	file, err := os.Open(openFile[1])
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%q\n", data[:count])
}
