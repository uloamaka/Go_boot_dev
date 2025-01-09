package main

import "strings"

func countDistinctWords(messages []string) int {
	uniqueWords  := make(map[string]bool)
	for _, message := range messages {
		words := strings.Fields(message)
		for _, word := range words {
			word = strings.ToLower(word)
			if _, exists := uniqueWords[word]; !exists { 
				uniqueWords[word] = true
			}
		}
	}
	return len(uniqueWords)
}

func main() {
	runTestCases()
}

// algorithm
// ------------------------------------
// 1. Iterate over each messages in the input slice:
// 	- For each message:
// 	  - Extract each word using the strings.Fields.
// 	  - For each word convert to lower case
//  - Check if the the word exists in the msg map(uniqueWords):
//    - If not, store the new word in the msg map(uniqueWords)
//    - then get the len of the msg map(uniqueWords).

// 2. Return the count of the msg map (uniqueWords):
//    - After processing all messages, return count
