package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	riskPath = "/transactions/risk-evaluations"
	port     = ":8080"
)

// End of Test  time Comments
// I could get more tests coverage and refactor the file structure a little bit more, I would be glad to extend more in an interview session
// this is not the normal way I build the base structure (but... is subjective to the workspace and other factos)
// Glad to show more projects in a call or exten my knowledge about how I would struct projects.
// Thanks for watching and sorry for all the mistakes and brain melting in the way
func main() {
	fmt.Println("starting server at localhost", port)
	http.HandleFunc(riskPath, HandleTransactionRiskEvaluation)

	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatal(err)
	}
	fmt.Println("server closed")
}
