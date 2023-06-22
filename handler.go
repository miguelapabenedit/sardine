package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func HandleTransactionRiskEvaluation(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, fmt.Sprintf("method '%v', not allowed", req.Method), http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var txnRequest TransactionRiskEvaluationRequest
	if err := json.Unmarshal(body, &txnRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := txnRequest.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var response TransactionRiskEvaluationResponse
	response.RiskRating = EvaluateTransactionsRisk(txnRequest.Transactions)

	rContent, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(rContent)
}

type TransactionRiskEvaluationRequest struct {
	Transactions Transactions `json:"transactions"`
}
type TransactionRiskEvaluationResponse struct {
	RiskRating []string `json:"risk_ratings"`
}

func (r *TransactionRiskEvaluationRequest) Validate() error {
	if len(r.Transactions) == 0 {
		return errors.New("empty transactions not allowed")
	}
	txnIDs := make(map[int]interface{})
	for _, txn := range r.Transactions {
		if _, exist := txnIDs[txn.ID]; exist {
			return errors.New("duplicated transactions ids not allowed")
		}
		txnIDs[txn.ID] = nil
		// more elegant check can be performend, leave it for later
		if txn.AmountUsCents < 1 || txn.CardID < 1 || txn.UserID < 1 || txn.ID < 1 {
			return errors.New("invalid transaction input, negative values not allowed")
		}
	}
	return nil
}
