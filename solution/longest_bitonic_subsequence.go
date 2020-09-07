package solution

import (
	"fmt"
	"grokingdp/mathutil"
)

func longestBintonicSubsequence(nums []int) int {
	countPerIndex := increasingSubsequenceCountPerIndex(nums)
	countPerIndexReverse := increasingSubsequenceCountPerIndexReverse(nums)
	max := 1
	for i := 0; i < len(nums); i++ {
		count := countPerIndex[i] + countPerIndexReverse[i]
		max = mathutil.MaxInt(max, count)
	}
	return max - 1
}

func increasingSubsequenceCountPerIndex(nums []int) []int {
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 0; i < len(nums); i++ {
		count := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j] > count {
				count = dp[j]
			}
		}
		dp[i] = count + 1
	}
	return dp
}

func increasingSubsequenceCountPerIndexReverse(nums []int) []int {
	dp := make([]int, len(nums))
	dp[len(nums)-1] = 1
	for i := len(nums) - 1; i >= 0; i-- {
		count := 0
		for j := len(nums) - 1; j > i; j-- {
			if nums[i] > nums[j] && dp[j] > count {
				count = dp[j]
			}
		}
		dp[i] = count + 1
	}
	return dp
}

// TestLongestBintonicSubsequence is used for testing
func TestLongestBintonicSubsequence() {
	fmt.Println(longestBintonicSubsequence([]int{4, 2, 3, 6, 10, 1, 12}))
	fmt.Println(longestBintonicSubsequence([]int{4, 2, 5, 9, 7, 6, 10, 3, 1}))
}
