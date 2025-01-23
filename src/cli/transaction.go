package cli

/*
    ## Transaction Service

    Transactions are a core part of the Cellframe network. 
    They are used to transfer tokens between wallets.

    This service provides a way to create and execute:
    1. transactions (& everything related to them)
    2. exchanges (& everything related to them)

*/

import (
    "log"
	"fmt"
)

type TransactionService struct {
    builder  *CommandBuilder
    executor *CommandExecutor
}

func NewTransactionService(cliPath string) *TransactionService {
    return &TransactionService{
        builder:  NewCommandBuilder(cliPath),
        executor: NewCommandExecutor(),
    }
}

func (ts *TransactionService) CreateUnsignedTransaction(params TransactionParams) (string, error) {
    path, args, err := ts.builder.BuildTxCreateCommand(params)
    if err != nil {
        return "", fmt.Errorf("invalid parameters: %w", err)
    }
    
    log.Printf("Executing transaction creation command: %s %v", path, args)
    return ts.executor.Execute(path, args)
}



func (ts *TransactionService) OrderCreate(params ExchangeParams) (string, error) {
    path, args, err := ts.builder.BuildOrderCreateCommand(params)
    if err != nil {
        return "", fmt.Errorf("invalid parameters: %w", err)
    }
    
    log.Printf("Executing transaction creation command: %s %v", path, args)
    return ts.executor.Execute(path, args)
}

func (ts *TransactionService) OrderPurchase(params ExchangePurchaseParams) (string, error) {
    path, args, err := ts.builder.BuildOrderPurchaseCommand(params)
    if err != nil {
        return "", fmt.Errorf("invalid parameters: %w", err)
    }
    
    log.Printf("Executing transaction creation command: %s %v", path, args)
    return ts.executor.Execute(path, args)
}



func (ts *TransactionService) WalletActivation(params WalletActivationParams) (string, error) {
    path, args, err := ts.builder.BuildWalletActivationCommand(params)
    if err != nil {
        return "", fmt.Errorf("invalid parameters: %w", err)
    }
    
    log.Printf("Executing transaction creation command: %s %v", path, args)
    return ts.executor.Execute(path, args)
}


func (ts *TransactionService) WalletInfo(params WalletInfoParams) (string, error) {
    path, args, err := ts.builder.BuildWalletInfoCommand(params)
    if err != nil {
        return "", fmt.Errorf("invalid parameters: %w", err)
    }
    
    log.Printf("Executing transaction creation command: %s %v", path, args)
    return ts.executor.Execute(path, args)
}


 