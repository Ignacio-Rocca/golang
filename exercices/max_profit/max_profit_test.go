package max_profit

import "testing"

func Test_maxProfit(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{
			name: "success",
			args: []int{7,1,5,3,6,4},
			want: 7,
		},
		{
			name: "success",
			args: []int{5,4,3,2,1},
			want: 0,
		},
		{
			name: "success",
			args: []int{1,2,3,4,5},
			want: 4,
		},
		{
			name: "success",
			args: []int{},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfit(tt.args); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
