package queue

import "fmt"

type MyCircularQueue struct {
	capacity int
	data []int
	size int
	head int
	tail int

}




/** Initialize your data structure here. Set the size of the queue to be k. */
func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		capacity:k,
		data: make([]int, k),
		size: 0,
		head: 0,
		tail: -1,
	}
}

/** Insert an element into the circular queue. Return true if the operation is successful. */
func (q *MyCircularQueue) EnQueue(value int) bool {
	if q.size + 1 > q.capacity {
		return false
	}
	q.size += 1
	q.tail = (q.tail + 1) % q.capacity
	q.data[q.tail] = value
	return true
}


/** Delete an element from the circular queue. Return true if the operation is successful. */
func (q *MyCircularQueue) DeQueue() bool {
	if q.size < 1 {
		return false
	}
	q.size -= 1
	q.data[q.head] = 0
	q.head = (q.head + 1)  % q.capacity
	return true
}


/** Get the front item from the queue. */
func (q *MyCircularQueue) Front() int {
	if q.size == 0 {
		return -1
	}
	return q.data[q.head]
}


/** Get the last item from the queue. */
func (q *MyCircularQueue) Rear() int {
	if q.size == 0 {
		return -1
	}
	return q.data[q.tail]
}


/** Checks whether the circular queue is empty or not. */
func (q *MyCircularQueue) IsEmpty() bool {
	return q.size == 0
}


/** Checks whether the circular queue is full or not. */
func (q *MyCircularQueue) IsFull() bool {
	return q.size == q.capacity
}

func (q MyCircularQueue) String() string {
	return fmt.Sprintf("data: %v, capacity: %d, size: %d, head: %d, tail: %d", q.data, q.capacity, q.size, q.head, q.tail)
}
