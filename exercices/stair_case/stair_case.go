package stair_case

// simple solution assumption that you can go up by one or two steps only
func NumWays(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return NumWays(n-1) + NumWays(n-2)
}

// simple solution with DP assumption that you can go up by one or two steps only
func NumWaysDP_BottomUp(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	nums := []int{1, 1}

	for i := 2; i <= n; i++ {
		nums = append(nums, nums[i-1] + nums[i-2])
	}

	return nums[n]
}

// complex solution receiving by parameter the set of steps that you can do
func NumWaysComplex(n int, steps []int) int {
	if n == 0 {
		return 1
	}
	nums := []int{1}
	for i := 1; i <= n; i++ {
		total := 0
		for _, s := range steps {
			if i - s >= 0 {
				total += nums[i - s]
			}
		}
		nums = append(nums, total)
	}
	return nums[n]
}
