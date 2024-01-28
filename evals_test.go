package main_test

import (
	"testing"

	main "sardine"

	"github.com/stretchr/testify/assert"
)

func TestSpendRisk(t *testing.T) {
	tt := []struct {
		Name  string
		Input main.User
		Exp   string
	}{{
		Name:  "when user with empty transactions, then return low",
		Input: main.User{},
		Exp:   main.HighRisk,
	}, {
		Name:  "when user with 6k spend transactions, then return Medium",
		Input: main.User{Transactions: main.Transactions{{AmountUsCents: 600000}}},
		Exp:   main.MediumRisk,
	}, {
		Name:  "when user with $10k+ spend transactions, then return High",
		Input: main.User{Transactions: main.Transactions{{AmountUsCents: 1000001}}},
		Exp:   main.HighRisk,
	}}

	for _, test := range tt {
		riskResult := main.SpendRisk(test.Input)
		assert.Equal(t, test.Exp, riskResult)
	}
}

func TestTotalSpendRisk(t *testing.T) {
	tt := []struct {
		Name  string
		Input main.User
		Exp   string
	}{{
		Name:  "when user with 1 dolar total transactions, then return low",
		Input: main.User{TotalAmountSpend: 1},
		Exp:   main.LowRisk,
	}, {
		Name:  "when user with 10k+ total spend transactions, then return Medium",
		Input: main.User{TotalAmountSpend: 1000001},
		Exp:   main.MediumRisk,
	}, {
		Name:  "when user with 20k+ total spend transactions, then return High",
		Input: main.User{TotalAmountSpend: 2000001},
		Exp:   main.HighRisk,
	}}

	for _, test := range tt {
		riskResult := main.TotalSpendRisk(test.Input)
		assert.Equal(t, test.Exp, riskResult)
	}
}

func TestCardsRisk(t *testing.T) {
	cardsFunc := func(numbOFcards int) map[int]interface{} {
		cards := make(map[int]interface{})
		for i := 0; i < numbOFcards; i++ {
			cards[i] = nil
		}
		return cards
	}

	tt := []struct {
		Name  string
		Input main.User
		Exp   string
	}{{
		Name:  "when user with 1 Card, then return low",
		Input: main.User{Cards: cardsFunc(1)},
		Exp:   main.LowRisk,
	}, {
		Name:  "when user with 2 Card, then return Medium",
		Input: main.User{Cards: cardsFunc(2)},
		Exp:   main.MediumRisk,
	}, {
		Name:  "when user with 3 Card, then return High",
		Input: main.User{Cards: cardsFunc(3)},
		Exp:   main.HighRisk,
	}}

	for _, test := range tt {
		riskResult := main.CardsRisk(test.Input)
		assert.Equal(t, test.Exp, riskResult)
	}
}
