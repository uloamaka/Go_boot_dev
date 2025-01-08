package main

func getCounts(messagedUsers []string, validUsers map[string]int) {
	for _, username := range messagedUsers { 
		if _, exists := validUsers[username]; exists {
			validUsers[username]++
		}
	}
}


func main() {
	runTestCases()
}

// algorithm
// ------------------------------------
	// 1. Loop through each username in the messagedUsers slice using a for range loop.
	
	// 2. For each username, check if it exists in the validUsers map.
	//    - Use the syntax `value, exists := map[key]` to check existence.
	
	// 3. If the username is a valid user (exists in the map):
	//    - Increment the corresponding value in the map by 1.

	// 4. If the username does not exist in validUsers:
	//    - Do nothing, as only valid users should be updated.

	// The validUsers map will be updated in place since maps are mutable.