package top

import topheap "github.com/a-castellano/topN/heap"

import "container/heap"

type TopN struct {
	Heap topheap.MaxHeap *
	Size int
}

function NewMatrix
