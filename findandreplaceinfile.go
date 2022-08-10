package main

import (
	"findandreplaceinfile/files"
	"fmt"
	"os"
)

func main() {
	occ, lines, err := files.FindReplaceFile("example.txt", "hello", "goodbye")
	if err != nil {
		fmt.Printf("Error: ''%v''\n", err)
		os.Exit(1)
	}

	fmt.Printf("=> Occurences:%v lines:%v\n", occ, lines)
}
