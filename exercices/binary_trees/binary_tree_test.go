package binary_trees

import (
	"testing"
)

func TestLookupOnBST(t *testing.T) {
	type args struct {
		root   *Node
		target int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success with a valid bst",
			args: args{
				root: &Node{
					Val: 4,
					Left: &Node{
						Val: 2,
						Left: &Node{Val: 1},
						Right: &Node{Val: 3},
					},
					Right: &Node{
						Val: 6,
						Left: &Node{Val:5},
						Right: &Node{Val:7},
					},
				},
				target: 7,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.args.root.LookupOnBST(tt.args.target); got != tt.want {
				t.Errorf("LookupOnBST() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestInsertOnBST(t *testing.T) {
	type args struct {
		root   *Node
		val int
	}
	tests := []struct {
		name string
		args args
		rootNodeVal int
	}{
		{
			name: "success with a valid bst",
			args: args{
				root: &Node{
					Val: 4,
					Left: &Node{
						Val: 2,
						Left: &Node{Val: 1},
						Right: &Node{Val: 3},
					},
					Right: &Node{
						Val: 6,
						Left: &Node{Val:5},
						Right: &Node{Val:7},
					},
				},
				val: 8,
			},
			rootNodeVal: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.root.InsertOnBST(tt.args.val)
			if got.Val != tt.rootNodeVal {
				t.Errorf("InsertOnBST() = %v, want %v", got.Val, tt.rootNodeVal)
			}
		})
	}
}

func TestSize(t *testing.T) {

	tests := []struct {
		name string
		args *Node
		size int
	}{
		{
			name: "success with a valid bst",
			args: &Node{
				Val: 4,
				Left: &Node{
					Val: 2,
					Left: &Node{Val: 1},
					Right: &Node{Val: 3},
				},
				Right: &Node{
					Val: 6,
					Left: &Node{Val:5},
					Right: &Node{Val:7},
				},
			},
			size: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.args.Size()
			if got != tt.size {
				t.Errorf("Size() = %v, want %v", got, tt.size)
			}
		})
	}
}
