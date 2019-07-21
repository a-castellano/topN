package top

import topheap "github.com/a-castellano/topN/heap"
import "container/heap"
import "errors"

type TopN struct {
	Heap *topheap.MaxHeap
	Size int
}

func NewTop(size int) (*TopN, error) {

	if size < 0 {
		return nil, errors.New("Size value must be greater than 0.")
	}

	top := new(TopN)
	top.Size = size
	top.Heap = &topheap.MaxHeap{}
	heap.Init(top.Heap)

	return top, nil
}

func (t *TopN) Push(item int) {
	heap.Push(t.Heap, item)
	if t.Heap.Len() > t.Size {
		elements := t.PopElements()
		for i := 0; i < t.Size; i++ {
			heap.Push(t.Heap, elements[i])
		}
	}
}

func (t *TopN) ShowMax() (int, error) {
	if (*t.Heap).Len() < 1 {
		return 0, errors.New("Size value must be greater than 0.")
	}
	return (*t.Heap)[0], nil
}

func (t *TopN) PopElements() []int {
	var elements []int
	var elementsToPop int
	var heapLen int = (*t.Heap).Len()

	if heapLen == 0 {
		return elements
	} else {
		if heapLen < t.Size {
			elementsToPop = heapLen
		} else {
			elementsToPop = t.Size
		}
	}
	for i := 0; i < elementsToPop; i++ {
		elements = append(elements, heap.Pop(t.Heap).(int))
	}
	return elements
}
