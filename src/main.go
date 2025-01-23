package main

/*
    ## Main

    A set of calls to showcase how it looks like running
    the CLI commands over the node via go's EXEC

    comment out regions you dont want to test

    I shall elaborate more on the documentation
*/

import (
	"fmt"
	"go-cellframe-cli/src/cli"
	"go-cellframe-cli/src/config"
)


func main() {

    // configuration, pass in the absolute path to the CLI's binary
	txService := cli.NewTransactionService(config.CliPath)


    /*
    * Set of examples for different types of CLI commands
    */


    // #region Transactions a.k.a. token transfers
	params := cli.TransactionParams{
		Network:    "Backbone",
		Chain:      "main",
		// Value:      "10000000000000000000", // 10 $CELL
		Value:      "100000000000000000", // 0.1 $CELL
		Token:      "CELL",
		ToAddress:  "Rj7J7MiX2bWy8sNyb2Zn74XqguhasMmiVvmM6amcQVyAmCTuqfutnFa8wfe4ts93cuDdEdUDQ59fCbbqBMisvzRySWCv6Hp6sDS1kbjz",
		FromWallet: "exchange",
		// Fee:        "0.05e+18",
		Fee: "50000000000000000", // Its more stable to have it in this format (Datoshi)
	}

    txHash, err := txService.CreateUnsignedTransaction(params)
	if err != nil {
		fmt.Printf("Error creating transaction: \n%v\n", err)
		return
	}

    fmt.Printf("Successfully created transaction:\n%s\n", txHash)
    // #endregion



    // #region Orders & Exchanges
	paramsOrder := cli.ExchangeParams{
		Network:   "Backbone",
		Value:     "10000000000000000000",
		TokenBuy:  "CELL",
		TokenSell: "QEVM",
		Wallet:    "exchange", // this is a wallet i created locally, i can access it and i have activated it
		Rate:      "1.0",      // Has to be in a decimal format 1.0 (example)
		Fee:       "50000000000000000",
	}

    // https://gist.github.com/Madjarx/e4c2e1ba09d1ca842c3d16773398d1e3
    // Look at the github link under the section (QEVM Sell - people selling $QEVM and we're giving them CELL)
    // pick any hash from there and use it as the OrderHash
	paramsOrderPurchase := cli.ExchangePurchaseParams{
		Network:   "Backbone",
		OrderHash: "0x3731105E08B5633B546A2D9FDC1C3A02FC7022EF6221AF6D4F7A6B80DB22C965",
		Wallet:    "exchange",
		Value:     "10000000000000000",
		Fee:       "50000000000000000",
	}

    orderHash, err := txService.OrderCreate(paramsOrder)
	if err != nil {
		fmt.Printf("Error creating order: \n%v\n", err)
		return
	}

	orderHashPurchase, err := txService.OrderPurchase(paramsOrderPurchase)
	if err != nil {
		fmt.Printf("Error purchasing order: \n%v\n", err)
		return
	}


	fmt.Printf("Successfully created order:\n%s\n", orderHash)
	fmt.Printf("Successfully purchased order:\n%s\n", orderHashPurchase)
    // #endregion



    // #region Wallets 
    paramsWalletInfo := cli.WalletInfoParams{
        Wallet:  "exchange", // wallet generated locally
        Network: "Backbone",
    }

    paramsWalletActivate := cli.WalletActivationParams{
        Wallet:   "exchange",
        Password: "changeme",
    }

    walletInfo, err := txService.WalletInfo(paramsWalletInfo)
    if err != nil {
        fmt.Printf("Error getting wallet info: \n%v\n", err)
        return
    }

    walletActivation, err := txService.WalletActivation(paramsWalletActivate)
    if err != nil {
        fmt.Printf("Error activating wallet: \n%v\n", err)
        return
    }


    fmt.Printf("Successfully got wallet info:\n%s\n", walletInfo)
    fmt.Printf("Successfully activated wallet:\n%s\n", walletActivation)
    // #endregion


}
