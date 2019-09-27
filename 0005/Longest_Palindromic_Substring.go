package main

import "fmt"

func main() {
	//fmt.Printf("%v\n", isHuiwen("aa"))
	fmt.Println(longestPalindrome2("abb"))
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
	len, maxLen, start := len(s), 0, 0

	var dp = make([][]int, len)
	for i := 0; i < len; i++ {
		dp[i] = make([]int, len)
		dp[i][i] = 1
		if maxLen < 2 {
			maxLen = 1
		}
		if i < len-1 && s[i] == s[i+1] {
			dp[i][i+1] = 1
			start = i
			maxLen = 2
		}
	}

	for i := 3; i <= len; i++ { // 每次检索的字符串长度，依次递增
		for j := 0; j+i-1 < len; j++ {
			if s[j] == s[j+i-1] && dp[j+1][j+i-2] == 1 {
				dp[j][j+i-1] = 1
				if i >= maxLen {
					start = j
					maxLen = i
				}
			}
		}
	}
	return s[start : start+maxLen]
}
