package issue_0234

import (
	. "github.com/tdytaylor/leetcode/structure"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	type args struct {
		head *ListNode
	}

	l1 := &ListNode{
		Val: 1,
	}
	l1.Append(2).Append(2).Append(1)

	l2 := &ListNode{
		Val: 1,
	}
	l2.Append(2)

	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{head: l1}, want: true},
		{name: "test2", args: args{head: l2}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
