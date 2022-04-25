package heap

import "testing"

func TestIntHeap(t *testing.T) {
	h := NewIntHeap(5)
	h.Insert(1)
	x := h.Min()
	expected := 1
	if x != expected {
		t.Errorf("expected %d got %d", expected, x)
	}
	h.Insert(2)
	x = h.Min()
	expected = 1
	if x != expected {
		t.Errorf("expected %d got %d", expected, x)
	}
	h.Insert(3)
	h.Insert(4)
	h.Insert(5)
	x = h.Min()
	expected = 1
	if x != expected {
		t.Errorf("expected %d got %d", expected, x)
	}
	h.Insert(6)
	x = h.Min()
	expected = 2
	if x != expected {
		t.Errorf("expected %d got %d", expected, x)
	}
	h.Insert(7)
	h.Insert(8)
	h.Insert(9)
	x = h.Min()
	expected = 5
	if x != expected {
		t.Errorf("expected %d got %d", expected, x)
	}
}
func TestIntHeap2(t *testing.T) {
	h := NewIntHeap(3)
	init := []int{5, 4, 8, 2}
	for _, v := range init {
		h.Insert(v)
	}
	h.Insert(3)
	check(h.Min(), 4, t)
	h.Insert(5)
	check(h.Min(), 5, t)
	h.Insert(10)
	check(h.Min(), 5, t)
	h.Insert(9)
	check(h.Min(), 8, t)
	h.Insert(4)
	check(h.Min(), 8, t)
}
func TestIntHeap3(t *testing.T) {
	h := NewIntHeap(2)
	h.Insert(0)
	// h.Insert(0)
	h.Insert(-1)
	check(h.Min(), -1, t)
	h.Insert(1)
	check(h.Min(), 0, t)
	h.Insert(-2)
	check(h.Min(), 0, t)
	h.Insert(-4)
	check(h.Min(), 0, t)
	h.Insert(3)
	check(h.Min(), 1, t)
}

func check(actual, expected int, t *testing.T) {
	if actual != expected {
		t.Errorf("expected %d got %d", expected, actual)
	}
}
