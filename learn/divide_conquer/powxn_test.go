package divide_conquer

import "testing"

func TestPow(t *testing.T) {
	actual := myPow(2, 10)
	expected := 1024.0
	if actual != expected {
		t.Errorf("expected %f actual %f", expected, actual)
	}
}
