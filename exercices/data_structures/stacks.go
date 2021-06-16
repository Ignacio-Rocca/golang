package data_structures

import "fmt"

// A basic stack (LIFO) data structure
func stackExample() {

	var stack []string

	stack = append(stack, "world!") // Push
	stack = append(stack, "Hello ")

	for len(stack) > 0 {
		n := len(stack) - 1 // Top element
		fmt.Print(stack[n])

		stack[n] = ""     // To prevent memory leaks
		stack = stack[:n] // Pop
	}
}
