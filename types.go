package main

import "encoding/xml"

type AccountSummary struct {
	XMLName        xml.Name `xml:"accountsummary"`
	Text           string   `xml:",chardata"`
	Account        string   `xml:"account"`
	Accountbalance struct {
		Text         string `xml:",chardata"`
		Account      string `xml:"account"`
		Accountvalue string `xml:"accountvalue"`
		Buyingpower  struct {
			Text                       string `xml:",chardata"`
			Cashavailableforwithdrawal string `xml:"cashavailableforwithdrawal"`
			Daytrading                 string `xml:"daytrading"`
			Equitypercentage           string `xml:"equitypercentage"`
			Options                    string `xml:"options"`
			Soddaytrading              string `xml:"soddaytrading"`
			Sodoptions                 string `xml:"sodoptions"`
			Sodstock                   string `xml:"sodstock"`
			Stock                      string `xml:"stock"`
		} `xml:"buyingpower"`
		Fedcall   string `xml:"fedcall"`
		Housecall string `xml:"housecall"`
		Money     struct {
			Text              string `xml:",chardata"`
			Accruedinterest   string `xml:"accruedinterest"`
			Cash              string `xml:"cash"`
			Cashavailable     string `xml:"cashavailable"`
			Marginbalance     string `xml:"marginbalance"`
			Mmf               string `xml:"mmf"`
			Total             string `xml:"total"`
			Uncleareddeposits string `xml:"uncleareddeposits"`
			Unsettledfunds    string `xml:"unsettledfunds"`
			Yield             string `xml:"yield"`
		} `xml:"money"`
		Securities struct {
			Text         string `xml:",chardata"`
			Longoptions  string `xml:"longoptions"`
			Longstocks   string `xml:"longstocks"`
			Options      string `xml:"options"`
			Shortoptions string `xml:"shortoptions"`
			Shortstocks  string `xml:"shortstocks"`
			Stocks       string `xml:"stocks"`
			Total        string `xml:"total"`
		} `xml:"securities"`
	} `xml:"accountbalance"`
	Accountholdings struct {
		Text        string `xml:",chardata"`
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
			Gainloss   string `xml:"gainloss"`
			Instrument struct {
				Text   string `xml:",chardata"`
				Cusip  string `xml:"cusip"`
				Desc   string `xml:"desc"`
				Factor string `xml:"factor"`
				Sectyp string `xml:"sectyp"`
				Sym    string `xml:"sym"`
			} `xml:"instrument"`
			Marketvalue       string `xml:"marketvalue"`
			Marketvaluechange string `xml:"marketvaluechange"`
			Price             string `xml:"price"`
			Purchaseprice     string `xml:"purchaseprice"`
			Qty               string `xml:"qty"`
			Quote             struct {
				Text      string `xml:",chardata"`
				Change    string `xml:"change"`
				Lastprice string `xml:"lastprice"`
			} `xml:"quote"`
			Underlying string `xml:"underlying"`
		} `xml:"holding"`
		Totalsecurities string `xml:"totalsecurities"`
	} `xml:"accountholdings"`
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
