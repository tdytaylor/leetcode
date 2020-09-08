package issue_0094

import (
	. "github.com/tdytaylor/leetcode/structure"
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "test case1",
			args: createTree(),
			want: []int{1, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createTree() struct{ root *TreeNode } {
	var name struct{ root *TreeNode }
	node3 := &TreeNode{Val: 3, Right: nil, Left: nil}
	node2 := &TreeNode{Val: 2, Right: nil, Left: node3}
	node1 := &TreeNode{Val: 1, Right: node2, Left: nil}
	name.root = node1
	return name
}
