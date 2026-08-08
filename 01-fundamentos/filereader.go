package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "No valid arguments written")
		os.Exit(1)
	}
	fileName := os.Args[1]
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not Read the file.\n Error: %v\n", err)
		os.Exit(1)
	}
	os.Stdout.Write(data)
}
