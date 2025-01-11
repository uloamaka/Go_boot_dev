package main

import (
	"fmt"
)

var withSubmit = true

func runTestCases() {
	type testCase struct {
		numBatches int
		expected   int
	}
	tests := []testCase{
		{3, 114},
		{4, 198},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{0, 0},
			{1, 15},
			{6, 435},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		numSentCh := make(chan int)
		go sendReports(test.numBatches, numSentCh)
		output := countReports(numSentCh)
		if output != test.expected {
			failCount++
			fmt.Printf(`
---------------------------------
Test Failed:
  numBatches: %v
  expected:   %v
  actual:     %v
`, test.numBatches, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`
---------------------------------
Test Passed:
  numBatches: %v
  expected:   %v
  actual:     %v
`, test.numBatches, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}