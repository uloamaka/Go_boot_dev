package main

import (
	"errors"
)

func divide(x, y float64) (float64, error) {
	if y == 0 {
		var err error = errors.New("no dividing by 0")
		return 0, err
	}
	return x / y, nil
}

func main() {	
	runTestCases()
}
