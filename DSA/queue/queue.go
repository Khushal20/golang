package queue

import "fmt"

type Queue struct {
	array  []int
	length int
}

func (queue *Queue) Add(val int) {
	if queue.array == nil {
		queue.array = []int{}
	}
	queue.array = append(queue.array, val)
	queue.length++
}

func (queue *Queue) Remove() int {
	if queue.array == nil {
		return -1
	}
	if queue.length == 0 {
		return -1
	}

	tmp := queue.array[0]
	queue.array = queue.array[1:]
	queue.length--
	return tmp
}

func (queue *Queue) Top() int{
	if queue.array == nil {
		return -1
	}
	if queue.length == 0 {
		return -1
	}
	tmp := queue.array[0]
	return tmp
}

func (queue *Queue) Print() {
	fmt.Println(queue.array)
}

func (queue *Queue) Size() int {
	return queue.length
}
