package stack

import "testing"

func TestQueue(t *testing.T) {
	q := Constructor()
	q.Push(1)
	x := q.Peek()
	if x != 1 {
		t.Errorf("peek should be 1 go %d", x)
	}
	q.Push(2)
	q.Push(3)
	x = q.Peek()
	if x != 1 {
		t.Errorf("peek should be 1 go %d", x)
	}
	x = q.Pop()
	if x != 1 {
		t.Errorf("peek should be 1 go %d", x)
	}
	x = q.Peek()
	if x != 2 {
		t.Errorf("peek should be 2 go %d", x)
	}
	q.Push(4)
	q.Push(5)
	q.Pop()
	q.Pop()
	x = q.Peek()
	if x != 4 {
		t.Errorf("peek should be 5 go %d", x)
	}
}
