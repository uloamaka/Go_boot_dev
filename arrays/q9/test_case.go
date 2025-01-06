package main

import (
	"fmt"
)

var withSubmit = true

func runTestCases() {
	type testCase struct {
		password string
		isValid  bool
	}
	testCases := []testCase{
		{"Pass123", true},
		{"pas", false},
		{"Password", false},
		{"123456", false},
	}
	if withSubmit {
		testCases = append(testCases,
			[]testCase{
				{"VeryLongPassword1", false},
				{"Short", false},
				{"1234short", false},
				{"Short5", true},
				{"P4ssword", true},
			}...,
		)
	}

	passCount := 0
	failCount := 0

	for _, tc := range testCases {
			result := isValidPassword(tc.password)
			if result != tc.isValid {
				failCount++
				fmt.Printf(`---------------------------------
Password:  "%s"
Expecting: %v
Actual:    %v
Fail`, tc.password, tc.isValid, result)
			} else {
				passCount++
				fmt.Printf(`---------------------------------
Password:  "%s"
Expecting: %v
Actual:    %v
Pass
`, tc.password, tc.isValid, result)
			}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}