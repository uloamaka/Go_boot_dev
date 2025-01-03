package main

import (
	"fmt"
)

var withSubmit = true
func runTestCases() {
	type testCase struct {
		thresh   int
		expected int
	}
	tests := []testCase{
		{103, 1},
		{205, 2},
		{1000, 9},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{100, 1},
			{3000, 26},
			{4000, 34},
			{5000, 41},
			{0, 0},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		output := maxMessages(test.thresh)
		if output != test.expected {
			failCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Fail`, test.thresh, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Pass
`, test.thresh, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}