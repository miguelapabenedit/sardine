package main_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	main "sardine"
	"testing"

	"github.com/stretchr/testify/assert"
)

const path = "/transactions/risk-evaluation"

type mockReader int

func (mra mockReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("mocked_reader_err")
}

func TestHandler_WhenGetMethod_ThenBadRequest(t *testing.T) {
	var (
		expErr = "method 'GET', not allowed\n"
		rr     = httptest.NewRecorder()
		req    = httptest.NewRequest(http.MethodGet, path, nil)
	)

	main.HandleTransactionRiskEvaluation(rr, req)

	assert.Equal(t, http.StatusMethodNotAllowed, rr.Code)
	assert.Equal(t, expErr, rr.Body.String())
}

func TestHandler_WhenReadAllError_ThenBadRequest(t *testing.T) {
	var (
		expErr = "mocked_reader_err\n"
		rr     = httptest.NewRecorder()
		req    = httptest.NewRequest(http.MethodPost, path, mockReader(1))
	)

	main.HandleTransactionRiskEvaluation(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, expErr, rr.Body.String())
}

func TestHandler_WhenValidationErr_ThenBadRequest(t *testing.T) {
	var (
		expErr = "empty transactions not allowed\n"
		rr     = httptest.NewRecorder()

		reqBody = main.TransactionRiskEvaluationRequest{}
	)

	bContent, _ := json.Marshal(reqBody)
	buff := bytes.NewBuffer(bContent)

	main.HandleTransactionRiskEvaluation(rr, httptest.NewRequest(http.MethodPost, path, buff))

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	assert.Equal(t, expErr, rr.Body.String())
}

func TestHandler_WhenAllOk_ThenSuccess(t *testing.T) {
	var (
		expected = main.TransactionRiskEvaluationResponse{
			RiskRating: []string{
				"low",
				"medium",
				"high",
				"low",
				"medium",
				"high",
			},
		}
		rr      = httptest.NewRecorder()
		reqBody = main.TransactionRiskEvaluationRequest{
			Transactions: []main.Transaction{
				{ID: 1, UserID: 1, AmountUsCents: 200000, CardID: 1},
				{ID: 2, UserID: 1, AmountUsCents: 600000, CardID: 1},
				{ID: 3, UserID: 1, AmountUsCents: 1100000, CardID: 1},
				{ID: 4, UserID: 2, AmountUsCents: 100000, CardID: 2},
				{ID: 5, UserID: 2, AmountUsCents: 100000, CardID: 3},
				{ID: 6, UserID: 2, AmountUsCents: 100000, CardID: 4},
			}}
		bContent, _ = json.Marshal(reqBody)
		buff        = bytes.NewBuffer(bContent)
	)

	main.HandleTransactionRiskEvaluation(rr, httptest.NewRequest(http.MethodPost, path, buff))

	var result main.TransactionRiskEvaluationResponse
	if err := json.Unmarshal(rr.Body.Bytes(), &result); err != nil {
		t.Log(err.Error())
	}
	assert.Equal(t, expected, result)
	assert.Equal(t, http.StatusOK, rr.Code)
}
