package cli

/*
Transactions (transfer of tokens)
https://wiki.cellframe.net/Resources/Node+Commands/Node+Command+-+TX_CREATE
*/
type TransactionParams struct {
	Network    string
	Chain      string
	Value      string
	Token      string
	ToAddress  string
	FromWallet string // HAS TO BE ACCESSIBLE VIA NODE
	Fee        string
}

/*
Exchanges
https://wiki.cellframe.net/02.+Learn/DEX/Decentralized+Exchange+Service+(DEX)
*/
type ExchangeParams struct {
	Network   string
	TokenSell string
	TokenBuy  string
	Wallet    string
	Value     string
	Rate      string // HAS TO BE IN A DECIMAL FORMAT 1.0 (example)
	Fee       string
}

type ExchangePurchaseParams struct {
	Network   string
	OrderHash string
	Wallet    string 
	Value     string // Again has to be accessible via node
	Fee       string
}


type WalletActivationParams struct {
	Wallet   string
	Password string
}

type WalletInfoParams struct {
	Wallet  string
	Network string
}

type CommandOptions map[string]string
