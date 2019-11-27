package main

import (
	"encoding/xml"
	"time"
)

type Accountbalance struct {
	XMLName      xml.Name `xml:"accountbalance"`
	Account      string   `xml:"account"`
	Accountvalue float64  `xml:"accountvalue"`
	Buyingpower  struct {
		Text                       string  `xml:",chardata"`
		Cashavailableforwithdrawal float64 `xml:"cashavailableforwithdrawal"`
		Daytrading                 float64 `xml:"daytrading"`
		Equitypercentage           float64 `xml:"equitypercentage"`
		Options                    float64 `xml:"options"`
		Soddaytrading              float64 `xml:"soddaytrading"`
		Sodoptions                 float64 `xml:"sodoptions"`
		Sodstock                   float64 `xml:"sodstock"`
		Stock                      float64 `xml:"stock"`
	} `xml:"buyingpower"`
	Fedcall   float64 `xml:"fedcall"`
	Housecall float64 `xml:"housecall"`
	Money     struct {
		Text              string  `xml:",chardata"`
		Accruedinterest   float64 `xml:"accruedinterest"`
		Cash              float64 `xml:"cash"`
		Cashavailable     float64 `xml:"cashavailable"`
		Marginbalance     float64 `xml:"marginbalance"`
		Mmf               float64 `xml:"mmf"`
		Total             float64 `xml:"total"`
		Uncleareddeposits float64 `xml:"uncleareddeposits"`
		Unsettledfunds    float64 `xml:"unsettledfunds"`
		Yield             float64 `xml:"yield"`
	} `xml:"money"`
	Securities struct {
		Text         string  `xml:",chardata"`
		Longoptions  float64 `xml:"longoptions"`
		Longstocks   float64 `xml:"longstocks"`
		Options      float64 `xml:"options"`
		Shortoptions float64 `xml:"shortoptions"`
		Shortstocks  float64 `xml:"shortstocks"`
		Stocks       float64 `xml:"stocks"`
		Total        float64 `xml:"total"`
	} `xml:"securities"`
}

type Accountholdings struct {
	XMLName     xml.Name `xml:"accountholdings"`
	Displaydata struct {
		Text            string `xml:",chardata"`
		Totalsecurities string `xml:"totalsecurities"`
	} `xml:"displaydata"`
	Holding []struct {
		Text        string `xml:",chardata"`
		Accounttype string `xml:"accounttype"`
		Costbasis   string `xml:"costbasis"`
		Displaydata struct {
			Text              string `xml:",chardata"`
			Accounttype       string `xml:"accounttype"`
			Assetclass        string `xml:"assetclass"`
			Change            string `xml:"change"`
			Costbasis         string `xml:"costbasis"`
			Desc              string `xml:"desc"`
			Lastprice         string `xml:"lastprice"`
			Marketvalue       string `xml:"marketvalue"`
			Marketvaluechange string `xml:"marketvaluechange"`
			Qty               string `xml:"qty"`
			Symbol            string `xml:"symbol"`
		} `xml:"displaydata"`
		Gainloss   float64 `xml:"gainloss"`
		Instrument struct {
			Text   string `xml:",chardata"`
			Cusip  string `xml:"cusip"`
			Desc   string `xml:"desc"`
			Factor string `xml:"factor"`
			Sectyp string `xml:"sectyp"`
			Sym    string `xml:"sym"`
		} `xml:"instrument"`
		Marketvalue       float64 `xml:"marketvalue"`
		Marketvaluechange float64 `xml:"marketvaluechange"`
		Price             float64 `xml:"price"`
		Purchaseprice     float64 `xml:"purchaseprice"`
		Qty               float64 `xml:"qty"`
		Quote             struct {
			Text      string  `xml:",chardata"`
			Change    float64 `xml:"change"`
			Lastprice float64 `xml:"lastprice"`
		} `xml:"quote"`
		Underlying string `xml:"underlying"`
	} `xml:"holding"`
	Totalsecurities float64 `xml:"totalsecurities"`
}

type AccountSummary struct {
	XMLName         xml.Name        `xml:"accountsummary"`
	Text            string          `xml:",chardata"`
	Account         string          `xml:"account"`
	Accountbalance  Accountbalance  `xml:"accountbalance"`
	Accountholdings Accountholdings `xml:"accountholdings"`
}

type AccountResponse struct {
	XMLName  xml.Name `xml:"response"`
	Text     string   `xml:",chardata"`
	ID       string   `xml:"id,attr"`
	Accounts struct {
		Text           string           `xml:",chardata"`
		Accountsummary []AccountSummary `xml:"accountsummary"`
	} `xml:"accounts"`
}

type AccountBalance struct {
	XMLName      xml.Name `xml:"accountbalance"`
	Account      int      `xml:"account"`
	AccountName  string   `xml:"accountname"`
	AccountValue float64  `xml:"accountvalue"`
}

type AccountBalanceResponse struct {
	XMLName         xml.Name         `xml:"response"`
	AccountBalances []AccountBalance `xml:"accountbalance"`
	TotalBalance    float64          `xml:"totalbalance>accountvalue"`
}

type AccountDetailResponse struct {
	XMLName         xml.Name        `xml:"response"`
	AccountBalance  Accountbalance  `xml:"accountbalance"`
	AccountHoldings Accountholdings `xml:"accountholdings"`
}

type AccountDetailBalanceResponse struct {
	XMLName        xml.Name       `xml:"response"`
	AccountBalance Accountbalance `xml:"accountbalance"`
}

type AccountDetailHoldingsResponse struct {
	XMLName         xml.Name        `xml:"response"`
	AccountHoldings Accountholdings `xml:"accountholdings"`
}

type SecurityDetail struct {
	XMLName      xml.Name `xml:"security"`
	Cusip        string   `xml:"cusip"`
	ID           string   `xml:"id"`
	SecurityType string   `xml:"sectyp"`
	Symbol       string   `xml:"symbol"`
}

type TransactionDetail struct {
	XMLName        xml.Name       `xml:"transaction"`
	AccountType    int            `xml:"accounttype"`
	Commission     float64        `xml:"commission"`
	Description    string         `xml:"description"`
	Fee            float64        `xml:"fee"`
	Price          float64        `xml:"price"`
	Quantity       float64        `xml:"quantity"`
	SecurityFee    float64        `xml:"secfee"`
	Security       SecurityDetail `xml:"security"`
	SettlementDate time.Time      `xml:"settlementdate"`
	Side           int            `xml:"side"`
	Source         string         `xml:"source"`
	TradeDate      time.Time      `xml:"tradedate"`
	TransactionID  int            `xml:"transactionid"`
}

type Transaction struct {
	XMLName     xml.Name          `xml:"transaction"`
	Activity    string            `xml:"activity"`
	Amount      float64           `xml:"amount"`
	Date        time.Time         `xml:"date"`
	Description string            `xml:"desc"`
	Symbol      string            `xml:"symbol"`
	Detail      TransactionDetail `xml:transaction>transaction`
}

type AccountHistoryResponse struct {
	XMLName      xml.Name      `xml:"response"`
	Transactions []Transaction `xml:"transactions>transaction"`
}
