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
	XMLName     xml.Name        `xml:"accountsummary"`
	Account     int             `xml:"account"`
	AccountName string          `xml:"accountname"`
	Balance     AccountBalance  `xml:"accountbalance"`
	Holdings    AccountHoldings `xml:"accountholdings"`
}

type Securities struct {
}

type AccountBalance struct {
	XMLName      xml.Name    `xml:"accountbalance"`
	Account      int         `xml:"account"`
	AccountValue float64     `xml:"accountvalue"`
	BuyingPower  BuyingPower `xml:"buyingpower"`
	FedCall      int         `xml:"fedcall"`
	HouseCall    int         `xml:"housecall"`
	Money        Money       `xml:"money"`
	Securities   Securities  `xml:"securities"`
}

type BuyingPower struct {
	XMLName                    xml.Name `xml:"buyingpower"`
	CashAvailableForWithdrawal float64  `xml:"cashavailableforwithdrawal"`
	DayTrading                 int      `xml:"daytrading"`
	EquityPercentage           int      `xml:"equitypercentage"`
	Options                    int      `xml:"options"`
	SodDayTrading              int      `xml:"soddaytrading"`
	SodOptions                 int      `xml:"sodoptions"`
	SodStock                   int      `xml:"sodstock"`
	Stock                      int      `xml:"stock"`
}

type Money struct {
	XMLName           xml.Name `xml:"money"`
	AccruedInterest   float64  `xml:"accruedinterest"`
	Cash              float64  `xml:"cash"`
	CashAvailable     float64  `xml:"cashavailable"`
	MarginBalance     float64  `xml:"marginbalance"`
	Mmf               float64  `xml:"mnf"`
	Total             float64  `xml:"total"`
	UnclearedDeposits float64  `xml:"uncleareddeposits"`
	UnsettledFunds    float64  `xml:"unsettledfunds"`
	Yield             float64  `xml:"yield"`
}
