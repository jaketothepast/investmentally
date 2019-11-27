package main

import (
	"fmt"
)

func main() {
	// Load our environment variables
	var api AllyApi
	api.Initialize()

	accountId := api.Accounts()[0].Account
	fmt.Printf("AccountDetail: %s\n", api.AccountHistory(accountId).Transactions[0].Detail.SettlementDate)
}
