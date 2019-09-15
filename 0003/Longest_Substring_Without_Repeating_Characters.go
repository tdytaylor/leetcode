package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abba"))
}

// 滑动窗口
func lengthOfLongestSubstring(s string) int {
	set, ans, start := make(map[rune]int, len(s)), 0, 0
	for i, val := range []rune(s) {
		if lastVal, ok := set[val]; ok && start < lastVal {
			start = lastVal + 1
		}
		if (i - start + 1) > ans {
			ans = i - start + 1
		}
		set[val] = i
	}
	return ans
}
