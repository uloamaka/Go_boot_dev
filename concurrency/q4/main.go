package main

func addEmailsToQueue(emails []string) chan string {
	emailQueue := make(chan string, len(emails))
	for _, email := range emails {
		emailQueue <- email
	}
	return emailQueue
}

func main() {
	runTestCases()
}
// algorithm
// ---------------------------------
// 1. get the lenght of the emails slice
// 2. create a channel of strings
// 3. iterate over the emails slice
// 4. send each email to the channel
// 5. return the channel