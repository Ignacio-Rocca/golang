package main

import "fmt"

//Given a sorted array nums, remove the duplicates in-place such that each element appears only once and returns the new length.
//Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

func main () {
	nums := []int{1,1,2,2,2,2,2,3,3,4,4,5,5}

	// nums is passed in by reference. (i.e., without making a copy)
	length := removeDuplicates(nums)

	// any modification to nums in your function would be known by the caller.
	// using the length returned by your function, it prints the first len elements.
	for i := 0; i < length; i++ {
		fmt.Println(nums[i])
	}
}

func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums)-1; {
		if nums[i] == nums[i+1] {
			nums = append(nums[:i], nums[i+1:]...)
		} else {
			i++
		}
	}

	return len(nums)
}
