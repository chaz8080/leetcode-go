package oddevenlinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	start := head
	end := head

	// get the last node
	for end.Next != nil {
		end = end.Next
	}

	// need to keep a reference to the last node for exit condition
	// continue adding to the tempEnd all of the even numbers
	tempEnd := end

	// reset head
	head = start

	// iterate through, putting the odds at the end
	for head != end {
		// shift next odd left

		// if it was an even number of nodes, the last needs to be moved to the end
		if head.Next == end {
			tempEnd.Next = end
			tempEnd = tempEnd.Next
			head.Next = head.Next.Next
			break
		}

		// move next (even) to end
		tempEnd.Next = head.Next
		tempEnd = tempEnd.Next

		head.Next = head.Next.Next
		head = head.Next
	}

	tempEnd.Next = nil

	return start
}
