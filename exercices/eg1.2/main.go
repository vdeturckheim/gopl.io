package main

import (
	"os"
	"fmt"
)

func main() {
	for rank, arg := range os.Args[1:] {
		fmt.Println(rank, arg)
	}
}
