package main

import (
	"DSA/trie"
	"fmt"
)

// "DSA/linkedlist"
// "DSA/stack"
// "DSA/queue"
// "DSA/binarytree"
func main() {
	// list := &linkedlist.LinkedList{}
	// list.AddStart(10)
	// list.AddStart(20)
	// list.AddStart(30)
	// list.AddEnd(15)
	// list.AddEnd(16)
	// list.Remove(30)
	// list.Remove(16)
	// list.Remove(20)
	// list.AddStart(35)
	// list.AddEnd(25)
	// list.Print()
	// stack := &stack.Stack{}
	// stack.Push(12)
	// stack.Push(13)
	// stack.Push(14)
	// stack.Push(15)
	// stack.Print()
	// stack.Pop()
	// stack.Print()
	// stack.Pop()
	// stack.Print()

	// queue := &queue.Queue{}

	// queue.Add(10)
	// queue.Add(20)
	// queue.Add(30)
	// queue.Add(40)
	// queue.Print()
	// fmt.Println(queue.Remove())
	// fmt.Println(queue.Remove())
	// queue.Print()

	// tree := binarytree.BinaryTree{}

	// tree.Add(10)
	// tree.Add(20)
	// tree.Add(5)
	// tree.Add(25)
	// tree.Add(27)
	// tree.Add(17)
	// tree.Add(18)
	// tree.Add(2)
	// tree.Add(7)
	// tree.Remove(20)
	// tree.Print()

	// hashTable := &hashtable.HashTable{}
	// hashTable.Add(12)
	// hashTable.Add(17)
	// hashTable.Add(19)
	// hashTable.Add(21)
	// hashTable.Add(2)
	// hashTable.Add(1)
	// hashTable.Print()

	// gh := graph.AdjacencyGraph{}
	// for i := 0; i < 6; i++ {
	// 	gh.AddVertix(i)
	// }

	// gh.AddEdge(0, 1)
	// gh.AddEdge(2, 1)
	// gh.AddEdge(0, 2)
	// gh.AddEdge(0, 4)
	// gh.AddEdge(3, 4)
	// gh.AddEdge(2, 5)
	// gh.Print()

	// maxHeap := &heap.MaxHeap{}
	// for _, val := range []int{2, 12, 11, 10, 20, 50, 30, 70, 1, 7, 8} {
	// 	maxHeap.Insert(val)
	// }
	// fmt.Println(maxHeap)
	// maxHeap.Remove()
	// fmt.Println(maxHeap)
	// maxHeap.Remove()
	// fmt.Println(maxHeap)
	// maxHeap.Remove()
	// fmt.Println(maxHeap)
	// maxHeap.Remove()
	// fmt.Println(maxHeap)
	// maxHeap.Remove()
	// fmt.Println(maxHeap)

	// str:="ATEST"

	// for _, ch := range str {
	// 	fmt.Println(ch-'A') 
	// }


	t := &trie.Trie{}
	t.Init()
	t.Add("TEST")
	t.Add("WORD")
	t.Add("TESTER")
	t.Add("WORKER")

	fmt.Println(t.Search("TEST"))
	fmt.Println(t.Search("WORD"))
	fmt.Println(t.Search("WORKER"))
	fmt.Println(t.Search("TESTER"))
	fmt.Println(t.Search("TESTED"))
	fmt.Println(t.Search("KING"))
	fmt.Println(t.Search("ALPHA"))
	t.Add("ALPHA")
	fmt.Println(t.Search("ALPHA"))
}
