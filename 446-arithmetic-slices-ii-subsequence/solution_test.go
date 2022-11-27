package arithmeticSlicesIISubsequence

import "testing"

func Test_numberOfArithmeticSlices(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				nums: []int{2, 4, 6, 8, 10},
			},
			want: 7,
		},
		{
			name: "example 2",
			args: args{
				nums: []int{7, 7, 7, 7, 7},
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := numberOfArithmeticSlices(tt.args.nums); got != tt.want {
				t.Errorf("numberOfArithmeticSlices() = %v, want %v", got, tt.want)
			}
		})
	}
}
