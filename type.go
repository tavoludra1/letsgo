package main

import "math/rand"

type Account struct {
	ID int
	FistName string
	LastName string
	Number int64
	Balance int64
}

func NewAccount(firstName, lastName string) *Account {
	return &Account{
		ID rand.Intn(100000),
		FistName: firstName,
		LastName: lastName,
		Number: rand.Intn(10000000),
	}
}