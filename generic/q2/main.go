package main

import (
	"fmt"
	"time"
)

func chargeForLineItem[T lineItem](newItem T, oldItems []T, balance float64) ([]T, float64, error) {
	if balance < newItem.GetCost() {
		return nil, 0, fmt.Errorf("insufficient funds") 
	}
	oldItems = append(oldItems, newItem)
	newBalance := balance - newItem.GetCost() 
	return oldItems, newBalance, nil
}

// don't edit below this line

type lineItem interface {
	GetCost() float64
	GetName() string
}

type subscription struct {
	userEmail string
	startDate time.Time
	interval  string
}

func (s subscription) GetName() string {
	return fmt.Sprintf("%s subscription", s.interval)
}

func (s subscription) GetCost() float64 {
	if s.interval == "monthly" {
		return 25.00
	}
	if s.interval == "yearly" {
		return 250.00
	}
	return 0.0
}

type oneTimeUsagePlan struct {
	userEmail        string
	numEmailsAllowed int
}

func (otup oneTimeUsagePlan) GetName() string {
	return fmt.Sprintf("one time usage plan with %v emails", otup.numEmailsAllowed)
}

func (otup oneTimeUsagePlan) GetCost() float64 {
	const costPerEmail = 0.03
	return float64(otup.numEmailsAllowed) * costPerEmail
}

// algorithm
// initialise the zero values for the variable Type
// Check the customer balance 
//  - if it is enough to pay for the newItem
//  - if it is not return "insufficient funds" error and zero values for the other return values.
//  - if it is enough :
//  - Add the line item to the user's history by appending the newItem to the slice of oldItems. This new slice is your first return value.
//  - Calculate the user's new balance by subtracting the cost of the new item from their balance. This is your second return value.