package issue_0888

import (
	"reflect"
	"testing"
)

func Test_fairCandySwap(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fairCandySwap(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fairCandySwap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fairCandySwap2(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{a: []int{1, 1}, b: []int{2, 2}}, want: []int{1, 2}},
		{name: "test2", args: args{a: []int{1, 2}, b: []int{2, 3}}, want: []int{1, 2}},
		{name: "test3", args: args{a: []int{2}, b: []int{1, 3}}, want: []int{2, 3}},
		{name: "test4", args: args{a: []int{1, 2, 5}, b: []int{2, 4}}, want: []int{5, 4}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fairCandySwap2(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fairCandySwap2() = %v, want %v", got, tt.want)
			}
		})
	}
}
