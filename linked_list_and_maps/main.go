// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Node struct {
	name string
	next *Node
}

type List struct {
	head *Node
}

func (list *List) addItem(name string) {
	newNode := &Node{name: name}

	if list.head == nil {
		list.head = newNode
		return
	}

	current := list.head

	if current.next != nil {
		current = current.next
	}

	current.next = newNode
}

func main() {
	linkedList := &List{}

	linkedList.addItem("item1")
	linkedList.addItem("item2")
	linkedList.addItem("item3")

	current := linkedList.head

	for current != nil {
		fmt.Println(current.name)

		current = current.next
	}

	var score = make(map[string]int)

	score["John"] = 10

	score["Jane"] = 20

	score["Joe"] = 30

	fmt.Println(score)

	delete(score, "John")

	fmt.Println(score)
}
