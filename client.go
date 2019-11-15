package main

import (
	"encoding/xml"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/mrjones/oauth"
	"io/ioutil"
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
	return
}

/* The /accounts endpoint of Ally */
func (c *AllyApi) Accounts() (resp AllyResponse, err error) {
	b, err := c.get("accounts")
	if err != nil {
		return
	}

	defer b.Body.Close()
	raw, err := ioutil.ReadAll(b.Body)
	if err != nil {
		return
	}

	err = xml.Unmarshal(raw, &resp)
	if err != nil {
		return
	}

	return
}
