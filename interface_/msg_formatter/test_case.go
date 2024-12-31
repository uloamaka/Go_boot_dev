package main

import (
	"fmt"
	"strconv"
)

var withSubmit = true

func runTestCases() {
	tests := []struct {
		formatter Formatter
		expected  string
	}{
		{PlainText{message: "Hello, World!"}, "Hello, World!"},
		{Bold{message: "Bold Message"}, "**Bold Message**"},
		{Code{message: "Code Message"}, "`Code Message`"},
	}

	if withSubmit {
		tests = append(tests,
			struct {
				formatter Formatter
				expected  string
			}{Code{message: ""}, "``"},
			struct {
				formatter Formatter
				expected  string
			}{Bold{message: ""}, "****"},
			struct {
				formatter Formatter
				expected  string
			}{PlainText{message: ""}, ""},
		)
	}

	passCount := 0
	failCount := 0

	for i, test := range tests {
		testName := "Test Case " + strconv.Itoa(i+1)
		formattedMessage := SendMessage(test.formatter)
		if formattedMessage != test.expected {
			failCount++
			fmt.Printf(`---------------------------------
%s
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Fail
`, testName, test.formatter, test.expected, formattedMessage)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
%s
Inputs:     (%v)
Expecting:  %v
Actual:     %v
Pass
`, testName, test.formatter, test.expected, formattedMessage)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}
