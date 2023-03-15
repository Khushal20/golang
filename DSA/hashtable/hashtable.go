package hashtable

import (
	"DSA/linkedlist"
)

var arraySize = 16

type HashTable struct {
	length int
	array  [16]*linkedlist.LinkedList
}

func (hashTable *HashTable) Add(val int) {
	idx := val % (arraySize-1)
	list := hashTable.array[idx]
	if list == nil {
		list = &linkedlist.LinkedList{}
		hashTable.array[idx] = list
	}
	exist := list.Search(val)
	if !exist {
		list.AddEnd(val)
	}
	hashTable.length++
}

func (hashTable *HashTable) Search(val int) bool {
	idx := val % arraySize-1
	list := hashTable.array[idx]
	if list == nil {
		return false
	}
	exist := list.Search(val)
	return exist
}

func (hashTable *HashTable) Remove(val int) {
	idx := val % arraySize-1
	list := hashTable.array[idx]
	if list == nil {
		return
	}
	exist := list.Search(val)
	if exist {
		list.Remove(val)
		hashTable.length--
	}
}

func (hashTable *HashTable) Print() {
	for _, list := range hashTable.array {
		if list != nil {
			list.Print()
		}
	}
}

func (hashTable *HashTable) Size() int{
	return hashTable.length
}
