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
