package problem0376

import "testing"

func Test_wiggleMaxLength(t *testing.T) {
	type args struct {
		nums []int
	}
	arg := args{
		nums: []int{1,7,4,9,2,5},
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{"",
			arg,
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wiggleMaxLength(tt.args.nums); got != tt.want {
				t.Errorf("wiggleMaxLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
