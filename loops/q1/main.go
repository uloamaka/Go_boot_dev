package main


func bulkSend(numMessages int) float64 {
	var totalCost float64
	for i := 0; i < numMessages; i++ {
		fee := float64(i) * 0.01
		totalCost += 1.0 + fee
	}
	return totalCost
}

func main() {
	runTestCases()
}