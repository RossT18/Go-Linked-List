package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

func addNode(head *Node, v int) {
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
	i := 1
	for current.next != nil {
		fmt.Println(i, ":", current.value)
		current = current.next
		i++
	}
	fmt.Println(i, ":", current.value)
}

func main() {

	head := Node{
		value: 10,
	}

	for i := 1; i < 10; i++ {
		addNode(&head, i+10)
	}

	showNodes(&head)
}
