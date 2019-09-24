package main

import "fmt"

func main() {
	//fmt.Printf("%v\n", isHuiwen("aa"))
	fmt.Println(longestPalindrome("a"))
}

// 暴力破解法
func longestPalindrome(s string) string {
	num, ans, l := 0, "", len(s)
	for i := 0; i < l; i++ {
		for j := i + 1; j <= l; j++ {
			if isPalindromicString(s[i:j]) {
				if (j - i) > num {
					num = j - i
					ans = s[i:j]
				}
			}
		}
	}
	return ans
}

func isPalindromicString(s string) bool {
	for i, j := 0, len(s); i < j/2; i++ {
		if s[i] != s[j-1-i] {
			return false
		}
	}
	return true
}

// 动态规划
func longestPalindrome2(s string) string {
	num, ans, l := 0, "", len(s)
	for i := 0; i < l; i++ {
		for j := i + 1; j <= l; j++ {
			if isPalindromicString(s[i:j]) {
				if (j - i) > num {
					num = j - i
					ans = s[i:j]
				}
			}
		}
	}
	return ans
}
