package queue

type Node struct {
	val  int
	prev *Node
	next *Node
}
type IntDeque struct {
	head *Node
	tail *Node
}

func (q *IntDeque) PushFront(val int) {
	n := &Node{
		val:  val,
		next: q.head,
	}
	q.head = n
}

func (q *IntDeque) PushBack(val int) {
	n := &Node{
		val:  val,
		prev: q.tail,
	}
	q.tail = n
}

func (q *IntDeque) PeekFront() int {
	// TODO: check empty
	return q.head.val
}

func (q *IntDeque) PopFront() int {
	// TODO: check empty
	r := q.head
	q.head = q.head.next
	return r.val
}

func (q *IntDeque) PopBack() int {
	// TODO: check empty
	r := q.tail
	q.tail = q.tail.next
	return r.val
}
