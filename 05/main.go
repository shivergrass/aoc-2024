package main

import (
	"fmt"
)

func main() {
	// part_1()
	part_2()
}

func part_1() {
	pairs, sequences := Read()

	compiled_pairs := compilePairs(pairs)
	total := 0
	for _, seq := range sequences {
		if validateSequence(seq, compiled_pairs) {
			mid_idx := len(seq) / 2
			total += seq[int(mid_idx)]
		}
	}

	fmt.Printf("total: %d\n", total)
}

func part_2() {
	pairs, sequences := Read()
	compiledPairs := compilePairs(pairs)

	for _, seq :=  range sequences {
		countDisplacement(seq, compiledPairs)
	}
}

// gonna use map to map out the precedences because i haven't warmed up in it
// i would use array instead because of the O(1) access if i have to do this seriously
func compilePairs(pairs [][2]int) map[int][]int {
	res := make(map[int][]int, len(pairs))

	for _, pair := range pairs {
		res[pair[0]] = append(res[pair[0]], pair[1])
	}

	return res
}

func validateSequence(seq []int, mapping map[int][]int) bool {
	
	res := true
	seq_len := len(seq)
	for i := seq_len-1; i >= 0; i-- {
		
		currentVal := seq[i]
		for pre_idx := i-1; pre_idx >= 0; pre_idx-- {
			preVal := seq[pre_idx]

			// precedence of current value
			preMap := mapping[currentVal]

			for _, v := range preMap {
				res = res && !(preVal == v)
				if !res { return false }
			}
		}
	}

	return true
}

func countDisplacement(seq []int, mapping map[int][]int) map[int]int {
	
	res := make(map[int]int, 0)
	seq_len := len(seq)
	for i := seq_len-1; i >= 0; i-- {
		currentVal := seq[i]
		res[currentVal] = -1

		for pre_idx := i-1; pre_idx >= 0; pre_idx-- {
			preVal := seq[pre_idx]

			// precedence of current value
			preMap := mapping[currentVal]

			for _, v := range preMap {
				if preVal != v { continue } // correctly placed
				res[currentVal] += 1

				break
			}
		}
	}

	return res
}

func checkIfPrecedes(num int, target int, mapping map[int][]int) bool {
	precendenceMap := mapping[num]

	for _, v := range precendenceMap {
		if target != v { continue }
		return true
	}

	return false
}

// misorderedIdx should start from the last index
func separateMisorderedIdx(seq []int, misorderedIdx []int) ([]int, []int) { // correct order, misorder
	ordered := make([]int, len(seq))
	copy(ordered, seq)

	misordered := make([]int, 0)


	for _, misIdx := range misorderedIdx {
		misordered = append(misordered, seq[misIdx])
		
		firstOrdered := ordered[:misIdx]
		secondOrdered := make([]int, 0)
		if misIdx + 1 < len(ordered) {
			secondOrdered = ordered[misIdx+1:]
		}

		ordered = append(firstOrdered, secondOrdered...)
	}

	return ordered, misordered
}