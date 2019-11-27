package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Load our environment variables
	var api AllyApi
	api.Initialize()

	acctId, _ := strconv.Atoi(api.Accounts()[0].Account)

	fmt.Printf("AccountDetail: %s\n", api.AccountDetail(acctId).AccountHoldings.Holding[0].Displaydata.Change)
}
