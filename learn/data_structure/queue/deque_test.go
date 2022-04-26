package queue

import "testing"

func TestDeque(t *testing.T) {
	q := &IntDeque{}

	q.PushBack(1)

	check(q.PeekBack(), 1, t)
	check(q.PeekFront(), 1, t)
	q.PushBack(2)
	q.PushBack(3)
	q.PushBack(4)
	check(q.PeekFront(), 1, t)
	check(q.PeekBack(), 4, t)
	check(q.PopBack(), 4, t)
	check(q.PopBack(), 3, t)
	check(q.PeekBack(), 2, t)
	check(q.PopFront(), 1, t)
	check(q.PeekFront(), 2, t)
}

func check(x, e int, t *testing.T) {
	if x != e {
		t.Errorf("expected %d got %d", e, x)
	}
}
