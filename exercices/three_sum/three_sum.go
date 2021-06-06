package three_sum

import (
	"sort"
)

// Given an array S of n integers, are there elements a, b, c in S such that a + b + c = 0?
// Find all unique triplets in the array which gives the sum of zero.

// Note: The solution set must not contain duplicate triplets.

// For example, given array S = [-1, 0, 1, 2, -1, -4],
// A solution set is:
// [
//   [-1, 0, 1],
//   [-1, -1, 2]
// ]

func getTripletsThatSumZero(data []int) [][3]int {
	var result [][3]int
	n := len(data)
	if n < 3 {
		return [][3]int{}
	}

	sort.Ints(data)
	for i := 0; i < n-2; i++ {
		for j := i+1; j < n -1; j++ {
			for k := j+1; k < n; k++ {
				v1 := data[i]
				v2 := data[j]
				v3 := data[k]

				if v1 + v2 	+ v3 == 0 {
					result = append(result, [3]int{v1,v2,v3})
				}
			}
		}
	}

	return unique(result)
}

func unique(tripletsArray [][3]int) [][3]int {
	list := [][3]int{}

	for _, triplet := range tripletsArray {
		hasEqual := false
		for _, t := range list {
			if t == triplet {
				hasEqual = true
				break
			}
		}
		if !hasEqual {
			list = append(list, triplet)
		}
	}
	return list
}
