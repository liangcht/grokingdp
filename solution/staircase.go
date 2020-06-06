package solution

import (
	"fmt"
)

func stairCase(n int) int {

	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2

	for i := 3; i < n+1; i++ {
		dp[i] = dp[i-1]
		if i > 1 {
			dp[i] += dp[i-2]
		}
		if i > 2 {
			dp[i] += dp[i-3]
		}
	}
	return dp[n]
}

// TestStairCase is used for testing
func TestStairCase() {
	fmt.Println(stairCase((3)))
	fmt.Println(stairCase((4)))
	fmt.Println(stairCase((5)))
}
