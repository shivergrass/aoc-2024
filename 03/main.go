package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dlclark/regexp2"
)

const RG0 string = `(?<=mul\()[0-9]+,[0-9]+(?=\))`
const RG1 string = `(?<=mul\()[0-9]+,[0-9]+(?=\))|do\(\)|don't\(\)`

type Op = int
const (
	Do Op = iota
	Dont
	Num
)

func main() {
	part_1()
	part_2()
}

func part_1() {
	rg := regexp2.MustCompile(RG0, 0)
	inp := Read()

	res, err := rg.FindStringMatch(inp)
	sum := 0
	for {
		if err != nil || res == nil {
			fmt.Printf("rg.FindStringMatch(inp) %s\n", err)
			break
		}

		// fmt.Printf("%s\n", res.Length)
		pair := splitString(res.Capture.String())
		sum += pair[0] * pair[1]

		res, err = rg.FindNextMatch(res)
	}

	fmt.Printf("%d\n", sum)
}

func splitString(pair string) [2]int {
	split_str := strings.Split(pair, ",")

	var res [2]int 
	var err interface{}

	res[0], err = strconv.Atoi(split_str[0])
	if err != nil {
		fmt.Printf("splitString() strconv.Atoi res[0] %s\n", err)
	}

	res[1], err = strconv.Atoi(split_str[1])
	if err != nil {
		fmt.Printf("splitString() strconv.Atoi res[1] %s\n", err)
	}

	return res
}

func checkOpCode(op string) Op { // false == is Don't
	if op == "do()" {
		return Do
	} else if op == "don't()" {
		return Dont
	}

	return Num
}

func part_2() {
	rg := regexp2.MustCompile(RG1, 0)
	inp := Read()

	res, err := rg.FindStringMatch(inp)
	sum := 0
	var op Op = Do

	for {
		if err != nil || res == nil {
			fmt.Printf("rg.FindStringMatch(inp) %s\n", err)
			break
		}

		new_op := checkOpCode(res.Capture.String())
		if new_op != Num {
			op = new_op
		} else {
			if op == Dont {
				goto ESC
			}

			pair := splitString(res.Capture.String())
			sum += pair[0] * pair[1]
		}


		ESC: 
		res, err = rg.FindNextMatch(res)
	}

	fmt.Printf("%d\n", sum)

}