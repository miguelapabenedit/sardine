package main

type Transactions []Transaction
type Transaction struct {
	ID            int `json:"id"`
	UserID        int `json:"user_id"`
	AmountUsCents int `json:"amount_us_cents"`
	CardID        int `json:"card_id"`
}

type User struct {
	ID               int
	TotalAmountSpend int
	Cards            map[int]interface{}
	Transactions     Transactions
}

func NewUser(id int) *User {
	return &User{
		ID:    id,
		Cards: make(map[int]interface{}),
	}
}

func (u *User) AddTransaction(txn Transaction) {
	if _, exists := u.Cards[txn.CardID]; !exists {
		u.Cards[txn.CardID] = nil
	}

	u.TotalAmountSpend += txn.AmountUsCents
	u.Transactions = append(u.Transactions, txn)
}

func (u *User) GetLastTransaction() Transaction {
	if len(u.Transactions) == 0 {
		return Transaction{}
	}
	return u.Transactions[len(u.Transactions)-1]
}
