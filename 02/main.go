package main

import (
	"fmt"
	"math"
)

func main() {
	// inp := Read()
	// fmt.Printf("%v\n", inp[len(inp)-1])
	part_1()
	part_2()
}

func part_2() {
	inp := Read()

	safe_count := 0
	for _, v := range inp {
		if !checkSafeWithDampener(v, true) { continue }
		safe_count++
	}

	fmt.Printf("safe_count %d\n", safe_count)

}

func part_1() {
	inp := Read()

	safe_count := 0
	for _, v := range inp {
		if !checkSafe(v) { continue }
		safe_count++
	}

	fmt.Printf("safe_count %d\n", safe_count)
}

func checkSafe(level []int) bool {
	iter_count := len(level) - 1
	sign_int := 0
	for i := 0; i < iter_count; i++ {
		new_sign_int, is_within_range := checkFluctuation(level[i], level[i+1], 3)

		if sign_int == 0 {
			sign_int = new_sign_int
		} 
		if new_sign_int != sign_int || new_sign_int == 0 || !is_within_range {
			return false
		}
	}

	return true
}

func checkSafeWithDampener(level []int, use_dampener bool) bool {
	iter_count := len(level) - 1
	sign_int := 0
	for i := 0; i < iter_count; i++ {
		new_sign_int, is_within_range := checkFluctuation(level[i], level[i+1], 3)

		if sign_int == 0 {
			sign_int = new_sign_int
		} 
		if new_sign_int != sign_int || new_sign_int == 0 || !is_within_range {
			if !use_dampener {
				return false
			}
			
			return checkSafeWithDampener(level[i:], false) || i+2 >= len(level) || checkSafeWithDampener(append([]int{level[i]}, level[i+2:]...), false)
		}
	}

	return true
}

func checkFluctuation(e0 int, e1 int, fluctuation_range int) (int, bool) { // return: sign_int, is_within_range
	diff := float64(e1 - e0)
	sign_int := Clamp(diff, -1, 1)

	return int(sign_int), math.Abs(diff) <= float64(fluctuation_range)
}




