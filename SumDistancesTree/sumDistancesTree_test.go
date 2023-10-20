package sumdistancestree

import (
	"reflect"
	"testing"
	"time"

	"problemsSolving/utils"
)

type TestCase struct {
	n      int
	edges  [][]int
	expect []int
}

func TestSumDistancesTree(t *testing.T) {
	defer utils.MemoryAlloc()
	defer utils.LogRuntime(time.Now())
	tests := []TestCase{
		{
			n:      1,
			edges:  [][]int{},
			expect: []int{0},
		},
		{
			n: 2,
			edges: [][]int{
				{1, 0},
			},
			expect: []int{1, 1},
		},
		{
			n:      6,
			edges:  [][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}, {2, 5}},
			expect: []int{8, 12, 6, 10, 10, 10},
		},
	}
	for _, test := range tests {
		res := sumOfDistancesInTree(test.n, test.edges)
		if !reflect.DeepEqual(test.expect, res) {
			t.Errorf("Expect %v, result %v at test [%+v]", test.expect, res, test)
		}
	}
}
