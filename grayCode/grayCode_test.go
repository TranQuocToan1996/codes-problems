package graycode

import (
	"reflect"
	"testing"
	"time"

	"problemsSolving/utils"
)

type GrayCodeTestCase struct {
	input  int
	expect []int
}

func TestGrayCode(t *testing.T) {
	defer utils.LogRuntime(time.Now())
	tests := []GrayCodeTestCase{
		{
			input:  1,
			expect: []int{0, 1},
		},
		{
			input:  2,
			expect: []int{0, 1, 3, 2},
		},
	}
	for _, test := range tests {
		if !reflect.DeepEqual(test.expect, grayCode(test.input)) {
			t.Errorf("Expect %v, result %v", test.expect, grayCode(test.input))
		}
	}
}
