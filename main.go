package main

import "fmt"

func main() {
	// Load our environment variables
	var api AllyApi
	api.Initialize()

	fmt.Printf("%s\n", api.Accounts()[0].Accountholdings.Displaydata.Totalsecurities)
}
