package top

import "testing"

func TestConstructorSize0(t *testing.T) {

	top := NewTop(0)
	if top.Size != 0 {
		t.Errorf("NewTop(0) should return a topN instance with size 0, but returns %d", top.Size)
	}
}
