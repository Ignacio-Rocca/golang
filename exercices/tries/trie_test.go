package tries

import "testing"

func TestTrie_Search(t1 *testing.T) {
	tests := []struct {
		name   string
		trie   []string
		word   string
		want   bool
	}{
		{
			name: "find word in trie",
			trie: []string{"hola", "mundo", "rocca", "ignacio"},
			word: "hola",
			want: true,
		},
		{
			name: "don't find word in trie",
			trie: []string{"hola", "mundo", "rocca", "ignacio"},
			word: "invalid",
			want: false,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := InitTrie()
			t.Insert(tt.trie[0])
			t.Insert(tt.trie[1])
			t.Insert(tt.trie[2])
			t.Insert(tt.trie[3])

			if got := t.Search(tt.word); got != tt.want {
				t1.Errorf("Search() = %v, want %v", got, tt.want)
			}
		})
	}
}
