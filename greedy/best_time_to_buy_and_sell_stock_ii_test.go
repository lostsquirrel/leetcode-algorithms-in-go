package greedy

import "testing"

type greedy struct {
	prices   []int
	expected int
}

func TestGreedy01(t *testing.T) {
	tests := []greedy{
		{
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 7,
		},
		{
			prices:   []int{1, 2, 3, 4, 5},
			expected: 4,
		},
		{
			prices:   []int{7, 6, 4, 3, 1},
			expected: 0,
		},
	}
	for _, test := range tests {
		check(t, test.expected, maxProfit(test.prices))
	}

}

func check(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}
