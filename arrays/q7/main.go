package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, Word := range msg {
		for _, badWord := range badWords {
			if Word == badWord {
				return i
			}
		}
	}
	return -1
}

func main() {
	runTestCases()
}
