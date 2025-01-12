package main

import (
	"fmt"
)

var withSubmit = true

func runTestCases() {
	tests := []struct {
		messages []string
		expect   []string
	}{
		{
			messages: []string{"Sunlit", "Man"},
			expect:   []string{"Man-processed", "Sunlit-processed"},
		},
		{
			messages: []string{"Nomad do you copy?"},
			expect:   []string{"Nomad do you copy?-processed"},
		},
		{
			messages: []string{"Scadriel", "Roshar", "Sel", "Nalthis", "Taldain"},
			expect:   []string{"Taldain-processed", "Roshar-processed", "Sel-processed", "Nalthis-processed", "Scadriel-processed"},
		},
	}

	if withSubmit {
		tests = append(tests,
			struct {
				messages []string
				expect   []string
			}{
				messages: []string{},
				expect:   []string{},
			},
			struct {
				messages []string
				expect   []string
			}{
				messages: []string{"Scadriel"},
				expect:   []string{"Scadriel-processed"},
			},
		)
	}

	passCount := 0
	failCount := 0

	for _, tc := range tests {
		fail := false
		result := processMessages(tc.messages)

		if len(result) != len(tc.expect) {
			fail = true
		}

		counts := make(map[string]int)
		for _, res := range result {
			counts[res]++
		}
		for _, exp := range tc.expect {
			counts[exp]--
			if counts[exp] < 0 {
				fail = true
			}
		}

		if fail {
			failCount++
			fmt.Printf(`
---------------------------------
Test Failed:
  inputs:   %v
  expected: %v
  actual:   %v
  `, tc.messages, tc.expect, result)
		} else {
			passCount++
			fmt.Printf(`
---------------------------------
Test Passed:
  inputs:   %v
  expected: %v
  actual:   %v
`, tc.messages, tc.expect, result)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}