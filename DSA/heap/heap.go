package heap

type MaxHeap struct {
	array []int
}

func (heap *MaxHeap) Insert(val int) {
	heap.array = append(heap.array, val)
	heap.maxHeapUp()
}

func (heap *MaxHeap) maxHeapUp() {
	idx := len(heap.array) - 1
	parent := heap.parent(idx)
	for heap.array[idx] > heap.array[parent] {
		heap.swap(idx, parent)
		idx = parent
		parent = heap.parent(idx)
	}
}

func (heap *MaxHeap) parent(idx int) int {
	return (idx - 1) / 2
}

func (heap *MaxHeap) left(idx int) int {
	return (2 * idx) + 1
}

func (heap *MaxHeap) right(idx int) int {
	return (2 * idx) + 2
}

func (heap *MaxHeap) swap(idx1, idx2 int) {
	heap.array[idx1], heap.array[idx2] = heap.array[idx2], heap.array[idx1]
}

func (heap *MaxHeap) Remove() int {
	val := heap.array[0]
	last := len(heap.array) - 1

	heap.array[0] = heap.array[last]
	heap.array = heap.array[:last]
	heap.minHeapDown(0)
	return val
}

func (heap *MaxHeap) minHeapDown(idx int) {
	left, right := heap.left(idx), heap.right(idx)
	last := len(heap.array) - 1
	child := 0
	for left <= last {
		if left == last {
			child = left
		} else if heap.array[left] > heap.array[right] {
			child = left
		} else {
			child = right
		}

		if heap.array[child] > heap.array[idx] {
			heap.swap(child, idx)
			idx = child
			left, right = heap.left(idx), heap.right(idx)
		} else {
			return
		}
	}
}
