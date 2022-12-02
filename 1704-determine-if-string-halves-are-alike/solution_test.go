package determineIfStringHalvesAreAlike

import "testing"

func Test_halvesAreAlike(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "example 1",
			args: args{
				s: "book",
			},
			want: true,
		},
		{
			name: "example 2",
			args: args{
				s: "textbook",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := halvesAreAlike(tt.args.s); got != tt.want {
				t.Errorf("halvesAreAlike() = %v, want %v", got, tt.want)
			}
		})
	}
}
