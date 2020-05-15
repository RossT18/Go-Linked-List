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
		for current != nil {
			fmt.Println(i, ":", current.value)
			current = current.next
			i++
		}
		fmt.Println("")
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
				return i
			}
			current = current.next
			i++
		}
		if current.value == v {
			// Have to check the last element too
			return i
		} else {
			// value is not in the list
			return -1
		}
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

func (ll *LinkedList) remove(v int) int {
	searchResult := ll.find(v)
	if searchResult > -1 {
		// This means it exists in the linked list
		return ll.removeAt(searchResult)
	} else {
		fmt.Println("ERROR: Value", v, "cannot be removed. It is not present")
		return -1
	}
}

func (ll *LinkedList) reverse() {
	llReversed := &LinkedList{}
	originalLength := ll.length
	for i := 0; i < originalLength; i++ {
		llReversed.push(ll.pop())
	}

	current := llReversed.head
	for current != nil {
		ll.push(current.value)
		current = current.next
	}
}

func (ll *LinkedList) clear() {
	ll.head = nil
	ll.length = 0
}

func (ll *LinkedList) sortAdd(v int) {
	if ll.length == 0 {
		// LinkedList is empty, so can just set this value to the head of the list.
		ll.head = &Node{value: v, next: nil}
	} else {
		var previous *Node = nil
		current := ll.head
		for current != nil {
			if current.value < v {
				// The current val in the ll is smaller, so should look at next one
				if current.next == nil {
					// There is no next one, so add to end
					newNode := &Node{value: v, next: nil}
					current.next = newNode
					break
				} else {
					// More list to go through, keep checking.
					previous = current
					current = current.next
				}
			} else {
				// Current val is bigger or same as value. Should insert just before this node
				if previous == nil {
					// Must mean the head has a bigger val than what needs to be inserted
					newNode := &Node{value: v, next: ll.head}
					ll.head = newNode
				} else {
					newNode := &Node{value: v, next: current}
					previous.next = newNode
				}
				break //Exit the loop. We have inserted
			}
		}
	}
	ll.length++
}

func (ll *LinkedList) sort() {
	sorted := &LinkedList{}

	current := ll.head
	for current != nil {
		sorted.sortAdd(current.value)
		current = current.next
	}

	ll.clear()

	newCurrent := sorted.head
	for newCurrent != nil {
		ll.push(newCurrent.value)
		newCurrent = newCurrent.next
	}
}

func main() {
	ll := &LinkedList{}

	for i := 0; i < 10; i++ {
		ll.push(i * 10)
	}

	ll.push(-1)
	ll.push(400)
	ll.push(300)
	ll.push(45)
	ll.push(30)

	ll.output()

	ll.sort()

	ll.output()
}
