package sort

import "testing"

func TestStringCompare(t *testing.T) {
	if "A" > "B" {
		t.Error("A < B")
	}
}

func TestSortASC(t *testing.T) {
	case0 := []string{"A", "B", "C", "D"}
	if !sorted(&case0) {
		t.Fail()
	}
}

func TestSortedDESC(t *testing.T) {
	case0 := []string{"D", "C", "B", "A" }
	if !sorted(&case0) {
		t.Fail()
	}
}

func TestSortedNotInOrder(t *testing.T) {
	case0 := []string{"D", "B", "A", "C" }
	if sorted(&case0) {
		t.Fail()
	}
}