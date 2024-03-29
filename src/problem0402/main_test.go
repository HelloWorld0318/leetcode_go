package problem0402

import "testing"

func Test_removeKdigits(t *testing.T) {
	type args struct {
		num string
		k   int
	}

	arg := args{
		num: "1432219",
		k:3,
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: arg,
			want: "1219",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeKdigits(tt.args.num, tt.args.k); got != tt.want {
				t.Errorf("removeKdigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
