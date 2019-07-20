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
