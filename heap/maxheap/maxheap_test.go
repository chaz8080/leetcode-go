package heap

import (
	"reflect"
	"testing"
)

func TestNewHeap(t *testing.T) {
	type args struct {
		items []Item
	}
	tests := []struct {
		name string
		args args
		want MaxHeap
	}{
		{
			name: "test new heap",
			args: args{
				items: []Item{
					{Key: 1, Value: nil},
					{Key: 2, Value: nil},
				},
			},
			want: NewHeap([]Item{
				{Key: 2, Value: nil},
				{Key: 1, Value: nil},
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHeap(tt.args.items); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHeap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeapPop(t *testing.T) {
	type args struct {
		items []Item
	}
	tests := []struct {
		name string
		args args
		want MaxHeap
	}{
		{
			name: "test pop",
			args: args{
				items: []Item{
					{Key: 1, Value: nil},
					{Key: 2, Value: nil},
				},
			},
			want: NewHeap([]Item{
				{Key: 1, Value: nil},
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHeap(tt.args.items)
			h.Pop()
			if got := h; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHeap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeapPush(t *testing.T) {
	type args struct {
		items []Item
	}
	tests := []struct {
		name string
		args args
		want MaxHeap
	}{
		{
			name: "test pop",
			args: args{
				items: []Item{
					{Key: 1, Value: nil},
					{Key: 2, Value: nil},
				},
			},
			want: NewHeap([]Item{
				{Key: 3, Value: nil},
				{Key: 1, Value: nil},
				{Key: 2, Value: nil},
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHeap(tt.args.items)
			h.Push(3, nil)
			if got := h; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHeap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHeapPeek(t *testing.T) {
	type args struct {
		items []Item
	}
	tests := []struct {
		name string
		args args
		want Item
	}{
		{
			name: "test peek",
			args: args{
				items: []Item{
					{Key: 1, Value: nil},
					{Key: 2, Value: nil},
					{Key: 3, Value: nil},
				},
			},
			want: Item{Key: 3, Value: nil},
		},
		{
			name: "test peek empty",
			args: args{
				items: []Item{},
			},
			want: Item{Key: 0, Value: nil},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHeap(tt.args.items)
			item := h.Peek()
			if got := item; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewHeap() = %v, want %v", got, tt.want)
			}
		})
	}
}
