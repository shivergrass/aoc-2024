package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var LIST_SPLIT string = "   "

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

func processInput(inp []string) [][]int { // level -> entry
	var res [][]int

	for _, v := range inp {
		level_str := strings.Split(v, " ")
		var level []int

		for _, el := range level_str {
			entry, err := strconv.Atoi(el)
			if err != nil {
				fmt.Printf("processInput() %s\n", err)
				os.Exit(1)
			}

			level = append(level, entry)
		}

		res = append(res, level)
	}

	return res
}

func Read() [][]int {
	str := readInputFile()
	split_str := splitInput(str)
	res := processInput(split_str)

	return res
}
