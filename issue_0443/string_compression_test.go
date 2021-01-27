package issue_0443

import (
	"reflect"
	"testing"
)

func Test_compress(t *testing.T) {
	type args struct {
		chars []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compress(tt.args.chars); got != tt.want {
				t.Errorf("compress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_compressToStr2(t *testing.T) {
	type args struct {
		chars []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{
			chars: []byte{'a', 'a', 'n', 'b', 'b', 'c', 'c', 'c'},
		}, want: []byte{'a', '2', 'n', 'b', '2', 'c', '3'}},
		{name: "test2", args: args{
			chars: []byte{'a', 'n', 'b', 'b', 'c'},
		}, want: []byte{'a', 'n', 'b', '2', 'c'}},
		{name: "test3", args: args{
			chars: []byte{'a', 'a', 'a', 'b', 'b', 'a', 'a'},
		}, want: []byte{'a', '3', 'b', '2', 'a', '2'}},
		{name: "test4", args: args{
			chars: []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'},
		}, want: []byte{'a', '2', 'b', '2', 'c', '3'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compressToStr2(tt.args.chars); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("compressToStr2() = %v, want %v", got, tt.want)
			}
		})
	}
}
