package pair_of_socks

import "testing"

func Test_contains(t *testing.T) {
	type args struct {
		data  []int
		value int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success 01",
			args: args{
				data: []int{1,2,3,4,5},
				value: 1,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := contains(tt.args.data, tt.args.value); got != tt.want {
				t.Errorf("contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetPairsOfSocks(t *testing.T) {
	type args struct {
		data  []int
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "success 01",
			args: args{
				data: []int{1,1,2,2,3,4,5,2,3},
				n: 9,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPairOfSocksWithoutSort(tt.args.n, tt.args.data); got != tt.want {
				t.Errorf("getPairOfSocksWithoutSort() = %v, want %v", got, tt.want)
			}
			if got := getPairOfSocksWithSort(tt.args.n, tt.args.data); got != tt.want {
				t.Errorf("getPairOfSocksWithSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
