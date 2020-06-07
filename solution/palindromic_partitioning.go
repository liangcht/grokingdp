package solution

import (
	"fmt"
	"grokingdp/mathutil"
)

func palindromicPartitioning(input string) int {
	// return recursivePalindromicPartitioning(input, 0, len(input)-1)
	n := len(input)
	if n < 1 {
		return 0
	}
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for i := range dp {
		dp[i][i] = true
	}
	for l := 2; l <= n; l++ {
		for start := 0; start <= n-l; start++ {
			end := start + l - 1
			if input[start] == input[end] {
				if start == end-1 || dp[start+1][end-1] {
					dp[start][end] = true
				}
			}
		}
	}

	cut := make([]int, n)
	for i := n - 2; i >= 0; i-- {
		cut[i] = n - i - 1
		for end := i; end < n; end++ {
			if dp[i][end] {
				if end == n-1 {
					cut[i] = 0
				} else {
					cut[i] = mathutil.MinInt(cut[i], 1+cut[end+1])
				}
			}
		}
	}
	return cut[0]
}

func recursivePalindromicPartitioning(input string, start int, end int) int {
	if start == end {
		return 0
	}
	minPartition := end - start
	for i := start; i <= end; i++ {
		if isPalindrome(input, start, i) {
			if i == end {
				return 0
			}
			minPartition = mathutil.MinInt(minPartition, 1+recursivePalindromicPartitioning(input, i+1, end))
		}
	}
	return minPartition
}

func isPalindrome(input string, start int, end int) bool {
	for start < end {
		if input[start] != input[end] {
			return false
		}
		start++
		end--
	}
	return true
}

// TestPalindromicPartitioning is used for testing
func TestPalindromicPartitioning() {
	fmt.Println(palindromicPartitioning("abdbca"))
	fmt.Println(palindromicPartitioning("cddpd"))
	fmt.Println(palindromicPartitioning("pqr"))
	fmt.Println(palindromicPartitioning("pp"))
	fmt.Println(palindromicPartitioning("madam"))
	fmt.Println(palindromicPartitioning("cabababcbc"))
	fmt.Println(palindromicPartitioning("abababcb"))
}

// Wrong recursive function
func wrongRecursivePalindromicPartitioning(input string, start int, end int) int {
	if start == end {
		return 0
	}
	if start == end-1 {
		if input[start] == input[end] {
			return 0
		}
		return 1
	}
	cut := wrongRecursivePalindromicPartitioning(input, start+1, end-1)
	if input[start] == input[end] && cut == 0 {
		return 0
	}
	cut += 2
	cut = mathutil.MinInt(cut, 1+wrongRecursivePalindromicPartitioning(input, start+1, end))
	cut = mathutil.MinInt(cut, 1+wrongRecursivePalindromicPartitioning(input, start, end-1))
	return cut
}
