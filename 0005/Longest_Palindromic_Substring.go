package main

import "fmt"

func main() {
	//fmt.Printf("%v\n", isHuiwen("aa"))
	fmt.Println(longestPalindrome("aa"))
}

// 暴力破解法
func longestPalindrome(s string) string {
	num, ans, l := 0, "", len(s)
	if l == 0 || l == 1 {
		return s
	}
	for i := 0; i < l; i++ {
		for j := i + 1; j <= l; j++ {
			if isHuiwen(s[i:j]) {
				if (j - i) > num {
					num = j - i
					ans = s[i:j]
				}
			}
		}
	}
	return ans
}

func isHuiwen(s string) bool {
	for i, j := 0, len(s); i < j/2; i++ {
		if s[i] != s[j-1-i] {
			return false
		}
	}
	return true
}
