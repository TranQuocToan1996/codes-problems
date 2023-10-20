package maxlenrepeatedsubarray

import (
	"reflect"
	"testing"
	"time"

	"problemsSolving/utils"
)

type TestCase struct {
	nums1, nums2 []int
	expect       int
}

func TestMaxLenRepeatedSubarray(t *testing.T) {
	defer utils.MemoryAlloc()
	defer utils.LogRuntime(time.Now())
	tests := []TestCase{
		{
			nums1:  []int{1, 2, 3, 2, 1},
			nums2:  []int{3, 2, 1, 4, 7},
			expect: 3, //[]int{3,2,1}
		},
		{
			nums1:  []int{0, 0, 0, 0, 0},
			nums2:  []int{0, 0, 0, 0, 0},
			expect: 5, //[]int{0, 0, 0, 0, 0}
		},
		{
			nums1:  []int{70, 39, 25, 40, 7},
			nums2:  []int{52, 20, 67, 5, 31},
			expect: 0,
		},
	}
	for _, test := range tests {
		res := findLength(test.nums1, test.nums2)
		if !reflect.DeepEqual(test.expect, res) {
			t.Errorf("Expect %v, result %v at test [%+v]", test.expect, res, test)
		}
	}
}
