type AllyResponse struct {
}

type AllyAccounts struct {
	Accounts []AllyAccount
}

type AllyAccount struct {
	Summary AccountSummary `xml:"accountsummary"`
}

type AccountSummary struct {
	Account     int    `xml:"accountsummary>account"`
	AccountName string `xml:"accountsummary>accountname"`
	Balance     AccountBalance
	Holdings    AccountHoldings
}

type AccountBalance struct {
	Account      int
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