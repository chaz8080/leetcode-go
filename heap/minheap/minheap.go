package heap

import "log"

// Item Key(s) is the priority and should be a positive integer
type Item struct {
	Key   int
	Value interface{}
}

func (i *Item) IsNil() bool {
	return i.Key == 0
}

// max heap always has the max value at the root
type MinHeap struct {
	Items []Item
}

func NewHeap(items []Item) MinHeap {
	var h MinHeap
	for _, item := range items {
		h.Push(item.Key, item.Value)
	}

	return h
}

// check whether heap is empty
func (h *MinHeap) IsEmpty() bool {
	return len(h.Items) <= 0
}

// if heap is not empty, return the head
func (h *MinHeap) Peek() Item {
	if h.IsEmpty() {
		return Item{}
	} else {
		return h.Items[0]
	}
}

// push item into heap
func (h *MinHeap) Push(key int, value interface{}) {
	h.Items = append(h.Items, Item{
		Key:   key,
		Value: value,
	})

	h.heapifyUp(len(h.Items) - 1)
}

// pop root item off of heap
func (h *MinHeap) Pop() Item {
	// if there are no items, return nil item
	if h.IsEmpty() {
		log.Println("pop failed, heap is empty")
		return Item{}
	}

	// pop the first item
	item := h.Items[0]

	last := len(h.Items) - 1
	// move the last item to the root, and
	h.Items[0] = h.Items[last]

	// trim the last item
	h.Items = h.Items[:last]

	h.heapifyDown(0)

	return item
}

// bubble items up
func (h *MinHeap) heapifyUp(index int) {
	// loop until parent key is less than child
	for h.Items[parent(index)].Key > h.Items[index].Key {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// bubble items down
func (h *MinHeap) heapifyDown(index int) {
	lastIndex := len(h.Items) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	// loop until node has at least one child
	for l <= lastIndex {
		if l == lastIndex { // when the left child is the only child
			childToCompare = l
		} else if h.Items[l].Key < h.Items[r].Key { // when the left child is smaller
			childToCompare = l
		} else { // when the right child is smaller
			childToCompare = r
		}

		// compare array value of the index to smaller child and swap if larger
		if h.Items[index].Key > h.Items[childToCompare].Key {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

/*
	 in a heap, the parent = (child_index - 1) / 2
	 takes advantage of the fact that ints will round down, e.g. 3 and 4's parent will be 1...

				0
			   / \
			 1     2        (1-1) / 2 = 0
			/ \
	       3   4            (3-1) / 2 = 1
*/
func parent(index int) int {
	return (index - 1) / 2
}

/*
	 in a heap, the left_child_index = (parent * 2) + 1

				0
			   / \
			 1     2        (0 * 2) + 1 = 1
			/ \
	       3   4            (1 * 2) + 1 = 3
*/
func left(index int) int {
	return 2*index + 1
}

/*
	 in a heap, the right_child_index = (parent * 2) + 2

				0
			   / \
			 1     2        (0 * 2) + 2 = 2
			/ \
	       3   4            (1 * 2) + 2 = 4
*/
func right(index int) int {
	return 2*index + 2
}

// swap 2 elements in the heap
func (h *MinHeap) swap(index1 int, index2 int) {
	h.Items[index1], h.Items[index2] = h.Items[index2], h.Items[index1]
}
