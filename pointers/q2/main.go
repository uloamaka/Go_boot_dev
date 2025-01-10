package main

import (
	"strings"
)

func removeProfanity(message *string) {
	replacements  := map[string]string{
        "fubb":  "****",
        "shiz":  "****",
        "witch": "*****",
    }

    for word, replacement := range replacements {
        *message = strings.ReplaceAll(*message, word, replacement)
    }
}


func main() {
	runTestCases()
}
// algorithm
// ------------------------------------