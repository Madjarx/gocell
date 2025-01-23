package main

import (
	"fmt"
	"os/exec"
)

// Transaction represents a blockchain transaction
type Transaction struct {
	From      string
	To        string
	Amount    string
	Signature string
}

// TransactionManager handles transaction operations
type TransactionManager struct {
	cliPath string
}

func NewTransactionManager(cliPath string) *TransactionManager {
	return &TransactionManager{
		cliPath: cliPath,
	}
}

func (tm *TransactionManager) executeCommand(args ...string) (string, error) {
	cmd := exec.Command(tm.cliPath, args...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), fmt.Errorf("command failed: %v\nOutput: %s", err, output)
	}
	return string(output), nil
}

func (tm *TransactionManager) CreateTransaction(from, to, amount string) (*Transaction, error) {
	output, err := tm.executeCommand("srv_xchange", "orders", "-net", "Backbone", "-status", "opened")
	if err != nil {
		return nil, fmt.Errorf("failed to create transaction: %v", err)
	}
	fmt.Printf("Command output: %s\n", output)

	return &Transaction{
		From:   from,
		To:     to,
		Amount: amount,
	}, nil
}

func main() {
	// Create transaction manager instance
	// tm := NewTransactionManager("/path/to/cellframe-cli")
	tm := NewTransactionManager("/opt/cellframe-node/bin/cellframe-node-cli")

	// Test executing a system command
	cmd := exec.Command("ls", "-l")
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("ls output:\n%s\n", output)

	// Test cellframe CLI command
	tx, err := tm.CreateTransaction("sender", "receiver", "100")
	if err != nil {
		fmt.Printf("Transaction error: %v\n", err)
		return
	}
	fmt.Printf("Transaction created: %+v\n", tx)
}
