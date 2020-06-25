package solution

import (
	"fmt"
	"grokingdp/mathutil"
)

func minimumDeletionToSorted(nums []int) int {
	n := len(nums)
	dp := make([]int, n)

	maxSeq := 0
	for i := 0; i < n; i++ {
		curMax := 0
		for j := 0; j < i; j++ {
			if nums[i] >= nums[j] && dp[j] > curMax {
				curMax = dp[j]
			}
		}
		dp[i] = curMax + 1
		maxSeq = mathutil.MaxInt(maxSeq, dp[i])
	}
	return n - maxSeq
}

// TestMinimumDeletionToSorted is used for testing
func TestMinimumDeletionToSorted() {
	fmt.Println(minimumDeletionToSorted([]int{4, 2, 3, 6, 10, 1, 12}))
	fmt.Println(minimumDeletionToSorted([]int{-4, 10, 3, 7, 15}))
	fmt.Println(minimumDeletionToSorted([]int{3, 2, 1, 0}))
}
