package main

func SpendRisk(u User) string {
	txn := u.GetLastTransaction()
	if txn.AmountUsCents > dolarToCents(10000) {
		return HighRisk
	}
	if txn.AmountUsCents > dolarToCents(5000) {
		return MediumRisk
	}
	return LowRisk
}

func TotalSpendRisk(u User) string {
	if u.TotalAmountSpend > dolarToCents(20000) {
		return HighRisk
	}
	if u.TotalAmountSpend > dolarToCents(10000) {
		return MediumRisk
	}
	return LowRisk
}

func CardsRisk(u User) string {
	if len(u.Cards) > 2 {
		return HighRisk
	}
	if len(u.Cards) > 1 {
		return MediumRisk
	}
	return LowRisk
}

func dolarToCents(amount int) int {
	return amount * 100
}
