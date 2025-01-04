package main


func getMessageCosts(messages []string) []float64 {
	msgLength := len(messages)
	costs := make([]float64, msgLength)
	
	for i := 0; i < msgLength; i++{
		costs[i] = float64(len(messages[i])) * 0.01
	}
	return costs
}


func main() {
	runTestCases()
}