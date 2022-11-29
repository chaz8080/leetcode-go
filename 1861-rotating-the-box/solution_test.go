package rotatingthebox

import (
	"reflect"
	"testing"
)

func Test_rotateTheBox(t *testing.T) {
	type args struct {
		box [][]byte
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		{
			name: "example 1",
			args: args{
				box: [][]byte{
					[]byte("#.#"),
				},
			},
			want: [][]byte{
				[]byte("."),
				[]byte("#"),
				[]byte("#"),
			},
		},
		{
			name: "example 2",
			args: args{
				box: [][]byte{
					[]byte("#.*."),
					[]byte("##*."),
				},
			},
			want: [][]byte{
				[]byte("#."),
				[]byte("##"),
				[]byte("**"),
				[]byte(".."),
			},
		},
		{
			name: "example 3",
			args: args{
				box: [][]byte{
					[]byte("##*.*."),
					[]byte("###*.."),
					[]byte("###.#."),
				},
			},
			want: [][]byte{
				[]byte(".##"),
				[]byte(".##"),
				[]byte("##*"),
				[]byte("#*."),
				[]byte("#.*"),
				[]byte("#.."),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rotateTheBox(tt.args.box); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rotateTheBox() = %v, want %v", got, tt.want)
			}
		})
	}
}
