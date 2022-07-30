package problem0206

import (
	"fmt"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
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
		head:
		head,
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
			got := reverseList(tt.args.head)
			for got != nil {
				fmt.Println(got.Val)
				got = got.Next
			}
			str := "abba"
			fmt.Println(string(str[0]))
		})
	}
}
