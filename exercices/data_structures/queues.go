package data_structures

import (
	"container/list"
	"fmt"
)

// FIFO QUEUE WITH SLICE
func exampleQueueFifoWithSlice() {
	var queueSlice []string

	queueSlice = append(queueSlice, "Hello ")
	queueSlice = append(queueSlice, "world !")

	for len(queueSlice) > 0 {
		fmt.Print(queueSlice[0]) 		// First element
		queueSlice = queueSlice[1:]   	// Dequeue
	}
}



//FIFO QUEUE WITH LINKED LIST
func exampleQueueFifoWithLinkedList() {
	queue := list.New()

	queue.PushBack("Hello ") // Enqueue
	queue.PushBack("world!")

	for queue.Len() > 0 {
		e := queue.Front() // First element
		fmt.Print(e.Value)

		queue.Remove(e) // Dequeue
	}
}

