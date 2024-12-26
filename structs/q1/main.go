package main
// Assignment
// Complete the definition of the messageToSend struct. It needs two fields:
// ------------------------------------------------------------
// type messageToSend struct { }
// ------------------------------------------------------------
// phoneNumber - an integer
// message - a string.
// ----------------solution-----------------------------------
import "fmt"

func main() {
	runTest()
}

func getMessageText(m messageToSend) string {
	return fmt.Sprintf("Sending message: '%s' to: %v", m.message, m.phoneNumber)
}


type messageToSend struct {
	phoneNumber int
	message     string
}
