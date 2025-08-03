package main

import (
	"fmt"
	"os"
	"strconv"
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

func splitInput(inp string) ([]string, []string) {
	split_inp := strings.Split(inp, "\n")

	var pairs []string
	var sequences []string

	do_pair := true

	for _, v := range split_inp {
		if v == "" {
			do_pair = false
			continue
		} else if do_pair {
			pairs = append(pairs, v)
		} else {
			sequences = append(sequences, v)
		}
	}

	return pairs, sequences
}

func convertPair(pairs []string) [][2]int { // input: ["12|18"]
	var res [][2]int

	for _, str_pair := range pairs {
		split_str_pair := strings.Split(str_pair, "|")

		var num_pair [2]int

		for i := 0; i < 2; i++ {
			num, err := strconv.Atoi(split_str_pair[i])
			if err != nil {
				fmt.Printf("convertPair(pairs []string) %s\n", err.Error())
				os.Exit(1)
			}

			num_pair[i] = num
		}

		res = append(res, num_pair)
	}

	return res
}

func splitNumOrder(input []string) [][]int { 
	// input: ["44,76,93,64,66"]
	// output: [[44,76,93,64,66]]
	var res [][]int

	for _, v :=  range input {
		str_arr := strings.Split(v, ",")
		var nums []int

		for _, str_num := range str_arr {
			num, err := strconv.Atoi(str_num)
			if err != nil {
				fmt.Printf("splitNumOrder(input []string) %s\n", err.Error())
				os.Exit(1)
			}

			nums = append(nums, num)
		}

		res = append(res, nums)
	}

	return res
}

func Read() ([][2]int, [][]int) {
	str_inp := readInputFile()
	str_pair, str_seq := splitInput(str_inp)

	pairs := convertPair(str_pair)
	sequences := splitNumOrder(str_seq)

	return pairs, sequences
}