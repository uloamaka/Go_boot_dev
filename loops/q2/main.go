package main

func maxMessages(thresh int) int {
	total := 0
	count := 0
	for  {
		fee := count
		cost := 100 + fee
		if total + cost > thresh {
			break
		}
		total += cost
		count ++
	}
	return count
}

func main() {
	runTestCases()
}