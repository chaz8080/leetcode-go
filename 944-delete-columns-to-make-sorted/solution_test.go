package deletecolumnstomakesorted

import "testing"

func Test_minDeletionSize(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{
				strs: []string{
					"cba",
					"daf",
					"ghi",
				},
			},
			want: 1,
		},
		{
			name: "example 2",
			args: args{
				strs: []string{
					"a",
					"b",
				},
			},
			want: 0,
		},
		{
			name: "example 3",
			args: args{
				strs: []string{
					"zyx",
					"wvu",
					"tsr",
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDeletionSize(tt.args.strs); got != tt.want {
				t.Errorf("minDeletionSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
