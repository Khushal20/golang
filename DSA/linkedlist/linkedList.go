package linkedlist

import "fmt"

type node struct {
	val  int
	next *node
}

type LinkedList struct {
	head   *node
	last   *node
	length int
}

func (list *LinkedList) AddStart(val int) {
	if list.length == 0 {
		list.head = &node{val: val}
		list.last = list.head
		list.length++
		return
	}

	tmp := &node{val: val, next: list.head}
	list.head = tmp
	list.length++
}

func (list *LinkedList) AddEnd(val int) {
	if list.length == 0 {
		list.head = &node{val: val}
		list.last = list.head
		list.length++
		return
	}

	list.last.next = &node{val: val}
	list.last = list.last.next
	list.length++
}

func (list *LinkedList) Remove(val int) {
	if list.length == 0 {
		return
	}

	pre := list.head
	if pre.val == val {
		list.head = pre.next
		list.length--
		if list.length <= 1 {
			list.last = list.head
		}
		return
	}

	for pre.next.val != val {
		pre = pre.next
	}
	tmp := pre.next
	pre.next = pre.next.next
	if tmp == list.last {
		list.last = pre
	}

	list.length--
}

func (list *LinkedList) Print() {
	pre := list.head
	for pre != nil {
		fmt.Print(pre.val, ", ")
		pre = pre.next
	}
}

func (list *LinkedList) Search(val int) bool{
	pre := list.head
	for pre!=nil{
		if pre.val == val{
			return true
		}
		pre = pre.next
	}
	return false
}


func (list *LinkedList) Size() int{
	return list.length
}
