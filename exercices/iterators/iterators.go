package main

import "fmt"

// Iterate calls the f function with n = 1, 2, and 3.
func Iterate01(f func(n int)) {
	for i := 1; i <= 3; i++ {
		f(i)
	}
}

// Iterate calls the f function with n = 1, 2, and 3.
// If f returns true, Iterate returns immediately
// skipping any remaining values.
func Iterate02(f func(n int) (skip bool)) {
	for i := 1; i <= 3; i++ {
		if f(i) {
			return
		}
	}
}

func main() {
	Iterate01(func(n int) { fmt.Println(n) })
	Iterate02(func(n int) (skip bool) {
		fmt.Println(n)
		return n == 2
	})
}