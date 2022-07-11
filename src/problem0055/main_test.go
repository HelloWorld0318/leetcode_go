package problem0055

import "testing"

func Test_canJump(t *testing.T) {
	type args struct {
		nums []int
	}

	arg := args{
		nums: []int{2, 3, 1, 1, 4},
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: arg,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
