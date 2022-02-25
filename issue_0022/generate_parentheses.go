package issue_0022

import (
	"container/list"
	"strings"
)

// 括号生成
// Example 1:
// Input: n = 3
// Output: ["((()))","(()())","(())()","()(())","()()()"]
func generateParenthesis(n int) []string {
	var li = list.New()
	var str strings.Builder
	backtrack(li, &str, 0, 0, n)
	return nil
}

func backtrack(list *list.List, str *strings.Builder, open int, close int, n int) {
	if str.Len() == 2*n {
		list.PushBack(str.String())
	}

	if open < n {
		str.WriteByte('(')
		backtrack(list, str, open+1, close, n)
	}

	if close < open {
		str.WriteByte(')')
		backtrack(list, str, open, close+1, n)
	}
}
