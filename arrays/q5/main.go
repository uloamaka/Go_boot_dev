package main


type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	costsByDay := make([]float64, 0)
	for _, c := range costs {
		for len(costsByDay) <= c.day {
			costsByDay = append(costsByDay, 0.0)
		}
		costsByDay[c.day] += c.value
	}
	return costsByDay
}

func main() {
	runTestCases()
}