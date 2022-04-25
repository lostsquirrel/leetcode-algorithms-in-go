package stack

type MyQueue struct {
	s1 *IntStack
	s2 *IntStack
}

func Constructor() MyQueue {
	var size = 10
	return MyQueue{
		s1: NewIntStack(size),
		s2: NewIntStack(size),
	}
}

func (q *MyQueue) Push(x int) {
	move(q.s2, q.s1)
	q.s1.Push(x)
}

func (q *MyQueue) Pop() int {
	move(q.s1, q.s2)
	r, _ := q.s2.Pop()
	return r
}

func (q *MyQueue) Peek() int {
	move(q.s1, q.s2)
	r, _ := q.s2.Top()
	return r
}

func (q *MyQueue) Empty() bool {
	return q.s1.Empty() && q.s2.Empty()
}

func move(source *IntStack, target *IntStack) {

	if !source.Empty() {
		for {
			r, e := source.Pop()
			if e == nil {
				target.Push(r)
			} else {
				break
			}
		}
	}

}
