// Package reverse_sentence implements two solutions of
package reverse_sentence

import (
	"errors"
	"reflect"
	"strings"
)

// The ReverseSentence function returns the reverse sentence of some sentence received by param
func ReverseSentence(sentence string) string {
	sentenceArray := strings.Split(sentence, " ")

	if len(sentenceArray) == 1 {
		return sentence
	}

	sentenceArrayReversed := reverseStringSlice(sentenceArray)

	return strings.Join(sentenceArrayReversed, " ")
}

// The ReverseSentenceWithReflection function returns the reverse sentence of some sentence received by param
func ReverseSentenceWithReflection(sentence string) string {
	sentenceArray := strings.Split(sentence, " ")

	if len(sentenceArray) == 1 {
		return sentence
	}

	sentenceArrayReversed, _ := reverseSlice(sentenceArray)

	return strings.Join(sentenceArrayReversed.([]string), " ")
}

// The reverseStringSlice function receives an slice and return the slice reversed
func reverseStringSlice(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

// The reverseSlice function receives an slice of any type and return the array reversed
func reverseSlice(s interface{}) (interface{}, error) {
	if !isSlice(s) {
		return nil, errors.New("The type of 's' must be Slice")
	}

	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}

	return s, nil
}

// The isSlice function returns true when the received value from param is a slice, otherwise returns false
func isSlice(value interface{}) bool {
	typeOfValue := reflect.TypeOf(value).Kind()
	return typeOfValue == reflect.Slice || typeOfValue == reflect.Array
}