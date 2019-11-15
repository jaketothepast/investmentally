package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

func main() {
	// Load our environment variables
	var api AllyApi
	api.Initialize()

	// Set up our HTTP client

	resp, err := api.Get("accounts")

	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)

	var acctSummary AllyResponse
	err = xml.Unmarshal(b, &acctSummary)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", b)
	fmt.Printf("%f\n", acctSummary.Accounts[0].AccountHoldingsInfo.TotalSecurities)
}
