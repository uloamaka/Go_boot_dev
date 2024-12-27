package main
/* Assignment
Create a SendMessage method for the User struct.

It should take a message string and messageLength int as inputs.

If the messageLength is within the user's message character limit, return the original message and true (indicating the message can be sent), otherwise, return an empty string and false. */
// ---------------------------------------------------------
/* package main

// ?

// don't touch below this line

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	membership := Membership{Type: membershipType}
	if membershipType == "premium" {
		membership.MessageCharLimit = 1000
	} else {
		membership.Type = "standard"
		membership.MessageCharLimit = 100
	}
	return User{Name: name, Membership: membership}
}*/
// ----------------solution---------------------------------


// ?
func (u User) SendMessage(message string, messageLength int) (string, bool) {
	if messageLength <= u.MessageCharLimit {
		return message, true
	} else {
		return "", false
	}
}
// don't touch below this line

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	membership := Membership{Type: membershipType}
	if membershipType == "premium" {
		membership.MessageCharLimit = 1000
	} else {
		membership.Type = "standard"
		membership.MessageCharLimit = 100
	}
	return User{Name: name, Membership: membership}
}

func main() {
	runTestCases()
}