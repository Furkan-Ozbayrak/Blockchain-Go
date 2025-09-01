package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("=== Hash Test / Debug ===")

	// Örnek transactionlar
	transactions := []Transaction{
		{Sender: "Alice", Receiver: "Bob", Amount: 25},
		{Sender: "Charlie", Receiver: "Alice", Amount: 10},
	}

	// Örnek blok
	testBlock := Block{
		Index:        1,
		Timestamp:    time.Now().String(),
		Transactions: transactions,
		PrevHash:     "0000000000000000",
		Nonce:        0,
	}

	// Adým adým transaction stringlerini oluþtur
	var transactionsData []string
	for _, tx := range testBlock.Transactions {
		txString := tx.Sender + tx.Receiver + fmt.Sprintf("%.2f", tx.Amount)
		transactionsData = append(transactionsData, txString)
		fmt.Printf("Transaction string: %s\n", txString)
	}

	// Blok record string’i oluþtur
	record := fmt.Sprintf("%d%s%s%s%d", testBlock.Index, testBlock.Timestamp, 
		strings.Join(transactionsData, ""), testBlock.PrevHash, testBlock.Nonce)
	fmt.Printf("\nCombined record string for hashing:\n%s\n", record)

	// Hash hesapla
	hash := calculateHash(testBlock)
	fmt.Printf("\nCalculated Hash: %s\n", hash)

	fmt.Println("\n--- End of Hash Test ---")
}
