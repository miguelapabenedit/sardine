package main

const (
	LowRisk    = "low"
	MediumRisk = "medium"
	HighRisk   = "high"
)

func EvaluateTransactionsRisk(txns Transactions) []string {
	users := make(map[int]*User)
	risks := make([]string, 0, len(txns))
	for _, txn := range txns {
		user, exists := users[txn.UserID]
		if !exists {
			user = NewUser(txn.UserID)
			users[txn.UserID] = user
		}

		user.AddTransaction(txn)
		risks = append(risks, RiskEvaluator(
			*user,
			SpendRisk,
			TotalSpendRisk,
			CardsRisk,
		))
	}
	return risks
}

func RiskEvaluator(u User, evals ...func(u User) string) string {
	riskResult := LowRisk
	for _, eval := range evals {
		evalRisk := eval(u)
		if evalRisk == LowRisk {
			continue
		}
		riskResult = evalRisk
		if riskResult == HighRisk {
			return riskResult
		}
	}
	return riskResult
}
