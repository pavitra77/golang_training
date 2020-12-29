package main

import (
	"fmt"
	"log"
	"os"
)

func CreateFile(filename, text string) {

	fmt.Printf("Writing file in Golang\n")

	file, err := os.Create(filename)

	if err != nil {
		log.Fatalf("failed to create a file: %s", err)
	}

	defer file.Close()
	len, err := file.WriteString(text)
	if err != nil {
		log.Fatalf("failed to write a file: %s", err)
	}

	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nLength: %d bytes", len)
}
func main() {

	CreateFile()
}