package main

func countReports(numSentCh chan int) int {
	totalReports := 0
	for i := 0;  ; i++ {
		v, ok := <-numSentCh
		if !ok {
			break
		}
		totalReports += v
	} 
	return totalReports 
}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}

func main() {
	runTestCases()
}
// algorithm
// ---------------------------------
// 1. Use an infinite for loop to read from the channel:
// 2. If the channel is closed, break out of the loop
// 3. Otherwise, keep a running total of the number of reports sent
// 4. Return the total number of reports sent