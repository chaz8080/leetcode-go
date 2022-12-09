package meetingscheduler

import (
	"reflect"
	"testing"
)

func Test_minAvailableDuration(t *testing.T) {
	type args struct {
		slots1   [][]int
		slots2   [][]int
		duration int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example 1",
			args: args{
				slots1: [][]int{
					{10, 50},
					{60, 120},
					{140, 210},
				},
				slots2: [][]int{
					{0, 15},
					{60, 70},
				},
				duration: 8,
			},
			want: []int{60, 68},
		},
		{
			name: "example 2",
			args: args{
				slots1: [][]int{
					{10, 50},
					{60, 120},
					{140, 210},
				},
				slots2: [][]int{
					{0, 15},
					{60, 70},
				},
				duration: 12,
			},
			want: []int{},
		},
		{
			name: "example 3",
			args: args{
				slots1: [][]int{
					{0, 2},
				},
				slots2: [][]int{
					{1, 3},
				},
				duration: 1,
			},
			want: []int{1, 2},
		},
		{
			name: "example 4",
			args: args{
				slots1: [][]int{
					{10, 50},
					{60, 120},
					{140, 210},
				},
				slots2: [][]int{
					{0, 15},
					{40, 50},
				},
				duration: 8,
			},
			want: []int{40, 48},
		},
		{
			name: "example 5",
			args: args{
				slots1: [][]int{
					{216397070, 363167701},
					{98730764, 158208909},
					{441003187, 466254040},
					{558239978, 678368334},
					{683942980, 717766451},
				},
				slots2: [][]int{
					{50490609, 222653186},
					{512711631, 670791418},
					{730229023, 802410205},
					{812553104, 891266775},
					{230032010, 399152578},
				},
				duration: 456085,
			},
			want: []int{98730764, 99186849},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minAvailableDuration(tt.args.slots1, tt.args.slots2, tt.args.duration); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("minAvailableDuration() = %v, want %v", got, tt.want)
			}
		})
	}
}
