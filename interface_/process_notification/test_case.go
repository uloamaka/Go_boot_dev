package main

import (
	"fmt"
	"strconv"
)

var withSubmit = true
func runTestCases() {
	tests := []struct {
		notification       notification
		expectedID         string
		expectedImportance int
	}{
		{
			directMessage{senderUsername: "Kaladin", messageContent: "Life before death", priorityLevel: 10, isUrgent: true},
			"Kaladin",
			50,
		},
		{
			groupMessage{groupName: "Bridge 4", messageContent: "Soups ready!", priorityLevel: 2},
			"Bridge 4",
			2,
		},
		{
			systemAlert{alertCode: "ALERT001", messageContent: "THIS IS NOT A TEST HIGH STORM COMING SOON"},
			"ALERT001",
			100,
		},
	}

	if withSubmit {
		tests = append(tests,
			struct {
				notification       notification
				expectedID         string
				expectedImportance int
			}{
				directMessage{senderUsername: "Shallan", messageContent: "I am that I am.", priorityLevel: 5, isUrgent: false},
				"Shallan",
				5,
			},
			struct {
				notification       notification
				expectedID         string
				expectedImportance int
			}{
				groupMessage{groupName: "Knights Radiant", messageContent: "For the greater good.", priorityLevel: 10},
				"Knights Radiant",
				10,
			},
			struct {
				notification       notification
				expectedID         string
				expectedImportance int
			}{
				directMessage{senderUsername: "Adolin", messageContent: "Duels are my favorite.", priorityLevel: 3, isUrgent: true},
				"Adolin",
				50,
			},
		)
	}

	passCount := 0
	failCount := 0

	for i, test := range tests {
		testName := "TestProcessNotification_" + strconv.Itoa(i+1)
		id, importance := processNotification(test.notification)

		if id != test.expectedID || importance != test.expectedImportance {
			failCount++
			fmt.Printf(`---------------------------------
Test Failed: %s
Notification: %+v
Expecting:    %v/%d
Actual:       %v/%d
Fail
`, testName, test.notification, test.expectedID, test.expectedImportance, id, importance)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Test Passed: %s
Notification: %+v
Expecting:    %v/%d
Actual:       %v/%d
Pass
`, testName, test.notification, test.expectedID, test.expectedImportance, id, importance)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}
