package solution

import (
	"fmt"
	"grokingdp/mathutil"
)

func longestAlternatingSubsequence(nums []int) int {
	increase := make([]int, len(nums))
	decrease := make([]int, len(nums))
	increase[0] = 1
	decrease[0] = 1
	for i := 0; i < len(nums); i++ {
		maxIncreasing := 0
		maxDecreasing := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && increase[j] > maxIncreasing {
				maxIncreasing = increase[j]
			}
			if nums[i] < nums[j] && decrease[j] > maxDecreasing {
				maxDecreasing = decrease[j]
			}
		}
		increase[i] = maxDecreasing + 1
		decrease[i] = maxIncreasing + 1
	}
	return mathutil.MaxInt(increase[len(nums)-1], decrease[len(nums)-1])
}

// TestLongestAlternatingSubsequence is used for testing
func TestLongestAlternatingSubsequence() {
	fmt.Println(longestAlternatingSubsequence([]int{1, 2, 3, 4}))
	fmt.Println(longestAlternatingSubsequence([]int{3, 2, 1, 4}))
	fmt.Println(longestAlternatingSubsequence([]int{1, 3, 2, 4}))
}
