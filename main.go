package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mrjones/oauth"
)

var endpoint string = "https://api.tradeking.com/v1/"

/* My configuration structure for oauth */
type config struct {
	ConsumerKey    string
	ConsumerSecret string
	AccessToken    string
	AccessSecret   string
}

func main() {
	c := config{
		ConsumerKey:    os.Getenv("CONSUMER_KEY"),
		ConsumerSecret: os.Getenv("CONSUMER_SECRET"),
		AccessToken:    os.Getenv("ACCESS_TOKEN"),
		AccessSecret:   os.Getenv("ACCESS_SECRET"),
	}

	// Set up our new oauth consumer
	cons := oauth.NewConsumer(
		c.ConsumerKey,
		c.ConsumerSecret,
		oauth.ServiceProvider{
			RequestTokenUrl:   "https://developers.tradeking.com/oauth/request_token",
			AuthorizeTokenUrl: "https://developers.tradeking.com/oauth/authorize",
			AccessTokenUrl:    "https://developers.tradeking.com/oauth/access_token",
		},
	)

	// Set up our HTTP client
	client, err := cons.MakeHttpClient(&oauth.AccessToken{Token: c.AccessToken, Secret: c.AccessSecret})
	if err != nil {
		panic(err)
	}

	resp, err := client.Get(fmt.Sprintf("%s\\%s", endpoint, "accounts"))

	defer resp.Body.Close()
	b, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("Got %s from Ally\n", b)
}
