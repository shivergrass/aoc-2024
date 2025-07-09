package main

import (
	"fmt"
	"os"
)



func readInputFile() string {
	res, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("ReadInputFile() %s\n", err)
		os.Exit(1)
	}
	
	return string(res)
}

func Read() string {
	return readInputFile()
}