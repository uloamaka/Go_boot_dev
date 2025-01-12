package main

import (
	"sync"
	"time"
)

func processMessages(messages []string) []string {
	resultChan := make(chan string, len(messages))
	var wg sync.WaitGroup

	for _, message := range messages {
		wg.Add(1)
		go func(msg string) {
			defer wg.Done()
			processed := process(msg)
			resultChan <- processed
		}(message)
	}

	wg.Wait()
	close(resultChan)

	var processedMessages []string
	for result := range resultChan {
		processedMessages = append(processedMessages, result)
	}

	return processedMessages
}

// don't touch below this line

func process(message string) string {
	time.Sleep(1 * time.Second)
	return message + "-processed"
}
func main() {
	runTestCases()
}

// algorithm
// ---------------------------------
// 1. Iterate through the message slice.
//   - for each message string call the process func with the goroutine

