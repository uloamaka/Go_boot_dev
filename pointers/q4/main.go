package main

import (
	"strings"
)

// func removeProfanity(message *string) {
// 	replacements  := map[string]string{
//         "fubb":  "****",
//         "shiz":  "****",
//         "witch": "*****",
//     }
//     if message == nil {
//         return
//     }

//     for word, replacement := range replacements {
//         *message = strings.ReplaceAll(*message, word, replacement)
//     }
// }
func removeProfanity(message *string) {
    if message == nil {
        return
    }
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "fubb", "****")
	messageVal = strings.ReplaceAll(messageVal, "shiz", "****")
	messageVal = strings.ReplaceAll(messageVal, "witch", "*****")
	*message = messageVal
}



func main() {
	runTestCases()
}
// algorithm
// ------------------------------------