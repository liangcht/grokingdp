package solution

import (
	"fmt"
)

func subsequencePatternMatchingDpSlow(str, pat string) int {
	n := len(str)
	dp := make([]int, n)
	//for i := range dp {
	//	dp[i] = make([]int, n)
	//}
	count := 0
	for i := 0; i < len(pat); i++ {
		for j := 0; j < n; j++ {
			if str[j] == pat[i] {
				max := 0
				for k := 0; k < j; k++ {
					if dp[k] > max {
						max = dp[k]
					}
					if dp[k] == len(pat)-1 {
						count++
					}
				}
				dp[j] = max + 1
			}
		}
	}
	return count
}

func subsequencePatternMatchingRecursive(str, pat string) int {
	return subsequencePatternMatchingRecur(str, pat, 0, 0)
}

func subsequencePatternMatchingRecur(str, pat string, strStart, patStart int) int {
	if patStart == len(pat) {
		return 1
	}
	if strStart == len(str) {
		return 0
	}

	count := subsequencePatternMatchingRecur(str, pat, strStart+1, patStart)
	if str[strStart] == pat[patStart] {
		count += subsequencePatternMatchingRecur(str, pat, strStart+1, patStart+1)
	}
	return count
}

func subsequencePatternMatchingDpFast(str, pat string) int {
	if len(str) == 0 || len(pat) == 0 {
		return 0
	}

	dp := make([][]int, len(str)+1)
	for i := range dp {
		dp[i] = make([]int, len(pat)+1)
	}

	for i := 0; i < len(str)+1; i++ {
		dp[i][0] = 1
	}

	for i := 1; i < len(str)+1; i++ {
		for j := 1; j < len(pat)+1; j++ {
			dp[i][j] = dp[i-1][j]
			if str[i-1] == pat[j-1] {
				dp[i][j] += dp[i-1][j-1]
			}
		}
	}
	return dp[len(str)][len(pat)]
}

// TestSubsequencePatternMatching is used for testing
func TestSubsequencePatternMatching() {
	fmt.Println(subsequencePatternMatchingDpFast("baxmx", "ax"))
	fmt.Println(subsequencePatternMatchingDpFast("tomorrow", "tor"))
}
