package main

import (
	"fmt"
)

var withSubmit = true
func runTestCases() {
	type testCase struct {
		msgToCustomer string
		msgToSpouse   string
		expectedCost  int
		expectedErr   error
	}
	tests := []testCase{
		{"Thanks for coming in to our flower shop today!", "We hope you enjoyed your gift.", 0, fmt.Errorf("can't send texts over 25 characters")},
		{"Thanks for joining us!", "Have a good day.", 76, nil},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{"Thank you.", "Enjoy!", 32, nil},
			{"We loved having you in!", "We hope the rest of your evening is fantastic.", 0, fmt.Errorf("can't send texts over 25 characters")},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		cost, err := sendSMSToCouple(test.msgToCustomer, test.msgToSpouse)
		errString := ""
		if err != nil {
			errString = err.Error()
		}
		expectedErrString := ""
		if test.expectedErr != nil {
			expectedErrString = test.expectedErr.Error()
		}
		if cost != test.expectedCost || errString != expectedErrString {
			failCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Fail`, test.msgToCustomer, test.msgToSpouse, test.expectedCost, test.expectedErr, cost, err)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  (%v, %v)
Actual:     (%v, %v)
Pass
`, test.msgToCustomer, test.msgToSpouse, test.expectedCost, test.expectedErr, cost, err)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}