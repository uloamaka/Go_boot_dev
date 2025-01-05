package main



func createMatrix(rows, cols int) [][]int {
	xy := make([][]int, rows) 
	for i := 0; i < rows; i++ {
		y := make([]int, cols) 
		for j := 0; j < cols; j++ {
			y[j] = i * j 
		}
		xy[i] = y 
	}
	return xy
}

func main() {
	runTestCases()
}
