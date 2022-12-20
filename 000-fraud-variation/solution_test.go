package solution

import "testing"

func Test_detectFraud(t *testing.T) {
	type args struct {
		grid [][]string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Fail",
			args: args{
				[][]string{
					{"P", "F", "F", "F"},
					{"P", "_", "F", "_"},
					{"_", "P", "F", "P"},
					{"P", "P", "F", "P"},
				},
			},
			want: FAIL,
		},
		{
			name: "Pass",
			args: args{
				[][]string{
					{"P", "F", "F", "F"},
					{"P", "_", "P", "_"},
					{"P", "P", "F", "P"},
					{"P", "P", "_", "P"},
				},
			},
			want: PASS,
		},
		{
			name: "Needs more info",
			args: args{
				[][]string{
					{"P", "F", "F", "F"},
					{"P", "_", "F", "_"},
					{"P", "P", "F", "P"},
					{"_", "P", "_", "P"},
				},
			},
			want: NEED_MORE_INFO,
		},
		{
			name: "Inconclusive",
			args: args{
				[][]string{
					{"P", "F", "F", "F"},
					{"P", "_", "F", "_"},
					{"P", "P", "F", "P"},
					{"F", "P", "P", "P"},
				},
			},
			want: INCONCLUSIVE,
		},
		{
			name: "Inconclusive",
			args: args{
				[][]string{
					{"P", "F", "F", "F"},
					{"P", "F", "F", "P"},
					{"F", "P", "F", "P"},
					{"P", "P", "P", "F"},
				},
			},
			want: INCONCLUSIVE,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectFraud(tt.args.grid); got != tt.want {
				t.Errorf("detectFraud() = %v, want %v", got, tt.want)
			}
		})
	}
}
