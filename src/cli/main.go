package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: drift render <input-file> --output <output-file> [--engine <engine>]")
		os.Exit(1)
	}

	// TODO: Implement full CLI argument parsing.
	command := os.Args[1]
	if command == "render" {
		// TODO: Implement the rendering logic.
		fmt.Println("Rendering diagram...")
	} else {
		fmt.Printf("Unknown command: %s\n", command)
		os.Exit(1)
	}
}
