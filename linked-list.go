package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func push(head *Node, v int) {
	newNode := &Node{value: v}
	if head.next == nil {
		head.next = newNode
	} else {
		current := head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
}

func showNodes(head *Node) {
	current := head
	i := 0
	for current.next != nil {
		fmt.Println(i, ":", current.value)
		current = current.next
		i++
	}
	fmt.Println(i, ":", current.value)
}

func insert(head *Node, v, p int) {
	current := head
	i := 0
	for current.next != nil {
		if i == p-1 {
			// Found the position of the previous node.
			// Change this one's next and if it had a next, use that in the new node's next.
			// Also break out of the loop. No point in carrying on.
			oldNext := current.next
			newNode := &Node{value: v, next: oldNext}
			current.next = newNode
			break
		}
		current = current.next
		i++
	}

}

func main() {
	head := Node{
		value: 10,
	}

	for i := 1; i < 10; i++ {
		push(&head, i+10)
	}

	insert(&head, 20, 2)
	showNodes(&head)
}
