package services_test

import (
	"testing"

	"github.com/amerikarno/exam7/services"
)

type filterTest struct {
	input, expect string
}

func TestUnitTest(t *testing.T) {

	tests := []filterTest{
		{"input1.", "input1"},
		{"input2,", "input2"},
		{"inpuT3", "input3"},
	}
	for _, test := range tests {
		if actual := services.Filter(test.input); actual != test.expect {
			t.Errorf("actual: %v, expect: %v\n", actual, test.expect)
		}
	}
}

func BenchmarkUnitTest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		services.Filter("input1.")
	}
}
