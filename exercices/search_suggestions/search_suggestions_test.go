package search_suggestions

import (
	"reflect"
	"testing"
)

func Test_searchSuggestions(t *testing.T) {
	type args struct {
		repository    []string
		customerQuery string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchSuggestions(tt.args.repository, tt.args.customerQuery); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchSuggestions() = %v, want %v", got, tt.want)
			}
		})
	}
}