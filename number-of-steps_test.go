package leetcode_test

import (
	"testing"

	lc "github.com/yogavredizon/LeetCodeTraining"
)

func TestNumberOfSteps(t *testing.T) {
	testTable := []struct {
		name            string
		value           int
		expectedOutCome int
	}{
		{name: "Divide by even number", value: 14, expectedOutCome: 6},
		{name: "Divide by odd number", value: 8, expectedOutCome: 4},
		{name: "Divide by odd number", value: 123, expectedOutCome: 12},
	}

	for _, tt := range testTable {
		t.Run(tt.name, func(t *testing.T) {
			want := lc.NumberOfSteps(tt.value)

			if want != tt.expectedOutCome {
				t.Errorf("Get : %v, Expected : %v", want, tt.expectedOutCome)
			}
		})
	}

}
