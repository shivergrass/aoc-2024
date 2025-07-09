package main

import (
	"fmt"
	"os"
	"strings"
)

func readInputFile() string {
	res, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Printf("ReadInputFile() %s\n", err)
		os.Exit(1)
	}
	
	return string(res)
}

func splitInput(inp string) []string {
	var res []string = strings.Split(inp, "\n")
	res = res[:len(res)-1]
	return res
}

func Read() []string {
	return splitInput(readInputFile())
}