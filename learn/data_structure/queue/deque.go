package queue

type Node struct {
	val  int
	prev *Node
	next *Node
}
type IntDeque struct {
	head *Node
	tail *Node
	size int
}

func (q *IntDeque) PushFront(val int) {
	n := &Node{
		val: val,
	}
	if q.head == nil {
		q.head = n
		q.tail = n
	} else {
		h := q.head
		h.prev = n
		n.next = h
		q.head = n
	}
	q.size++
}

func (q *IntDeque) PushBack(val int) {
	n := &Node{
		val: val,
	}
	if q.tail == nil {
		q.tail = n
		q.head = n
	} else {
		t := q.tail
		n.prev = t
		t.next = n
		q.tail = n
	}
	q.size++
}

func (q *IntDeque) PeekFront() int {
	// TODO: check empty
	return q.head.val
}

func (q *IntDeque) PeekBack() int {
	return q.tail.val
}

func (q *IntDeque) PopFront() int {
	// TODO: check empty
	r := q.head
	if q.head != nil {
		q.head = q.head.next

		if q.head == nil {
			q.tail = nil
		} else {
			q.head.prev = nil
		}
	}
	q.size--
	return r.val
}

func (q *IntDeque) PopBack() int {
	// TODO: check empty
	r := q.tail
	if q.tail != nil {
		q.tail = q.tail.prev

		if q.tail == nil {
			q.head = nil
		} else {
			q.tail.next = nil
		}
	}
	q.size--
	return r.val
}

func (q *IntDeque) Empty() bool {
	return q.size == 0
}
