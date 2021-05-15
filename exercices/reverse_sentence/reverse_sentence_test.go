package reverse_sentence

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ReverseSentence(t *testing.T) {
	tests := []struct {
		name             string
		sentence         string
		sentenceExpected string
	}{
		{
			name:             "success scenario 01",
			sentence:         "Hello world",
			sentenceExpected: "world Hello",
		},
		{
			name:             "success scenario 02",
			sentence:         "Hello world !",
			sentenceExpected: "! world Hello",
		},
		{
			name:             "success scenario 03",
			sentence:         "today is the first day of the rest of my life",
			sentenceExpected: "life my of rest the of day first the is today",
		},
		{
			name:             "success scenario 04",
			sentence:         " ",
			sentenceExpected: " ",
		},
		{
			name:             "success scenario 05",
			sentence:         ".",
			sentenceExpected: ".",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sentenceReversed := ReverseSentence(tt.sentence)
			assert.Equal(t, tt.sentenceExpected, sentenceReversed)
		})
	}
}

func Test_ReverseStringSlice(t *testing.T) {
	tests := []struct {
		name                  string
		sentenceArray         []string
		sentenceArrayExpected []string
	}{
		{
			name:                  "success scenario 01",
			sentenceArray:         []string{"Hello", "world"},
			sentenceArrayExpected: []string{"world", "Hello"},
		},
		{
			name:                  "success scenario 02",
			sentenceArray:         []string{"Hello", "world", "!"},
			sentenceArrayExpected: []string{"!", "world", "Hello"},
		},
		{
			name:                  "success scenario 03",
			sentenceArray:         []string{"today", "is", "the", "first", "day", "of", "the", "rest", "of", "my", "life"},
			sentenceArrayExpected: []string{"life", "my", "of", "rest", "the", "of", "day", "first", "the", "is", "today"},
		},
		{
			name:                  "success scenario 04",
			sentenceArray:         []string{"", ""},
			sentenceArrayExpected: []string{"", ""},
		},
		{
			name:                  "success scenario 05",
			sentenceArray:         []string{"."},
			sentenceArrayExpected: []string{"."},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sentenceReversed := reverseStringSlice(tt.sentenceArray)
			assert.Equal(t, tt.sentenceArrayExpected, sentenceReversed)
		})
	}
}

func Test_ReverseSentenceWithReflection(t *testing.T) {
	tests := []struct {
		name             string
		sentence         string
		sentenceExpected string
	}{
		{
			name:             "success scenario 01",
			sentence:         "Hello world",
			sentenceExpected: "world Hello",
		},
		{
			name:             "success scenario 02",
			sentence:         "Hello world !",
			sentenceExpected: "! world Hello",
		},
		{
			name:             "success scenario 03",
			sentence:         "today is the first day of the rest of my life",
			sentenceExpected: "life my of rest the of day first the is today",
		},
		{
			name:             "success scenario 04",
			sentence:         " ",
			sentenceExpected: " ",
		},
		{
			name:             "success scenario 05",
			sentence:         ".",
			sentenceExpected: ".",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sentenceReversed := ReverseSentenceWithReflection(tt.sentence)
			assert.Equal(t, tt.sentenceExpected, sentenceReversed)
		})
	}
}

func Test_ReverseSlice(t *testing.T) {
	tests := []struct {
		name                  string
		sentenceArray         interface{}
		sentenceArrayExpected interface{}
		errorExpected         error
	}{
		{
			name:                  "success scenario 01",
			sentenceArray:         []string{"Hello", "world"},
			sentenceArrayExpected: []string{"world", "Hello"},
		},
		{
			name:                  "success scenario 02",
			sentenceArray:         []string{"Hello", "world", "!"},
			sentenceArrayExpected: []string{"!", "world", "Hello"},
		},
		{
			name:                  "success scenario 03",
			sentenceArray:         []string{"today", "is", "the", "first", "day", "of", "the", "rest", "of", "my", "life"},
			sentenceArrayExpected: []string{"life", "my", "of", "rest", "the", "of", "day", "first", "the", "is", "today"},
		},
		{
			name:                  "success scenario 04",
			sentenceArray:         []string{"", ""},
			sentenceArrayExpected: []string{"", ""},
		},
		{
			name:                  "success scenario 05",
			sentenceArray:         []string{"."},
			sentenceArrayExpected: []string{"."},
		},
		{
			name:                  "success scenario 05",
			sentenceArray:         []int{1, 2, 3, 4, 5},
			sentenceArrayExpected: []int{5, 4, 3, 2, 1},
		},
		{
			name:          "error invalid type",
			sentenceArray: 1,
			errorExpected: errors.New("The type of 's' must be Slice"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sentenceReversed, err := reverseSlice(tt.sentenceArray)

			if tt.errorExpected != nil {
				assert.Error(t, err)
				assert.Equal(t, tt.errorExpected.Error(), err.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.sentenceArrayExpected, sentenceReversed)
			}
		})
	}
}

func Test_IsSlice(t *testing.T) {
	tests := []struct {
		name           string
		param          interface{}
		resultExpected bool
	}{
		{
			name:           "success 02",
			param:          []string{"Hello", "world"},
			resultExpected: true,
		},
		{
			name:           "success 02",
			param:          []int{1, 2, 3, 4},
			resultExpected: true,
		},
		{
			name:           "invalid type 01",
			param:          "string",
			resultExpected: false,
		},
		{
			name:           "invalid type 02",
			param:          1,
			resultExpected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			isSlice := isSlice(tt.param)
			assert.Equal(t, tt.resultExpected, isSlice)
		})
	}
}
