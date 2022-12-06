package oddevenlinkedlist

import (
	"testing"

	"gotest.tools/assert"
)

func Test_oddEvenList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "example 1",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  5,
									Next: nil,
								},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val:  4,
								Next: nil,
							},
						},
					},
				},
			},
		},
		{
			name: "example 1.1",
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val: 5,
									Next: &ListNode{
										Val:  6,
										Next: nil,
									},
								},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 5,
						Next: &ListNode{
							Val: 2,
							Next: &ListNode{
								Val: 4,
								Next: &ListNode{
									Val:  6,
									Next: nil,
								},
							},
						},
					},
				},
			},
		},
		{
			name: "example 2",
			args: args{
				head: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 1,
						Next: &ListNode{
							Val: 3,
							Next: &ListNode{
								Val: 5,
								Next: &ListNode{
									Val: 6,
									Next: &ListNode{
										Val: 4,
										Next: &ListNode{
											Val:  7,
											Next: nil,
										},
									},
								},
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 7,
							Next: &ListNode{
								Val: 1,
								Next: &ListNode{
									Val: 5,
									Next: &ListNode{
										Val:  4,
										Next: nil,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := oddEvenList(tt.args.head)
			for got != nil {
				assert.Equal(t, got.Val, tt.want.Val)
				got = got.Next
				tt.want = tt.want.Next
			}
		})
	}
}
