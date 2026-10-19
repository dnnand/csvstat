package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: csvstat <file.csv>")
		os.Exit(2)
	}

	path := os.Args[1]

	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot open file:", err)
		os.Exit(2)
	}
	defer f.Close()

	header, err := csv.NewReader(f).Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, "cannot read header", err)
		os.Exit(2)
	}

	fmt.Println("columns:", len(header))
	for i, name := range header {
		fmt.Println(i, name)
	}

}
