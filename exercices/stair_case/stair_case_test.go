package stair_case

import "testing"

func TestNumWays(t *testing.T) {

	tests := []struct {
		name string
		n int
		want int
	}{
		{
			name: "01",
			n: 3,
			want: 3,
		},
		{
			name: "02",
			n: 2,
			want: 2,
		},
		{
			name: "03",
			n: 1,
			want: 1,
		},
		{
			name: "04",
			n: 0,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumWays(tt.n); got != tt.want {
				t.Errorf("NumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumWaysDP_BottomUp(t *testing.T) {

	tests := []struct {
		name string
		n int
		want int
	}{
		{
			name: "01",
			n: 3,
			want: 3,
		},
		{
			name: "02",
			n: 2,
			want: 2,
		},
		{
			name: "03",
			n: 1,
			want: 1,
		},
		{
			name: "04",
			n: 0,
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumWaysDP_BottomUp(tt.n); got != tt.want {
				t.Errorf("NumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNumWaysComplex(t *testing.T) {

	tests := []struct {
		name string
		n int
		steps []int
		want int
	}{
		{
			name: "01",
			n: 3,
			steps: []int{1,3,5},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NumWaysComplex(tt.n, tt.steps); got != tt.want {
				t.Errorf("NumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
