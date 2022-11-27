package main

import (
	"fmt"
	maxheap "heap/maxheap"
)

func main() {
	m := &maxheap.MaxHeap{}
	for range []int{1, 20, 3, 4, 5} {
		m.Pop()
		fmt.Println(m)
	}

	for _, val := range []int{1, 20, 3, 4, 5, 6} {
		m.Push(val, nil)
		fmt.Println(m)
	}

	for range []int{1, 20, 3, 4, 5} {
		m.Pop()
		fmt.Println(m)
	}
}
