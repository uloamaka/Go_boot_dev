package main

import (
	"fmt"
	"slices"
)

var withSubmit = true

func runTestCases() {
	type testCase struct {
		n        int
		expected []int
	}
	tests := []testCase{
		{5, []int{0, 1, 1, 2, 3}},
		{3, []int{0, 1, 1}},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{0, []int{}},
			{1, []int{0}},
			{7, []int{0, 1, 1, 2, 3, 5, 8}},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		actual := concurrentFib(test.n)
		if !slices.Equal(actual, test.expected) {
			failCount++
			fmt.Printf(`
---------------------------------
Test Failed:
  n:        %v
  expected: %v
  actual:   %v
`, test.n, test.expected, actual)
		} else {
			passCount++
			fmt.Printf(`
---------------------------------
Test Passed:
  n:        %v
  expected: %v
  actual:   %v
`, test.n, test.expected, actual)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}