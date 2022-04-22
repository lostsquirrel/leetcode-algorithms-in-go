package stack

import "testing"

func TestBacksapceCompare(t *testing.T) {
	r := backspaceCompare("ab#c", "ad#c")
	if !r {
		t.Error("should equal")
	}
}

func TestBacksapceCompareE2(t *testing.T) {
	x := "xywrrmp"
	x1 := "xywrrmu#p"
	if !backspaceCompare(x, x1) {
		t.Errorf("%s %s equal", x, x1)
	}
}
func TestBacksapceCompareE3(t *testing.T) {
	x := "ab##"
	x1 := "c#d#"
	if !backspaceCompare(x, x1) {
		t.Errorf("%s %s equal", x, x1)
	}
}

func TestBacksapceCompareNE(t *testing.T) {
	if backspaceCompare("a#c", "b") {
		t.Error("should not equal")
	}
}
