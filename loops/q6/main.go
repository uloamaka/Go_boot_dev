package main

// import "fmt"

func countConnections(groupSize int) int {
	connections := 0
	for i := 0; i < groupSize; i++ {
		connections += i
	}
	return connections
}

func main() {
	runTestCases()
}
