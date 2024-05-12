package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in file - Hello world One"

	file, err := os.Create("./mylocalFile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}

	fmt.Println(length)
	defer file.Close()

	ReadFile("./mylocalFile.txt")
}

func ReadFile(file string) {
	dataInBytes, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(dataInBytes))
}
