package queue

import (
	"fmt"
	"strings"
	"testing"
)

func TestStringSplit(t *testing.T) {
	strArray := strings.Split("0000", "")
	for i, v := range strArray {
		fmt.Printf("%d => %T\n", i, v)
	}
	fmt.Println(len(strArray))
}

func TestToArray(t *testing.T) {
	a := toArray("0000")
	for i, v := range a {
		fmt.Printf("%d => %d\n", i, v)
	}
}

func TestToString(t *testing.T) {
	ints := []int{0, 0, 0, 0}
	if "0000" != toString(&ints) {
		t.Fail()
	}
}

func TestOpen1(t *testing.T) {
	deadends := []string{"0201","0101","0102","1212","2002"}
	target := "0202"
	result := openLock(deadends, target)
	fmt.Println(result)
	if 6 != result {
		t.Fail()
	}
}