package main
// Assignment
// Our over-engineering boss is at it again. He's heard about memory layout and wants to squeeze every last byte out of our structs.

// Run the tests to see the current size of the structs, then update the struct definitions to minimize memory usage
// ---------------------------------------------------------
/* package main

type contact struct {
	sendingLimit int32
	userID       string
	age          int32
}

type perms struct {
	canSend         bool
	canReceive      bool
	permissionLevel int
	canManage       bool
} */
// ----------------solution---------------------------------

type contact struct {
	sendingLimit int32
	age          int32
	userID       string
}

type perms struct {
	permissionLevel int
	canSend         bool
	canReceive      bool
	canManage       bool
}


func main() {
	runTestCases()
}