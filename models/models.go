package models

type Account struct {
	ID      int64 `json:"id"`
	Balance int64 `json:"balance"`
}

type Deposit struct {
	Destination int64 `json:"destination"`
	Amount      int64 `json:"amount"`
}

type Withdrawal struct {
	Origin int64 `json:"origin"`
	Amount int64 `json:"amount"`
}

type Transfer struct {
	Origin      int64 `json:"origin"`
	Destination int64 `json:"destination"`
	Amount      int64 `json:"amount"`
}
