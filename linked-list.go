package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	length int
	head   *Node
}

func (ll *LinkedList) push(v int) {
	newNode := &Node{value: v}

	if ll.length == 0 {
		ll.head = newNode
	} else {
		current := ll.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	ll.length++
}

func (ll *LinkedList) pop() int {
	if ll.length == 0 {
		fmt.Println("ERROR: Cannot pop off of an empty list.") //Actually handle this
		return -1
	} else if ll.length == 1 {
		headVal := ll.head.value
		ll.head = nil
		ll.length--
		return headVal
	} else {
		var previous *Node = nil
		current := ll.head
		for current.next != nil {
			previous = current
			current = current.next
		}
		val := current.value
		previous.next = nil
		ll.length--
		return val
	}
}

func (ll *LinkedList) output() {
	if ll.length == 0 {
		fmt.Println("LinkedList is empty!")
	} else {
		fmt.Println("Linked List. Length:", ll.length)
		current := ll.head
		i := 0
		for current.next != nil {
			fmt.Println(i, ":", current.value)
			current = current.next
			i++
		}
		fmt.Println(i, ":", current.value)
		fmt.Println("")
	}
}

func (ll *LinkedList) insert(v, p int) {
	if ll.length == p {
		// The position to insert into is one after the end of the list, so it's a push
		ll.push(v)
		ll.length++
	} else if p > ll.length {
		// The position is outside of it's range. If this worked, it would create gaps.
		// So don't allow this to happen
		fmt.Println("ERROR: Cannot insert outside of linked list.")
		// TODO: Should handle errors by returning an error instead of this.
	} else {
		if p == 0 {
			// Insert before rest of list.
			oldHead := ll.head
			ll.head = &Node{value: v, next: oldHead}
		} else {
			current := ll.head
			i := 0
			for current.next != nil {
				if i == p-1 {
					// Found the position of the previous node.
					// Change this one's next and if it had a next, use that in the new node's next.
					// Also break out of the loop. No point in carrying on.
					oldNext := current.next
					current.next = &Node{value: v, next: oldNext}
					break
				}
				current = current.next
				i++
			}
		}
		ll.length++
	}
}

func (ll *LinkedList) removeAt(p int) int {
	// TODO: Make switch-case
	if ll.length == 0 {

		fmt.Println("ERROR: Cannot remove from empty list")
		return -1

	} else if p >= ll.length || p < 0 {

		// This is out of range of the list.
		// If the length is 10, the maximum P is 9, as this is the 10th element
		fmt.Println("ERROR: Cannot remove from outside the list")
		return -1

	} else if p == ll.length-1 {

		// Last item was selected
		return ll.pop()

	} else if p == 0 {

		// Need to remove head and make second element, if any, the head
		// No need to check if head has any nodes. It must, or else the condition above (p == ll.length-1) will be true
		headVal := ll.head.value
		ll.head = ll.head.next
		ll.length--
		return headVal

	} else {
		// Need to remove an item from the middle.
		var removedVal int
		var previous *Node = nil
		current := ll.head
		i := 0
		for current.next != nil {
			if i == p {
				// Current is now the node to be removed
				removedVal = current.value
				previous.next = current.next
			}
			previous = current
			current = current.next
			i++
		}
		// It won't be the last value, as that's already checked for and popped above
		ll.length--
		return removedVal

	}
}

func (ll *LinkedList) find(v int) int {
	if ll.length == 0 {
		fmt.Println("ERROR: LinkedList is empty")
		return -1
	} else {
		current := ll.head
		i := 0
		for current.next != nil {
			if current.value == v {
				fmt.Println(v, "found at", i)
				return i
			}
			current = current.next
			i++
		}
		if current.value == v {
			// Have to check the last element too
			fmt.Println(v, "found at", i)
			return i
		} else {
			fmt.Println("LinkedList does not contain", v)
			return -1
		}
	}
}

func (ll *LinkedList) reverse() *LinkedList {
	llReversed := &LinkedList{}
	originalLength := ll.length
	for i := 0; i < originalLength; i++ {
		llReversed.push(ll.pop())
	}
	return llReversed
}

func main() {
	ll := &LinkedList{}

	for i := 0; i < 10; i++ {
		ll.push(i * 10)
	}

	ll.output()

	ll.removeAt(0)
	ll.removeAt(1)
	ll.removeAt(2)
	ll.removeAt(3)
	ll.removeAt(4)

	ll.output()
}
