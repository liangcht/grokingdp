package solution

import (
	"fmt"
	"grokingdp/mathutil"
)

func houseThief(wealth []int) int {
	// return findMaxWealthRecursive(wealth, len(wealth)-1
	dp := make([]int, len(wealth))
	for i := range dp {
		dp[i] = 0
	}
	dp[0] = wealth[0]
	dp[1] = wealth[1]
	for i := 2; i < len(wealth); i++ {
		dp[i] = dp[i-1]
		dp[i] = mathutil.MaxInt(dp[i], wealth[i]+dp[i-2])
	}
	return dp[len(wealth)-1]
}

func findMaxWealthRecursive(wealth []int, index int) int {
	if index < 2 {
		return wealth[index]
	}
	return mathutil.MaxInt(findMaxWealthRecursive(wealth, index-1), wealth[index]+findMaxWealthRecursive(wealth, index-2))
}

// TestHouseThief is used for testing
func TestHouseThief() {
	fmt.Println(houseThief(([]int{2, 5, 1, 3, 6, 2, 4})))
	fmt.Println(houseThief(([]int{2, 10, 14, 8, 1})))
}
