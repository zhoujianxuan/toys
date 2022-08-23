package dynamic

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	l := len(nums)
	dp := make([][2]int, l)

	dp[0][0] = 0
	dp[0][1] = nums[0]

	for i, num := range nums {
		if i == 0 {
			continue
		}

		dp[i][0] = max(dp[i-1][0], dp[i-1][1])
		dp[i][1] = dp[i-1][0] + num
	}

	return max(dp[l-1][0], dp[l-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
