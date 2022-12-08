package countprimes

import "testing"

func Test_countPrimes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				n: 10,
			},
			want: 4,
		},
		{
			name: "example 2",
			args: args{
				n: 0,
			},
			want: 0,
		},
		{
			name: "example 3",
			args: args{
				n: 1,
			},
			want: 0,
		},
		{
			name: "example 4",
			args: args{
				n: 500,
			},
			want: 95,
		},
		{
			name: "example 5",
			args: args{
				n: 5_000_000,
			},
			want: 348513,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPrimes(tt.args.n); got != tt.want {
				t.Errorf("countPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
