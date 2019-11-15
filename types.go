package main

import "encoding/xml"

/**
Response structure from Ally
*/
type AllyResponse struct {
	XMLName     xml.Name         `xml:"response"`
	Error       string           `xml:"error"`
	ResponseId  string           `xml:"id,attr"`
	ElapsedTime int              `xml:"elapsedtime"`
	Accounts    []AccountSummary `xml:"accounts>accountsummary"`
}

/**
TODO: We need to work on AccountHoldings next.
*/
type AccountHoldings struct {
}

/**
Account summary information
*/
type AccountSummary struct {
	XMLName     xml.Name        `xml:"accountsummary"`
	Account     int             `xml:"account"`
	AccountName string          `xml:"accountname"`
	Balance     AccountBalance  `xml:"accountbalance"`
	Holdings    AccountHoldings `xml:"accountholdings"`
}

/**
Different securites currently held with an account
*/
type Securities struct {
	XMLName      xml.Name `xml:"securities"`
	LongOptions  float64  `xml:"longoptions"`
	LongStocks   float64  `xml:"longstocks"`
	Options      float64  `xml:"options"`
	ShortOptions float64  `xml:"shortoptions"`
	ShortStocks  float64  `xml:"shortstocks"`
}

/**
Various values representing account balance
*/
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

/**
Various values representing buying power of this account
*/
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

/**
Various values representing money in this account
*/
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
