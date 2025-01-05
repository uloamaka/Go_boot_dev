package main

func sum(nums ...int) int {
	final := 0
	for i := 0; i < len(nums); i++ {
		final += nums[i]
	}
	return final
}

func main() {
	runTestCases()
}