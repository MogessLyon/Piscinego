package main

import (
	"fmt"
)

/*
Instructions

You have two numbers represented by a linked list, where each NodeAddL contains a single digit. The digits are stored in reverse order, such that the 1’s digit is at the head of the list. Write a function that adds the two numbers and returns the sum as a linked list
*/

type NodeAddL struct {
	Next *NodeAddL
	Num  int
}

func Revertlinklist(node *NodeAddL) {
	current := node
	var prev *NodeAddL
	prev = nil

	for current != nil {
		next := current.Next
		current.Next = prev

		prev = current
		current = next

	}
}

/*
func AddLinkedNumbers(num1, num2 *NodeAddL) *NodeAddL {

}
*/

func pushFront(node *NodeAddL, num int) *NodeAddL {
	data := &NodeAddL{Num: num}
	if node == nil {
		node = data
	} else {
		data.Next = node
		node = data
	}

	return node
}

func main() {
	// 3 -> 1 -> 5
	num1 := &NodeAddL{Num: 5}
	num1 = pushFront(num1, 1)
	num1 = pushFront(num1, 3)

	// 5 -> 9 -> 2
	num2 := &NodeAddL{Num: 2}
	num2 = pushFront(num2, 9)
	num2 = pushFront(num2, 5)

	Revertlinklist(num1)

	// 9 -> 0 -> 7
	// result := AddLinkedNumbers(num1, num2)
	for tmp := num1; tmp != nil; tmp = tmp.Next {
		fmt.Print(tmp.Num)
		if tmp.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}
