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

func main() {
	ll := &LinkedList{}

	for i := 0; i < 10; i++ {
		ll.push(i * 10)
	}

	ll.output()

	ll.insert(999, 0)

	ll.output()

	ll.find(999)
	fmt.Println("40 at", ll.find(40))
	ll.find(90)
	ll.find(123)
}
