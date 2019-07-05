package queue

import (
	"fmt"
	"strconv"
	"strings"
)

func openLock(deadends []string, target string) int {
	var tryOnAsc [4]bool // try to find the number from 1-9 otherwise  9-1
	targetArray := toArray(target)
	tryArray := make([]int, 4)
	// init the try direction
	for i, v := range targetArray {
		tryOnAsc[i] = v < 5
	}
	if isDead(toString(&tryArray), deadends) {
		return -1
	}

	var step int
	// try each slot
	for i, v := range targetArray {
		var deadCount int
		if tryOnAsc[i] {

			for j := 1; j <= v; j++ {
				tryArray[i] = j
				if isDead(toString(&tryArray), deadends) {
					deadCount += 1
					break
				}
			}
			fmt.Printf("deadCount = %d\n", deadCount)
			if deadCount > 0 {
				for j := 9; j >= v; j-- {
					fmt.Printf("%d, %d\n", i, j)
					tryArray[i] = j
					if isDead(toString(&tryArray), deadends) {
						deadCount += 1
						break
					} else {
						step += 1
					}
				}
			}

			if deadCount > 1 {
				step = -1
				break
			}
			step += v
			fmt.Printf("%s %s\n", "asc", toString(&tryArray))
		} else {
			for j := 9; j >= v; j-- {
				tryArray[i] = j
				if isDead(toString(&tryArray), deadends) {
					deadCount += 1
					break
				}
			}
			if deadCount > 0 {
				for j := 1; j <= v; j++ {
					tryArray[i] = j
					if isDead(toString(&tryArray), deadends) {
						deadCount += 1
						break
					}
				}
			}
			if deadCount > 1 {
				step = -1
				break
			}
			step += 10 -v
			fmt.Printf("%s %s\n", "desc", toString(&tryArray))
		}

	}
	return step
}

func isDead(s string, deadends []string) bool {
	for _, v := range deadends {
		if s == v {
			fmt.Println(s)
			return true
		}
	}
	return false
}

func toArray(try string) []int {
	strArray := strings.Split(try, "")
	//size := len(strArray)
	intArray := make([]int, 0);
	for _, v :=range strArray {
		i, e := strconv.Atoi(v)
		if e == nil {
			intArray = append(intArray, i)
		}

	}

	return intArray
}

func toString(s *[]int) string {
	var b strings.Builder
	for _, v :=range *s {
		fmt.Fprintf(&b, "%d", v)
	}
	return b.String()
}