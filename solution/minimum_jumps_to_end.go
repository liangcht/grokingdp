package solution

import (
	"fmt"
	"grokingdp/mathutil"
	"math"
)

func minimumJumpToEnd(jumps []int) int {
	n := len(jumps)
	if n <= 1 {
		return 0
	}
	dp := make([]int, n)
	for i := range dp {
		dp[i] = math.MaxInt32
	}

	dp[0] = 0
	for i := 0; i < n; i++ {
		end := i + 1
		for end <= i+jumps[i] && end < n {
			dp[end] = mathutil.MinInt(dp[end], dp[i]+1)
			end++
		}
	}
	return dp[n-1]
}

// TestMinimumJumpToEnd is used for testing
func TestMinimumJumpToEnd() {
	fmt.Println(minimumJumpToEnd(([]int{2, 1, 1, 1, 4})))
	fmt.Println(minimumJumpToEnd(([]int{1, 1, 3, 6, 9, 3, 0, 1, 3})))
}
