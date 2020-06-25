package solution

import (
	"fmt"
)

func maximumSumIncreasingSubsequence(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	maxSum := 0
	for i := 1; i < n; i++ {
		dp[i] = 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j] > dp[i] {
				dp[i] = dp[j]
			}
		}
		dp[i] += nums[i]
		if dp[i] > maxSum {
			maxSum = dp[i]
		}
	}
	return maxSum
}

// TestMaximumSumIncreasingSubsequence is used for testing
func TestMaximumSumIncreasingSubsequence() {
	fmt.Println(maximumSumIncreasingSubsequence([]int{4, 1, 2, 6, 10, 1, 12}))
	fmt.Println(maximumSumIncreasingSubsequence([]int{-4, 10, 3, 7, 15}))
	fmt.Println(maximumSumIncreasingSubsequence([]int{4, 10, 4, 3, 8, 9}))
}
