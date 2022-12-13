package twosumclosestk

import (
	"reflect"
	"testing"
)

func Test_twosumclosestk(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want [2]int
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{2, 7, 11, 15},
				k:    9,
			},
			want: [2]int{1, 2},
		},
		{
			name: "example 2",
			args: args{
				nums: []int{2, 3, 4},
				k:    6,
			},
			want: [2]int{1, 3},
		},
		{
			name: "example 3",
			args: args{
				nums: []int{-1, 0},
				k:    -1,
			},
			want: [2]int{1, 2},
		},
		{
			name: "example 4",
			args: args{
				nums: []int{-1, 0},
				k:    -1,
			},
			want: [2]int{1, 2},
		},
		{
			name: "example 5",
			args: args{
				nums: []int{-10, 5, 3, 56, 10, 0},
				k:    1,
			},
			want: [2]int{1, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twosumclosestk(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twosumclosestk() = %v, want %v", got, tt.want)
			}
		})
	}
}
