package other

import (
	"reflect"
	"testing"
)

func TestXxx(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9

	check(twoSum(nums, target), []int{0, 1}, t)
	nums = []int{3, 2, 4}
	check(twoSum(nums, 6), []int{1, 2}, t)
	check(twoSum([]int{3, 3}, 6), []int{0, 1}, t)
}

func check(actual, expected []int, t *testing.T) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected %d got %d", expected, actual)
	}
}
