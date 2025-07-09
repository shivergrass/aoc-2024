package main

import (
	"fmt"
	"os"
)

const TARGET_WORD string = "XMAS"

func get_func_ptrs() [8]func(word string, x int, y int, width int, height int)([][2]int, bool) {
	return [8]func(word string, x int, y int, width int, height int) ([][2]int, bool){
		directionN,
		directionNE,
		directionE,
		directionSE,
		directionS,
		directionSW,
		directionW,
		directionNW,
	}
}

func main() {
	part_1()
}

func part_1() {
	inp := Read()

	count := 0

	for y, _ := range inp {
		for x, _ := range inp[0] {
			sequences := generateSequences(TARGET_WORD, [2]int{x, y}, len(inp), len(inp[0]))
			for _, seq := range sequences {
				is_valid := checkSequence(seq, TARGET_WORD, inp)
				if !is_valid { continue }
				count++
			}
		}
	}

	fmt.Printf("%d\n", count)
}

func generateSequences(word string, coord [2]int, inp_width int, inp_height int) [][][2]int {
	funcs := get_func_ptrs()
	var res [][][2]int

	for _, f :=  range funcs {
		sequence, is_valid := f(word, coord[0], coord[1], int(inp_width), int(inp_height))
		if !is_valid { continue }

		res = append(res, sequence)
	}

	return res
}

func checkSequence(coords [][2]int, word string, inp []string) bool {
	if len(word) != len(coords) {
		fmt.Printf("len(word) != len(coords)\n")
		os.Exit(1)
	}

	for i, letter := range []byte(word) {
		coord := coords[i]
		target := inp[coord[1]][coord[0]]

		if letter != target {
			return false
		}
	}

	return true
}

func directionN(word string, x int, y int, width int, height int) ([][2]int, bool) { // indexes, is within bounds
	if !checkBound(y - (len(word) - 1), height) {
		return [][2]int{}, false
	}

	var res [][2]int
	for i, _ := range word {
		res = append(res, [2]int{x, y-i})
	}

	return res, true
}

func directionNE(word string, x int, y int, width int, height int) ([][2]int, bool) { // indexes, is within bounds
	if !checkBound(x + (len(word) - 1), width) || !checkBound(y - (len(word) - 1), height) {
		return [][2]int{}, false
	}

	var res [][2]int
	for i, _ := range word {
		res = append(res, [2]int{x+i, y-i})
	}

	return res, true
}

func directionE(word string, x int, y int, width int, height int) ([][2]int, bool) { // indexes, is within bounds
	if !checkBound(x + (len(word) - 1), width) {
		return [][2]int{}, false
	}

	var res [][2]int
	for i, _ := range word {
		res = append(res, [2]int{x+i, y})
	}

	return res, true
}

func directionSE(word string, x int, y int, width int, height int) ([][2]int, bool) { // indexes, is within bounds
	if !checkBound(x + (len(word) - 1), width) || !checkBound(y + (len(word) - 1), height) {
		return [][2]int{}, false
	}

	var res [][2]int
	for i, _ := range word {
		res = append(res, [2]int{x+i, y+i})
	}

	return res, true
}

func directionS(word string, x int, y int, width int, height int) ([][2]int, bool) { // indexes, is within bounds
	if !checkBound(y + (len(word) - 1), height) {
		return [][2]int{}, false
	}

	var res [][2]int
	for i, _ := range word {
		res = append(res, [2]int{x, y+i})
	}

	return res, true
}

func directionSW(word string, x int, y int, width int, height int) ([][2]int, bool) { // indexes, is within bounds
	if !checkBound(x - (len(word) - 1), width) || !checkBound(y + (len(word) - 1), height) {
		return [][2]int{}, false
	}

	var res [][2]int
	for i, _ := range word {
		res = append(res, [2]int{x-i, y+i})
	}

	return res, true
}

func directionW(word string, x int, y int, width int, height int) ([][2]int, bool) { // indexes, is within bounds
	if !checkBound(x - (len(word) - 1), width) {
		return [][2]int{}, false
	}

	var res [][2]int
	for i, _ := range word {
		res = append(res, [2]int{x-i, y})
	}

	return res, true
}

func directionNW(word string, x int, y int, width int, height int) ([][2]int, bool) { // indexes, is within bounds
	if !checkBound(x - (len(word) - 1), width) || !checkBound(y - (len(word) - 1), height) {
		return [][2]int{}, false
	}

	var res [][2]int
	for i, _ := range word {
		res = append(res, [2]int{x-i, y-i})
	}

	return res, true
}

func checkBound(idx int, exc_bound int) bool {
	return idx >= 0 && idx < exc_bound
}



