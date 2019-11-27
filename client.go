package main

import (
	"encoding/xml"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mrjones/oauth"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

/* The trading endpoint for Ally */
const endpoint string = "https://api.tradeking.com/v1/"

type AllyApi struct {
	consumer *oauth.Consumer
	Client   *http.Client
}

func (c *AllyApi) Initialize() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// Set up our new oauth consumer
	c.consumer = oauth.NewConsumer(
		os.Getenv("CONSUMER_KEY"),
		os.Getenv("CONSUMER_SECRET"),
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://developers.tradeking.com/oauth/request_token",
			AuthorizeTokenUrl: "https://developers.tradeking.com/oauth/authorize",
			AccessTokenUrl:    "https://developers.tradeking.com/oauth/access_token",
		},
	)

	c.Client, err = c.consumer.MakeHttpClient(
		&oauth.AccessToken{Token: os.Getenv("ACCESS_TOKEN"), Secret: os.Getenv("ACCESS_SECRET")})

	if err != nil {
		panic(err)
	}
}

func (c *AllyApi) get(path string) (resp *http.Response, err error) {
	resp, err = c.Client.Get(fmt.Sprintf("%s\\%s", endpoint, path))
	return resp, err
}

//GetAndMarshal call GET on the path, and marshal the resopnse to interface
// type
func (c *AllyApi) getAndRead(path string) []byte {
	resp, err := c.get(path)
	if err != nil {
		log.Fatal("Could not make request")
	}

	defer resp.Body.Close()
	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Could not read body!")
		return nil
	}
	return raw
}

/* The /accounts endpoint of Ally */
func (c *AllyApi) Accounts() []AccountSummary {
	var resp AccountResponse
	_ = xml.Unmarshal(c.getAndRead("accounts"), &resp)
	return resp.Accounts.Accountsummary
}

func (c *AllyApi) AccountBalances() (balances []AccountBalance) {
	var resp AccountBalanceResponse
	_ = xml.Unmarshal(c.getAndRead("accounts/balances"), &resp)
	return resp.AccountBalances
}

func (c *AllyApi) AccountDetail(accountId string) AccountDetailResponse {
	var resp AccountDetailResponse
	_ = xml.Unmarshal(c.getAndRead(fmt.Sprintf("accounts/%s", accountId)), &resp)
	return resp
}

func (c *AllyApi) AccountBalance(accountId string) AccountDetailBalanceResponse {
	var resp AccountDetailBalanceResponse
	_ = xml.Unmarshal(c.getAndRead(fmt.Sprintf("accounts/%s/balances", accountId)), &resp)
	return resp
}

func (c *AllyApi) AccountHoldings(accountId string) AccountDetailHoldingsResponse {
	var resp AccountDetailHoldingsResponse
	_ = xml.Unmarshal(c.getAndRead(fmt.Sprintf("accounts/%s/holdings", accountId)), &resp)
	return resp
}
