package main

import (
	"fmt"
)

var withSubmit = true


func runTestCases() {
	tests := []struct {
		name           string
		membershipType string
	}{
		{"Syl", "standard"},
		{"Pattern", "premium"},
		{"Pattern", "standard"},
	}
	if withSubmit {
		submitCases := []struct {
			name           string
			membershipType string
		}{
			{"Renarin", "standard"},
			{"Lift", "premium"},
			{"Dalinar", "standard"},
		}
		tests = append(tests, submitCases...)
	}

	passCount := 0
	failCount := 0

	for _, tc := range tests {
		user := newUser(tc.name, tc.membershipType)

		msgCharLimit := 100
		if tc.membershipType == "premium" {
			msgCharLimit = 1000
		}

		if user.Name != tc.name {
			failCount++
			fmt.Printf(`---------------------------------
Test Failed (name):
Inputs:     (name: %v, membershipType: %v)
Expecting:  %v
Actual:     %v
`, tc.name, tc.membershipType, tc.name, user.Name)
		} else if user.Type != tc.membershipType {
			failCount++
			fmt.Printf(`---------------------------------
Test Failed (membership type):
Inputs:     (name: %v, membershipType: %v)
Expecting:  %v
Actual:     %v
`, tc.name, tc.membershipType, tc.membershipType, user.Type)
		} else if user.MessageCharLimit != msgCharLimit {
			failCount++
			fmt.Printf(`---------------------------------
Test Failed (message character limit):
Inputs:     (name: %v, membershipType: %v)
Expecting:  %v
Actual:     %v
`, tc.name, tc.membershipType, msgCharLimit, user.MessageCharLimit)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed:
Inputs:     (name: %v, membershipType: %v)
Expecting:  %v, %v, %v
Actual:     %v, %v, %v
`, tc.name, tc.membershipType, tc.name, tc.membershipType, msgCharLimit, user.Name, user.Type, user.MessageCharLimit)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}