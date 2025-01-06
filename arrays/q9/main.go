package main

import "unicode"

func isValidPassword(password string) bool {
	if len(password) < 5 || len(password) > 12 {
		return false
	}

	hasUppercase := false
	hasDigit := false

	for _, char := range password {

		if unicode.IsUpper(char) {
			hasUppercase = true
		}
		if unicode.IsDigit(char) {
			hasDigit = true
		}
		if hasUppercase && hasDigit {
			break
		}
	}
	return hasUppercase && hasDigit
}

func main() {
	runTestCases()
}