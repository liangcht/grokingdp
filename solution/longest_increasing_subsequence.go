package solution

import (
	"fmt"
	"grokingdp/mathutil"
)

func longestIncreasingSubsequenceClean(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1
	maxNum := 0
	for i := 1; i < n; i++ {
		dp[i] = 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j] > dp[i] {
				dp[i] = dp[j]
			}
		}
		dp[i]++
		if dp[i] > maxNum {
			maxNum = dp[i]
		}
	}
	return maxNum
}

func longestIncreasingSubsequenceONSpace(nums []int) int {
	//return longestIncreasingSubsequenceRecur(nums, 0, -1)
	n := len(nums)
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i; j >= 0; j-- {
			if j == 0 || nums[i] > nums[j-1] {
				dp[i%2][j] = 1 + dp[(i+1)%2][i+1]
			}
			dp[i%2][j] = mathutil.MaxInt(dp[i%2][j], dp[(i+1)%2][j])
		}
	}
	return dp[0][0]
}

func longestIncreasingSubsequence(nums []int) int {
	//return longestIncreasingSubsequenceRecur(nums, 0, -1)
	n := len(nums)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			if j == 0 || nums[i] > nums[j-1] {
				dp[i][j] = 1 + dp[i+1][i+1]
			}
			dp[i][j] = mathutil.MaxInt(dp[i][j], dp[i+1][j])
		}
	}
	return dp[0][0]
}

func longestIncreasingSubsequenceRecur(nums []int, index int, prevIndex int) int {
	if index == len(nums) {
		return 0
	}
	count := 0
	if prevIndex == -1 || nums[index] > nums[prevIndex] {
		count = 1 + longestIncreasingSubsequenceRecur(nums, index+1, index)
	}
	count = mathutil.MaxInt(count, longestIncreasingSubsequenceRecur(nums, index+1, prevIndex))
	return count
}

// TestLongestIncreasingSubsequence is used for testing
func TestLongestIncreasingSubsequence() {
	fmt.Println(longestIncreasingSubsequenceClean([]int{4, 2, 3, 6, 10, 1, 12}))
	fmt.Println(longestIncreasingSubsequenceClean([]int{-4, 10, 3, 7, 15}))
	fmt.Println(longestIncreasingSubsequenceClean([]int{4, 10, 4, 3, 8, 9}))
}
