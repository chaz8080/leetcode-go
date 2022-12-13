package twosumii

import (
	"reflect"
	"testing"
)

func Test_twosumii(t *testing.T) {
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
			name: "example 3",
			args: args{
				nums: []int{5, 25, 75},
				k:    100,
			},
			want: [2]int{2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twosumii(tt.args.nums, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twosumii() = %v, want %v", got, tt.want)
			}
		})
	}
}
