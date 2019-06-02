package queue

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {

	fmt.Println(Constructor(3))
}

func TestMyCircularQueue_EnQueue(t *testing.T) {
	cq := Constructor(3)
	if !cq.EnQueue(1) {
		t.Fail()
	}
	fmt.Println(cq)
	if !cq.EnQueue(2) {
		t.Fail()
	}
	fmt.Println(cq)
	if !cq.EnQueue(3) {
		t.Fail()
	}
	fmt.Println(cq)
	if cq.EnQueue(5) {
		t.Fail()
	}
	fmt.Println(cq)
}

func TestMyCircularQueue_DeQueue(t *testing.T) {
	cq := Constructor(3)
	if !cq.EnQueue(1) {
		t.Fail()
	}
	fmt.Println(cq)
	if !cq.EnQueue(2) {
		t.Fail()
	}
	fmt.Println(cq)
	if !cq.EnQueue(3) {
		t.Fail()
	}
	fmt.Println(cq)
	if cq.EnQueue(5) {
		t.Fail()
	}
	fmt.Println(cq)
	if !cq.DeQueue() {
		t.Fail()
	}
	fmt.Println(cq)
	if !cq.DeQueue() {
		t.Fail()
	}
	fmt.Println(cq)
	if !cq.DeQueue() {
		t.Fail()
	}
	fmt.Println(cq)
	if cq.DeQueue() {
		t.Fail()
	}
	fmt.Println(cq)
	if !cq.EnQueue(1) {
		t.Error("68")
	}
	fmt.Println(cq)
	if !cq.EnQueue(2) {
		t.Error("72")
	}
	fmt.Println(cq)
	if !cq.EnQueue(3) {
		t.Fail()
	}
	fmt.Println(cq)
	if !cq.DeQueue() {
		t.Fail()
	}
	fmt.Println(cq)
	if !cq.DeQueue() {
		t.Fail()
	}
	fmt.Println(cq)
	if !cq.DeQueue() {
		t.Fail()
	}
	fmt.Println(cq)
	if cq.DeQueue() {
		t.Fail()
	}
	fmt.Println(cq)

	if !cq.EnQueue(1) {
		t.Error("68")
	}
	fmt.Println(cq)
	if !cq.DeQueue() {
		t.Fail()
	}
	fmt.Println(cq)
	if !cq.EnQueue(2) {
		t.Error("72")
	}
	fmt.Println(cq)
	if !cq.DeQueue() {
		t.Fail()
	}
	fmt.Println(cq)
	if !cq.EnQueue(3) {
		t.Fail()
	}
	fmt.Println(cq)
	if !cq.DeQueue() {
		t.Fail()
	}
	fmt.Println(cq)
}