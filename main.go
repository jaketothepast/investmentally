package main

import (
	"fmt"
)

func main() {
	// Load our environment variables
	var api AllyApi
	api.Initialize()

	acctSummary, _ := api.Accounts()

	fmt.Printf("%f\n", acctSummary.Accounts[0].AccountHoldingsInfo.TotalSecurities)
}
