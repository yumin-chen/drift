package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: drift render <input file> -o <output file>")
		os.Exit(1)
	}

	renderCmd := flag.NewFlagSet("render", flag.ExitOnError)
	outputFile := renderCmd.String("o", "", "output file")

	renderCmd.Parse(os.Args[2:])

	if renderCmd.NArg() != 1 {
		fmt.Println("Usage: drift render <input file> -o <output file>")
		os.Exit(1)
	}

	inputFile := renderCmd.Arg(0)

	// TODO: Implement file reading and rendering logic
	fmt.Println("Rendering", inputFile, "to", *outputFile)
}
