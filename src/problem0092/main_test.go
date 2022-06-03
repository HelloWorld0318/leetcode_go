package problem0092

import (
	"fmt"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}

	head := &ListNode{
		Val:  1,
		Next: nil,
	}
	second := &ListNode{
		Val:  2,
		Next: nil,
	}
	head.Next = second
	third := &ListNode{
		Val:  3,
		Next: nil,
	}
	second.Next = third
	forth := &ListNode{
		Val:  4,
		Next: nil,
	}
	third.Next = forth
	fifth := &ListNode{
		Val:  5,
		Next: nil,
	}
	forth.Next = fifth

	arg := args{
		head:  head,
		left:  2,
		right: 4,
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{"",
			arg,
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := reverseBetween(tt.args.head, tt.args.left, tt.args.right)
			for got != nil {
				fmt.Println(got.Val)
				got = got.Next
			}
		})
	}
}
