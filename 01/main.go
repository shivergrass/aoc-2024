package main

import (
	"fmt"
	"math"
	"os"
)

/*
	overall, this approach is heavy on the space
*/

/*
	on part 1
	i could do something funnier instead of repetitively renewing the slices of input
	instead of [2][]int, i could've done something like [2][][2]int32

	in [2]int32, the additional int would store the info on whether this element "exists" or not
	like an option type
	i didn't think of it earlier, so- oh well
*/

func main() {
	part_1()
	part_2()
}

func part_1() {
	var inp [2][]int = Read()
	// fmt.Printf("%d\n", len(inp))
	res := 0.

	for len(inp[0]) != 0 {
		if len(inp[0]) != len(inp[1]) {
			fmt.Println("len(inp[0]) != len(inp[1])")
			os.Exit(1)
		}

		count := len(inp[0])

		idx_sl_0 := 0
		idx_sl_1 := 0
		
		// find smallest
		for i := 1; i < count; i++ {
			idx_sl_0 = updateIndexMinimum(inp[0], idx_sl_0, i)
			idx_sl_1 = updateIndexMinimum(inp[1], idx_sl_1, i)
		}

		// add to res
		res += math.Abs(float64(inp[0][idx_sl_0] - inp[1][idx_sl_1]))

		// remove element
		inp[0] = withoutIndex(inp[0], idx_sl_0)
		inp[1] = withoutIndex(inp[1], idx_sl_1)
	}

	fmt.Printf("%f\n", res)

}

func withoutIndex(sl []int, target int) []int {
	res := append(sl[:target], sl[target+1:]...)

	return res
}

func updateIndexMinimum(sl []int, idx_0 int, idx_1 int) int {
	if sl[idx_1] >= sl[idx_0] {
		return idx_1
	}
	return idx_0
}

/*
comment on part 2
i will be using array to store the count of elements instead of dictionary
*/

func part_2() {
	// read input
	inp := Read()

	// find floor and ceiling
	idx_extremes := findFloorAndCeil(inp[1])
	min_val := inp[1][idx_extremes[0]]
	max_val := inp[1][idx_extremes[1]]
	distance := max_val - min_val

	// allocate
	tallies := make([]int, distance + 1)

	// count
	for i := min_val; i <= max_val; i++ {
		el_count := countElementPresence(inp[1], i)
		tallies[i - min_val] = el_count
	}

	// sum of presence
	sum := 0
	for _, v := range inp[0] {
		target := v - min_val
		if v > max_val || v < min_val { 
			continue 
		}
		sum += tallies[target] * v
	}

	fmt.Printf("%d\n", sum)
}

func findFloorAndCeil(sl []int) [2]int {
	count := len(sl)
	idx_floor := 0
	idx_ceil := 0
	for i := 1; i < count; i++ {
		if sl[i] < sl[idx_floor] {
			idx_floor = i
		} else if  sl[i] > sl[idx_ceil] {
			idx_ceil = i
		}
	}

	return [2]int {idx_floor, idx_ceil}
}

func countElementPresence(sl []int, el int) int {
	res := 0
	for _, v := range sl {
		if v != el { continue }
		
		res++
	}

	return res
}



