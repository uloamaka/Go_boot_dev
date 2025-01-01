package main

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (dm directMessage)importance () int {
	if dm.isUrgent == true {
		return 50
	} else {
		return dm.priorityLevel
	}
}

func (gm groupMessage)importance () int {
	return gm.priorityLevel
}

func (sa systemAlert)importance () int {
	return 100
}

func processNotification(n notification) (string, int) {
	switch v := n.(type) {
		case directMessage:
			return v.senderUsername, v.importance()
		case groupMessage:
			return v.groupName, v.importance()
		case systemAlert:
			return v.alertCode, v.importance()
		default:
			return "", 0
	}
}

func main() {
	runTestCases()
}