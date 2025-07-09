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
	return res
}

func processInput(inp []string) [2][]int {
	res_0 := make([]int, 0)
	res_1 := make([]int, 0)

	for _, v := range inp {
		if v == "" {
			break
		}
		nums_str := strings.Split(v, LIST_SPLIT)

		i, err := strconv.Atoi(nums_str[0])
		if err != nil {
			fmt.Printf("ProcessInput() nums_str[0] %s\n", v)
			fmt.Printf("ProcessInput() nums_str[0] %s\n", err)
			os.Exit(1)
		}
		res_0 = append(res_0, i)

		i, err = strconv.Atoi(nums_str[1])
		if err != nil {
			fmt.Printf("ProcessInput() nums_str[1] %s\n", v)
			fmt.Printf("ProcessInput() nums_str[1] %s\n", err)
			os.Exit(1)
		}
		res_1 = append(res_1, i)
	}

	return [2][]int{
		res_0,
		res_1,
	}
}

func Read() [2][]int {
	str := readInputFile()
	split_str := splitInput(str)
	res := processInput(split_str)

	return res
}
