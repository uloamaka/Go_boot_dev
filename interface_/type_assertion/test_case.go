package main

import (
	"fmt"
)

var withSubmit = true

func runTestCases() {
	type testCase struct {
		expense      expense
		expectedTo   string
		expectedCost float64
	}
	tests := []testCase{
		{
			email{isSubscribed: true, body: "Whoa there!", toAddress: "soldier@monty.com"},
			"soldier@monty.com",
			0.11,
		},
		{
			sms{isSubscribed: false, body: "Halt! Who goes there?", toPhoneNumber: "+155555509832"},
			"+155555509832",
			2.1,
		},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{
				email{
					isSubscribed: false,
					body:         "It is I, Arthur, son of Uther Pendragon, from the castle of Camelot. King of the Britons, defeator of the Saxons, sovereign of all England!",
					toAddress:    "soldier@monty.com",
				},
				"soldier@monty.com",
				6.95,
			},
			{
				email{
					isSubscribed: true,
					body:         "Pull the other one!",
					toAddress:    "arthur@monty.com",
				},
				"arthur@monty.com",
				0.19,
			},
			{
				sms{
					isSubscribed:  true,
					body:          "I am. And this my trusty servant Patsy.",
					toPhoneNumber: "+155555509832",
				},
				"+155555509832",
				1.17,
			},
			{
				invalid{},
				"",
				0.0,
			},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		to, cost := getExpenseReport(test.expense)
		if to != test.expectedTo || cost != test.expectedCost {
			failCount++
			fmt.Printf(`---------------------------------
Inputs:     %+v
Expecting:  (%v, %v)
Actual:     (%v, %v)
Fail`, test.expense, test.expectedTo, test.expectedCost, to, cost)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     %+v
Expecting:  (%v, %v)
Actual:     (%v, %v)
Pass
`, test.expense, test.expectedTo, test.expectedCost, to, cost)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}