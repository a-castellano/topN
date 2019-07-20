package top

import "testing"
import "math/rand"

func TestConstructorSize0(t *testing.T) {

	top := NewTop(0)
	if top.Size != 0 {
		t.Errorf("NewTop(0) should return a topN instance with size 0, but returns %d", top.Size)
	}
}

func TestConstructorSize1(t *testing.T) {
	rand.Seed(42)

	size := rand.Int()
	top := NewTop(size)
	if top.Size != size {
		t.Errorf("NewTop(0) should return a topN instance with size %d, but returns %d", size, top.Size)
	}
}
