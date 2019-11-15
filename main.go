package main

import "fmt"

func main() {
	// Load our environment variables
	var api AllyApi
	api.Initialize()

	fmt.Printf("%f\n", api.Accounts()[0].Accountbalance.Money.Cash)
}
