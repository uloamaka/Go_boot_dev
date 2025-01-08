package main

import "errors"

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	elem, ok := users[name]
	if !ok {
		return false, errors.New("not found")
	}
	if !elem.scheduledForDeletion {
		return false, nil
	}
	if elem.scheduledForDeletion {
		delete(users, name)
	}
	return true, nil
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func main() {
	runTestCases()
}