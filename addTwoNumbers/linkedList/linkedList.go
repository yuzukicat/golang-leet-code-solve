package linkedList

import "fmt"

type Node struct {
	prev *Node
	next *Node
	key  interface{}
}

type List struct {
	head *Node
	tail *Node
}

func (L *List) Insert(key interface{}) {
	list := &Node{
		next: L.head,
		key:  key,
	}
	if L.head != nil {
		L.head.prev = list
	}
	L.head = list

	l := L.head
	for l.next != nil {
		l = l.next
	}
	L.tail = l
}

func (l *List) Display() {
	list := l.head
	for list != nil {
		fmt.Printf("%+v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

func Display(list *Node) {
	for list != nil {
		fmt.Printf("%v ->", list.key)
		list = list.next
	}
	fmt.Println()
}

func AddTwoLinkedList(linkedListOne *List, linkedListTwo *List) {
	listOneNode := linkedListOne.tail
	listTwoNode := linkedListTwo.tail
	newLinkedList := List{}
	var carry *int
	zero := 0
	carry = &zero
	one := 1
	for listOneNode != nil && listTwoNode != nil {
		sum := listOneNode.key.(int) + listTwoNode.key.(int) + *carry
		if sum >= 10 {
			carry = &one
		} else {
			carry = &zero
		}
		newLinkedList.Insert(sum % 10)
		newLinkedList.Display()
		listOneNode = listOneNode.prev
		listTwoNode = listTwoNode.prev
	}
	if *carry == 1 {
		newLinkedList.Insert(*carry)
	}
	newLinkedList.Display()
}

func ShowBackwards(list *Node) {
	for list != nil {
		fmt.Printf("%v <-", list.key)
		list = list.prev
	}
	fmt.Println()
}

func (l *List) Reverse() {
	curr := l.head
	var prev *Node
	l.tail = l.head

	for curr != nil {
		next := curr.next
		curr.next = prev
		prev = curr
		curr = next
	}
	l.head = prev
	Display(l.head)
}
