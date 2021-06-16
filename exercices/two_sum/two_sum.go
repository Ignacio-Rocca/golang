package two_sum

import "fmt"

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	min := func() int {
		n := 1
		val := -2
		for i := 1; i <= 32; i++ {
			n *= val
		}
		return n
	}

	fmt.Println(min())

	for i := 0; i < len(nums); i++ {
		numsMap[nums[i]] = i
	}

	for i :=0; i < len(nums); i++ {
		complement := target - nums[i]
		if j, ok := numsMap[complement]; ok && j != i{
			return []int{i, j}
		}
	}

	return []int{}
}