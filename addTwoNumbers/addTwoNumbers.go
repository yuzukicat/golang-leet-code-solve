package main

import (
	"github.com/yuzukicat/golang-leet-code-solve/addTwoNumbers/linkedList"
)

func main() {
	linkedListOne := linkedList.List{}
	linkedListOne.Insert(5)
	linkedListOne.Insert(2)
	linkedListOne.Insert(1)
	linkedListTwo := linkedList.List{}
	linkedListTwo.Insert(2)
	linkedListTwo.Insert(7)
	linkedListTwo.Insert(3)
	linkedListOne.Display()
	linkedListTwo.Display()
	linkedList.AddTwoLinkedList(&linkedListOne, &linkedListTwo)
}
