package main

import "unicode/utf8"

func getNameCounts(names []string) map[rune]map[string]int {
	result := make(map[rune]map[string]int)
	for _, name := range names {
		if len(name) == 0 {
			continue
		}
		// firstChar := rune(name[0])
		firstChar, _ := utf8.DecodeRuneInString(name)
		if _, exists := result[firstChar]; !exists {
			result[firstChar] = make(map[string]int)
		}
		result[firstChar][name]++
	}
	return result
}

func main() {
	runTestCases()
}

// algorithm
// ------------------------------------
// 1. Initialize an empty nested map:
// 		- result := make(map[rune]map[string]int).

// 2. Iterate over each name in the input slice:
// 	- For each name:
// 	  - Handle empty names: Skip if len(name) == 0.
// 	  - Extract the first character (rune(name[0])).
//  - Check if the first character exists in result:
//    - If not, create a new inner map for that character.
//    - Increment the count for the name in the inner map.

// 3. Return the nested map:
//    - After processing all names, return result
