package main

import "fmt"

// given 2 arrays and 1 target we need to find what values reach the target
// assume that the arrays has the same length

// Bruth force solution O(n^2)
func findTheTarget_01(a1 []int, a2 []int, target int) bool {

	for i := 0; i < len(a1); i++ {
		for j := 0; j < len(a1); j++ {
			if a1[i] + a2[j] == target {
				return true
			}
		}
	}
	return false

}

// Thinking the solution O(n)
func findTheTarget_02(a1 []int, a2 []int, target int) bool {
	map1 := make(map[int]bool)
	map2 := make(map[int]bool)

	for i := 0; i < len(a1); i++ {
		map1[a1[i]] = true
		map2[a2[i]] = true
	}

	for k, _ := range map1 {
		if _, ok := map2[target - k]; ok {
			return true
		}
	}

	return false

}

func main() {
	fmt.Println(findTheTarget_01([]int{1,2,3}, []int{1,2,3}, 3)) // true
	fmt.Println(findTheTarget_01([]int{2,2,5}, []int{1,1,1}, 4)) // false
	fmt.Println(findTheTarget_02([]int{1,2,3}, []int{1,2,3}, 3)) // true
	fmt.Println(findTheTarget_02([]int{2,2,5}, []int{1,1,1}, 4)) // false
}