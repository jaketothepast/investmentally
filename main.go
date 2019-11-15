package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/joho/godotenv"
	"github.com/mrjones/oauth"
)

var endpoint string = "https://api.tradeking.com/v1/"

func main() {
	// Load our environment variables
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// Set up our new oauth consumer
	cons := oauth.NewConsumer(
		os.Getenv("CONSUMER_KEY"),
		os.Getenv("CONSUMER_SECRET"),
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://developers.tradeking.com/oauth/request_token",
			AuthorizeTokenUrl: "https://developers.tradeking.com/oauth/authorize",
			AccessTokenUrl:    "https://developers.tradeking.com/oauth/access_token",
		},
	)

	// Set up our HTTP client
	client, err := cons.MakeHttpClient(&oauth.AccessToken{Token: os.Getenv("ACCESS_TOKEN"), Secret: os.Getenv("ACCESS_SECRET")})
	if err != nil {
		panic(err)
	}

	resp, err := client.Get(fmt.Sprintf("%s\\%s", endpoint, "accounts"))

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
