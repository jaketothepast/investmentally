package main

import "encoding/xml"

/**
Our Response structure from ALLY
*/
type AllyResponse struct {
	XMLName     xml.Name         `xml:"response"`
	Error       string           `xml:"error"`
	ResponseId  string           `xml:"id,attr"`
	ElapsedTime int              `xml:"elapsedtime"`
	Accounts    []AccountSummary `xml:"accounts>accountsummary"`
}

type AccountHoldings struct {
}

type AccountSummary struct {
	XMLName     xml.Name `xml:"accountsummary"`
	Account     int      `xml:"account"`
	AccountName string   `xml:"accountname"`
	Balance     AccountBalance
	Holdings    AccountHoldings
}

type Securities struct {
}

type AccountBalance struct {
	Account      int `xml:"accounts>accountsummary>accountbalance>account"`
	AccountValue float64
	BuyingPower  BuyingPower
	FedCall      int
	HouseCall    int
	Money        Money
	Securities   Securities
}

type BuyingPower struct {
	CashAvailableForWithdrawal float64
	DayTrading                 int
	EquityPercentage           int
	Options                    int
	SodDayTrading              int
	SodOptions                 int
	SodStock                   int
	Stock                      int
}

type Money struct {
	AccruedInterest   float64
	Cash              float64
	CashAvailable     float64
	MarginBalance     int
	Mmf               int
	Total             int
	UnclearedDeposits float64
	UnsettledFunds    float64
	Yield             float64
}
