// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileNames:= make(map[string]map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileNames, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fileNames, arg)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			files := ""
			for fileName := range fileNames[line]{
				files += fileName + " "
			}
			fmt.Printf("%d\t%s\t%s\n", n, line, files)
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileNames map[string]map[string]bool, file string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if fileNames[input.Text()] == nil {
			fileNames[input.Text()] = make(map[string]bool)
		}
		fileNames[input.Text()][file] = true
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
