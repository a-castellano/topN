package top

import "testing"
import "math/rand"

func TestConstructorSize0(t *testing.T) {

	top, _ := NewTop(0)
	if top.Size != 0 {
		t.Errorf("NewTop(0) should return a topN instance with size 0, but returns %d", top.Size)
	}
}

func TestConstructorAnySizeOver0(t *testing.T) {
	rand.Seed(42)

	size := rand.Int()
	top, _ := NewTop(size)
	if top.Size != size {
		t.Errorf("NewTop(0) should return a topN instance with size %d, but returns %d", size, top.Size)
	}
}

func TestConstructorSizeUnder0(t *testing.T) {
	rand.Seed(42)

	size := rand.Int() % 100
	size = -size
	_, err := NewTop(size)
	if err == nil {
		t.Errorf("Should return an error if size is under 0. With  %d it doesn't", size)
	}
}

func TestShowEmptyTop(t *testing.T) {
	top, _ := NewTop(10)
	erroredElement, err := top.ShowMax()
	if erroredElement != 0 {
		t.Errorf("With empty top, max element should be 0.")
	}
	if err == nil {
		t.Errorf("With empty top, return an error.")
	}
}

func TestShowWithPush(t *testing.T) {
	top, _ := NewTop(10)
	top.Push(3)
	max, _ := top.ShowMax()
	if max != 3 {
		t.Errorf("Max element should be 3, not %d", max)
	}

	top.Push(2)
	max, _ = top.ShowMax()
	if max == 2 {
		t.Errorf("Max element should not be 2, not %d", max)
	}

	top.Push(4)
	max, _ = top.ShowMax()
	if max != 4 {
		t.Errorf("Max element should be 4, not %d", max)
	}
}

func TestPopNoElements(t *testing.T) {
	top, _ := NewTop(10)
	numberOfElements := len(top.PopElements())
	if numberOfElements != 0 {
		t.Errorf("PopElements of empty top should return 0, not %d.", numberOfElements)
	}
}

func TestPopLessElementsThanSize(t *testing.T) {
	top, _ := NewTop(10)
	top.Push(2)
	top.Push(3)
	top.Push(4)
	top.Push(2)
	numberOfElements := len(top.PopElements())
	if numberOfElements != 4 {
		t.Errorf("PopElements of this top should return 4, not %d.", numberOfElements)
	}
}

func TestPopSameElementsThanSize(t *testing.T) {
	top, _ := NewTop(10)
	top.Push(1)
	top.Push(2)
	top.Push(3)
	top.Push(4)
	top.Push(5)
	top.Push(6)
	top.Push(7)
	top.Push(8)
	top.Push(9)
	top.Push(10)
	numberOfElements := len(top.PopElements())
	if numberOfElements != 10 {
		t.Errorf("PopElements of this top should return 10, not %d.", numberOfElements)
	}
}

func TestPopMoreElementsThanSize(t *testing.T) {
	top, _ := NewTop(10)
	top.Push(1)
	top.Push(2)
	top.Push(3)
	top.Push(4)
	top.Push(1)
	top.Push(5)
	top.Push(6)
	top.Push(7)
	top.Push(8)
	top.Push(9)
	top.Push(10)
	elements := top.PopElements()
	numberOfElements := len(elements)
	if numberOfElements != 10 {
		t.Errorf("PopElements of this top should return 10, not %d.", numberOfElements)
	}

	if elements[0] != 10 {
		t.Errorf("PopElements first element should be 10, not %d.", elements[0])
	}

	if elements[1] != 9 {
		t.Errorf("PopElements second element should be 9, not %d.", elements[1])
	}

}

func TestActualSizeLessElementsThanSize(t *testing.T) {
	top, _ := NewTop(3)
	if top.Heap.Len() != 0 {
		t.Errorf("Empty topN should have 0 elements, it hasn't.")
	}
	top.Push(1)
	if top.Heap.Len() != 1 {
		t.Errorf("After pushing one element topN should have 1 elements, it hasn't.")
	}
	top.Push(1)
	if top.Heap.Len() != 2 {
		t.Errorf("After pushing second element topN should have 2 elements, it hasn't.")
	}
	top.Push(1)
	if top.Heap.Len() != 3 {
		t.Errorf("After pushing third element topN should have 3 elements, it hasn't.")
	}
	top.Push(1)
	if top.Heap.Len() != 3 {
		t.Errorf("After pushing more than three elements, topN should maintain only 3 elements, it hasn't.")
	}
}
