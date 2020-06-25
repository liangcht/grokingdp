package solution

import (
	"fmt"
	"grokingdp/mathutil"
	"math"
)

func shortestCommonSuperSequence(s1 string, s2 string) int {
	// return shortestCommonSuperSequenceRecur(s1, s2, len(s1)-1, len(s2)-1)
	dp := make([][]int, len(s1)+1)
	for i := range dp {
		dp[i] = make([]int, len(s2)+1)
	}
	dp[0][0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i][0] = i
	}
	for j := 1; j < len(s2)+1; j++ {
		dp[0][j] = j
	}
	for i := 1; i < len(s1)+1; i++ {
		for j := 1; j < len(s2)+1; j++ {
			count := math.MaxInt32
			if s1[i-1] == s2[j-1] {
				count = 1 + dp[i-1][j-1]
			}
			count = mathutil.MinInt(count, 1+dp[i-1][j])
			count = mathutil.MinInt(count, 1+dp[i][j-1])
			dp[i][j] = count
		}
	}
	return dp[len(s1)][len(s2)]
}

func shortestCommonSuperSequenceRecur(s1 string, s2 string, i1, i2 int) int {
	if i1 == -1 && i2 == -1 {
		return 0
	}
	if i1 == -1 {
		return i2 + 1
	}
	if i2 == -1 {
		return i1 + 1
	}
	count := math.MaxInt32
	if s1[i1] == s2[i2] {
		count = 1 + shortestCommonSuperSequenceRecur(s1, s2, i1-1, i2-1)
	}
	count = mathutil.MinInt(count, 1+shortestCommonSuperSequenceRecur(s1, s2, i1-1, i2))
	count = mathutil.MinInt(count, 1+shortestCommonSuperSequenceRecur(s1, s2, i1, i2-1))
	return count
}

// TestShortestCommonSuperSequence is used for testing
func TestShortestCommonSuperSequence() {
	fmt.Println(shortestCommonSuperSequence("abcf", "bdcf"))
	fmt.Println(shortestCommonSuperSequence("dynamic", "programming"))
}
