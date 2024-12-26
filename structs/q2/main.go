package main
// Assignment
// Textio has a bug, we've been sending texts that are missing critical bits of information! Before we send text messages in Textio, we must check to make sure the required fields have non-zero values.
// -----------------------------------------------------------------
// package main

// type messageToSend struct {
// 	message   string
// 	sender    user
// 	recipient user
// }

// type user struct {
// 	name   string
// 	number int
// }

// func canSendMessage(mToSend messageToSend) bool {
// 	// ?
// 	return true
// }
// ----------------------------------------------------------------------
// Notice that the user struct is a nested struct within the messageToSend struct. Both sender and recipient are user struct types.

// Complete the canSendMessage function. It should return true only if the sender and recipient fields each contain a name and a number. If any of the default zero values are present, return false instead.

// ----------------solution-----------------------------------

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	return mToSend.sender.number != 0 && mToSend.recipient.number != 0 && mToSend.sender.name != "" && mToSend.recipient.name != ""
}

func main() {
	runTestCases()
}
