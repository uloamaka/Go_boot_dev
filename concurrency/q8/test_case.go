package main

import (
	"fmt"
	"slices"
	"time"
)

var withSubmit = true

func runTestCases() {
	expectedLogs := []string{
		"Nothing to do, waiting...",
		"Nothing to do, waiting...",
		"Taking a backup snapshot...",
		"Nothing to do, waiting...",
		"Nothing to do, waiting...",
		"Taking a backup snapshot...",
		"Nothing to do, waiting...",
		"Taking a backup snapshot...",
		"Nothing to do, waiting...",
		"All backups saved!",
	}

	snapshotTicker := time.Tick(800 * time.Millisecond)
	saveAfter := time.After(2800 * time.Millisecond)
	logChan := make(chan string)
	go saveBackups(snapshotTicker, saveAfter, logChan)
	actualLogs := []string{}
	for actualLog := range logChan {
		fmt.Println(actualLog)
		actualLogs = append(actualLogs, actualLog)
	}

	if !slices.Equal(expectedLogs, actualLogs) {
		fmt.Printf(`
---------------------------------
Test Failed:
expected:
%v
actual:
%v
`, sliceWithBullets(expectedLogs), sliceWithBullets(actualLogs))
	} else {
		fmt.Printf(`
---------------------------------
Test Passed:
expected:
%v
actual:
%v
`, sliceWithBullets(expectedLogs), sliceWithBullets(actualLogs))
	}
}

func sliceWithBullets[T any](slice []T) string {
	if slice == nil {
		return "  <nil>"
	}
	if len(slice) == 0 {
		return "  []"
	}
	output := ""
	for i, item := range slice {
		form := "  - %#v\n"
		if i == (len(slice) - 1) {
			form = "  - %#v"
		}
		output += fmt.Sprintf(form, item)
	}
	return output
}