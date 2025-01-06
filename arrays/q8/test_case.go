package main

import (
	"fmt"
)

var withSubmit = true

func runTestCases() {
	messages := []Message{
		TextMessage{"Alice", "Hello, World!"},
		MediaMessage{"Bob", "image", "A beautiful sunset"},
		LinkMessage{"Charlie", "http://example.com", "Example Domain"},
		TextMessage{"Dave", "Another text message"},
		MediaMessage{"Eve", "video", "Cute cat video"},
		LinkMessage{"Frank", "https://boot.dev", "Learn Coding Online"},
	}

	testCases := []struct {
		filterType    string
		expectedCount int
		expectedType  string
	}{
		{"text", 2, "text"},
		{"media", 2, "media"},
		{"link", 2, "link"},
	}

	if withSubmit {
		testCases = append(testCases,
			struct {
				filterType    string
				expectedCount int
				expectedType  string
			}{"media", 2, "media"},
			struct {
				filterType    string
				expectedCount int
				expectedType  string
			}{"text", 2, "text"},
		)
	}

	passCount := 0
	failCount := 0

	for idx, tc := range testCases {
		filtered := filterMessages(messages, tc.filterType)

		if len(filtered) != tc.expectedCount {
			failCount++
			fmt.Printf(`---------------------------------
Test Case %d - Filtering for "%s"
Expecting:  %d messages
Actual:     %d messages
Fail
`, idx+1, tc.filterType, tc.expectedCount, len(filtered))
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Case %d - Filtering for "%s"
Expecting:  %d messages
Actual:     %d messages
Pass
`, idx+1, tc.filterType, tc.expectedCount, len(filtered))
		}

		for _, m := range filtered {
			if m.Type() != tc.expectedType {
				failCount++
				fmt.Printf(`---------------------------------
Test Case %d - Message Type Check
Expecting:  "%s" message
Actual:     "%s" message
Fail
`, idx+1, tc.expectedType, m.Type())
			} else {
				passCount++
				fmt.Printf(`---------------------------------
Test Case %d - Message Type Check
Expecting:  "%s" message
Actual:     "%s" message
Pass
`, idx+1, tc.expectedType, m.Type())
			}
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}
