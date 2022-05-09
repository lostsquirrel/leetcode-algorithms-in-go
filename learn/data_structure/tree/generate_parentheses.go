package tree

import "fmt"

type generatorItem struct {
	ps     []rune
	amount [2]int
	n      int
}

func (item *generatorItem) isValid() bool {
	if item.amount[1] > 0 && item.amount[0] < item.amount[1] {
		return false
	}
	return item.amount[0] <= item.n && item.amount[1] <= item.n
}

func (item *generatorItem) isCompleted() bool {
	return len(item.ps) >= 2*item.n
}
func (item *generatorItem) add(a rune) generatorItem {
	x := generatorItem{
		n:      item.n,
		ps:     append(item.ps, a),
		amount: [2]int{item.amount[0], item.amount[1]},
	}
	x.amount[a-40] += 1
	return x
}
func (item *generatorItem) String() string {
	return string(item.ps)
}

func generateParenthesis(n int) []string {
	p := []rune("()")
	// left, right := p[0], p[1]
	c := generatorItem{
		n: n,
	}
	var s []string
	generator(c, rune(p[0]))
	return s
}

func generator(c generatorItem, a rune) {
	c = c.add(a)
	fmt.Println(c.String())
	if !c.isValid() {
		fmt.Printf("invalid %s\n", c.String())
		return
	}
	if c.isCompleted() {
		fmt.Printf("complete %s\n", c.String())
		return
	}
	for _, p := range "()" {
		generator(c, rune(p))
	}
}
