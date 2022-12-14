package threesumclosest

import (
	"reflect"
	"testing"
)

func Test_threesumclosest(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "example 1",
			args: args{
				[]int{-1, 0, 1, 2, -1, -4},
			},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			name: "example 2",
			args: args{
				[]int{0, 1, 1},
			},
			want: [][]int{},
		},
		{
			name: "example 3",
			args: args{
				[]int{0, 0, 0},
			},
			want: [][]int{
				{0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threesumclosest(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threesumclosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
