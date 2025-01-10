package main

type Analytics struct {
	MessagesTotal     int
	MessagesFailed    int
	MessagesSucceeded int
}

type Message struct {
	Recipient string
	Success   bool
}

// don't touch above this line

func getMessageText(analytic *Analytics, message Message) {
    if message.Success {
        analytic.MessagesSucceeded++
    } else {
        analytic.MessagesFailed++
    }
    analytic.MessagesTotal++
}


func main() {
	runTestCases()
}
// algorithm
// ------------------------------------
// checking for success message
// 1. If success == true 
//   - increment analytic messageSucceed count
// 2. else 
//   - increament analytic messageFailed count