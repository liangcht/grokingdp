package solution

import (
	"fmt"
)

func stringsInterleaving(m, n, p string) bool {
	if len(m)+len(n) != len(p) {
		return false
	}
	dp := make([][]bool, len(m)+1)
	for i := range dp {
		dp[i] = make([]bool, len(n)+1)
	}
	dp[0][0] = true
	for i := 1; i < len(m)+1; i++ {
		if m[i-1] == p[i-1] {
			dp[i][0] = dp[i-1][0]
		}
	}
	for i := 1; i < len(n)+1; i++ {
		if n[i-1] == p[i-1] {
			dp[0][i] = dp[0][i-1]
		}
	}

	for i := 1; i < len(m)+1; i++ {
		for j := 1; j < len(n)+1; j++ {
			dp[i][j] = (m[i-1] == p[i+j-1] && dp[i-1][j]) || (n[j-1] == p[i+j-1] && dp[i][j-1])
		}
	}
	return dp[len(m)][len(n)]
}

// TestStringsInterleaving is used for testing
func TestStringsInterleaving() {
	/*
		fmt.Println(stringsInterleaving("abd", "cef", "abcdef"))
		fmt.Println(stringsInterleaving("abd", "cef", "adcbef"))
		fmt.Println(stringsInterleaving("abc", "def", "abdccf"))
		fmt.Println(stringsInterleaving("abcdef", "mnop", "mnaobcdepf"))
	*/
	fmt.Println(stringsInterleaving("ab", "bc", "babc"))
}
