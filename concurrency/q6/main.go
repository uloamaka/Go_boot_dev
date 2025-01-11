package main


func concurrentFib(n int) []int {
	ch := make(chan int)
	num := []int{}
	go fibonacci(n, ch)
	for val := range ch {
		num = append(num, val)
	}
	return num
}

// don't touch below this line

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {
	runTestCases()
}

// algorithm
// ---------------------------------
// 1. Create a new channel of ints
// 2. Call fibonacci concurrently
// 3. Use a range loop to read from the channel and append the values to a slice
// 4. Return the slice
