package data

import (
	"github.com/rogerpoliver/ebanx-api/models"
)

var accounts []models.Account

func Reset() {
	accounts = accounts[:0]
}

func Deposit(destination int64, amount int64) {
	for _, account := range accounts {
		if account.ID == destination {
			account.Balance = account.Balance + amount
		}
	}
}

func Withdraw(origin int64, amount int64) {
	for _, account := range accounts {
		if account.ID == origin {
			account.Balance = account.Balance - amount
		}
	}
}

func Transfer(origin int64, destination int64, amount int64) {
	Withdraw(origin, amount)
	Deposit(destination, amount)
}
