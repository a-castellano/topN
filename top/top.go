package top

import topheap "github.com/a-castellano/topN/heap"
import "container/heap"

type TopN struct {
	Heap *topheap.MaxHeap
	Size int
}

func NewTop(size int) *TopN {
	top := new(TopN)
	top.Size = size
	top.Heap = &topheap.MaxHeap{}
	heap.Init(top.Heap)

	return top
}
