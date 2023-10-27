package chapter_searching

import "testing"

func Test_maxSales(t *testing.T) {
	type args struct {
		sales []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "t1", args: args{
			sales: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		}, want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSales(tt.args.sales); got != tt.want {
				t.Errorf("maxSales() = %v, want %v", got, tt.want)
			}
		})
	}
}
