package solution

import (
	"testing"
)

func Test_exist(t *testing.T) {
	type args struct {
		board [][]byte
		word  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				board: [][]byte{
					[]byte("ABCE"),
					[]byte("SFCS"),
					[]byte("ADEE"),
				},
				word: "ABCCED",
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				board: [][]byte{
					[]byte("ABCE"),
					[]byte("SFCS"),
					[]byte("ADEE"),
				},
				word: "SEE",
			},
			want: true,
		},
		{
			name: "example 3",
			args: args{
				board: [][]byte{
					[]byte("ABCE"),
					[]byte("SFCS"),
					[]byte("ADEE"),
				},
				word: "ABCB",
			},
			want: false,
		},
		{
			name: "example 4",
			args: args{
				board: [][]byte{
					[]byte("ABCE"),
					[]byte("SFES"),
					[]byte("ADEE"),
				},
				word: "ABCESEEEFS",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := exist(tt.args.board, tt.args.word); got != tt.want {
				t.Errorf("exist() = %v, want %v", got, tt.want)
			}
		})
	}
}
