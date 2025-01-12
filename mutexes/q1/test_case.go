package main

import (
	"fmt"
	"sync"
)

var withSubmit = true

func runTestCases() {
	type testCase struct {
		email string
		count int
	}
	var tests = []testCase{
		{"norman@bates.com", 23},
		{"marion@bates.com", 67},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{"lila@bates.com", 31},
			{"sam@bates.com", 453},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		sc := safeCounter{
			counts: make(map[string]int),
			mu:     &sync.Mutex{},
		}
		var wg sync.WaitGroup
		for i := 0; i < test.count; i++ {
			wg.Add(1)
			go func(email string) {
				sc.inc(email)
				wg.Done()
			}(test.email)
		}
		wg.Wait()

		if output := sc.val(test.email); output != test.count {
			failCount++
			fmt.Printf(`
---------------------------------
Test Failed:
  email: %v
  count: %v
  expected count: %v
  actual count:   %v
`, test.email, test.count, test.count, output)
		} else {
			passCount++
			fmt.Printf(`
---------------------------------
Test Passed:
  email: %v
  count: %v
  expected count: %v
  actual count:   %v
`, test.email, test.count, test.count, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}