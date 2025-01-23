package cli

/*
    ## Command builder

    Composes the parameters into a proper CLI command
    returns the command and the arguments as a slice of strings
*/


import "fmt"


type CommandBuilder struct {
    cliPath string
}


func NewCommandBuilder(cliPath string) *CommandBuilder {
    return &CommandBuilder{cliPath: cliPath}
}


func (cb *CommandBuilder) BuildTxCreateCommand(params TransactionParams) (string, []string, error) {
    if params.Network == "" || params.Chain == "" {
        return "", nil, fmt.Errorf("network and chain are required")
    }
    
    args := []string{
        "tx_create",
        "-net", params.Network,
        "-chain", params.Chain,
        "-value", params.Value,
        "-token", params.Token,
        "-to_addr", params.ToAddress,
        "-from_wallet", params.FromWallet,
        "-fee", params.Fee,
    }
    
    return cb.cliPath, args, nil
}




// #region Exchange Service (srv_xchange_...) section

// Build the srv_xchange command
// https://wiki.cellframe.net/Resources/Node+Commands/Node+Command+-+SRV_XCHANGE+ORDER+CREATE
//
func (cb *CommandBuilder) BuildOrderCreateCommand(params ExchangeParams) (string, []string, error) {
    if params.Network == "" || params.TokenSell == "" || params.TokenBuy == "" {
        return "", nil, fmt.Errorf("network, token sell and token buy are required")
    }

    args := []string{
        "srv_xchange",
        "order",
        "create",
        "-net", params.Network,
        "-token_sell", params.TokenSell,
        "-token_buy", params.TokenBuy,
        "-w", params.Wallet,
        "-value", params.Value,
        "-rate", params.Rate,
        "-fee", params.Fee,
    }

    return cb.cliPath, args, nil
}



// Build the srv_xchange purchase command
// https://wiki.cellframe.net/Resources/Node+Commands/Node+Command+-+SRV_XCHANGE+PURCHASE
//
func (cb *CommandBuilder) BuildOrderPurchaseCommand(params ExchangePurchaseParams) (string, []string, error) {
    if params.Network == "" || params.OrderHash == "" {
        return "", nil, fmt.Errorf("network and order hash are required")
    }

    args := []string{
        "srv_xchange",
        "purchase",
        "-order", params.OrderHash,
        "-net", params.Network,
        "-w", params.Wallet,
        "-value", params.Value,
        "-fee", params.Fee,
    }

    return cb.cliPath, args, nil
}

// #endregion



// #region Wallet services (wallet info, etc.)

// Build the wallet info command
// https://wiki.cellframe.net/Resources/Node+Commands/Node+Command+-+WALLET+INFO
//
func (cb *CommandBuilder) BuildWalletInfoCommand (params WalletInfoParams) (string, []string, error) {
    if params.Network == "" || params.Wallet == "" {
        return "", nil, fmt.Errorf("network and wallet are required")
    }

    args := []string{
        "wallet",
        "info",
        "-net", params.Network,
        "-w", params.Wallet,
    }

    return cb.cliPath, args, nil
}

// Build the wallet activation command
// https://wiki.cellframe.net/Resources/Node+Commands/Node+Command+-+WALLET+ACTIVATE
//
func (cb *CommandBuilder) BuildWalletActivationCommand (params WalletActivationParams) (string, []string, error) {
    if params.Wallet == "" || params.Password == "" {
        return "", nil, fmt.Errorf("wallet and password are required")
    }

    args := []string{
        "wallet",
        "activate",
        "-w", params.Wallet,
        "-p", params.Password,
    }

    return cb.cliPath, args, nil
}
// #endregion