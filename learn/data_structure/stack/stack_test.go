package stack

import (
	"fmt"
	"testing"
)

func TestStringToInt(t *testing.T) {
	s := "abcdefghijklmnopqrstuv#"
	for _, c := range s {
		fmt.Println(c)
	}
}
func TestIntToString(t *testing.T) {
	fmt.Println(string(rune(35)))
}

func TestIntStack(t *testing.T) {
	s := NewIntStack(3)
	_, e := s.Pop()
	if e == nil {
		t.Error("pop empty should be error")
	}
	e = s.Push(1)
	if e != nil {
		t.Error("push 1 to size 3 should success")
	}
	s.Push(2)
	s.Push(3)
	e = s.Push(4)
	if e == nil {
		t.Error("push to full stack shuld be error")
	}
}
