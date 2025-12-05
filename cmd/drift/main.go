package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 || os.Args[1] != "render" {
		fmt.Println("Usage: drift render <input file> <output file>")
		os.Exit(1)
	}

	// TODO: Implement file reading and rendering logic
	inputFile := os.Args[2]
	outputFile := os.Args[3]
	fmt.Println("Rendering", inputFile, "to", outputFile)
}
