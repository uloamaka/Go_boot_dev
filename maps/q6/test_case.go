package main

import (
	"fmt"
	"reflect"
	"sort"
)

var withSubmit = true

func runTestCases() {
	friendships := map[string][]string{
		"Dalinar": {"Kaladin", "Pattern", "Shallan"},
		"Kaladin": {"Dalinar", "Syl", "Teft", "Shallan"},
		"Pattern": {"Dalinar", "Teft", "Shallan"},
		"Syl":     {"Kaladin"},
		"Teft":    {"Kaladin", "Pattern"},
		"Moash":   {},
		"Shallan": {"Pattern", "Kaladin", "Dalinar"},
	}

	testCases := []struct {
		username string
		expected []string
	}{
		{"Dalinar", []string{"Syl", "Teft"}},
		{"Kaladin", []string{"Pattern"}},
		{"Pattern", []string{"Kaladin"}},
		{"Syl", []string{"Dalinar", "Shallan", "Teft"}},
		{"Teft", []string{"Dalinar", "Shallan", "Syl"}},
		{"Moash", nil},
	}

	if withSubmit {
		testCases = append(testCases,
			struct {
				username string
				expected []string
			}{
				"Odium", nil,
			},
			struct {
				username string
				expected []string
			}{
				"Shallan", []string{"Syl", "Teft"},
			},
		)
	}

	passCount := 0
	failCount := 0

	for _, tc := range testCases {
		result := findSuggestedFriends(tc.username, friendships)
		sort.Strings(result) 
		sort.Strings(tc.expected)

		if !reflect.DeepEqual(result, tc.expected) {
			failCount++
			fmt.Printf(`---------------------------------
Test Failed for username %s:
Expecting:  %v
Actual:     %v
Fail
`, tc.username, formatSlice(tc.expected), formatSlice(result))
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed for username %s:
Expecting:  %v
Actual:     %v
Pass
`, tc.username, formatSlice(tc.expected), formatSlice(result))
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

func formatSlice(slice []string) string {
	if slice == nil {
		return "nil"
	}
	return fmt.Sprintf("%+v", slice)
}