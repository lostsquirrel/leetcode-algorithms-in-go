package queue

import (
	"fmt"
	"testing"
)

func TestSlidingWindowDouble(t *testing.T) {
	nums := []int{1, 3, -1, -3, 5, 3, 6, 7}
	k := 3
	r := maxSlidingWindow(nums, k)
	fmt.Println(r)
}
func TestSlidingWindowDouble2(t *testing.T) {
	nums := []int{1, -1}
	k := 1
	r := maxSlidingWindow(nums, k)
	fmt.Println(r)
}
