package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)


type Transaction struct {
	Sender   string
	Receiver string
	Amount   float64
}


type Block struct {
	Index        int
	Timestamp    string
	Transactions []Transaction
	Hash         string
	PrevHash     string
	Nonce        int
}

// Wallet system (baþlangýç cüzdanlarý)
var wallets = map[string]float64{
	"Furkan":   100,
	"Begüm":     100,
	"Mahmut": 100,
	"Beyza" : 100,
}

// Calculate hash
func calculateHash(block Block) string {
	var transactionsData []string
	for _, tx := range block.Transactions {
		transactionsData = append(transactionsData, tx.Sender+tx.Receiver+fmt.Sprintf("%.2f", tx.Amount))
	}
	record := strconv.Itoa(block.Index) + block.Timestamp + strings.Join(transactionsData, "") + block.PrevHash + strconv.Itoa(block.Nonce)
	hash := sha256.New()
	hash.Write([]byte(record))
	return hex.EncodeToString(hash.Sum(nil))
}

// Proof-of-Work
func mineBlock(block *Block, difficulty int) {
	target := strings.Repeat("0", difficulty)
	for {
		block.Hash = calculateHash(*block)
		if strings.HasPrefix(block.Hash, target) {
			break
		}
		block.Nonce++
	}
}

// Generate new block
func generateBlock(oldBlock Block, transactions []Transaction, difficulty int) Block {
	newBlock := Block{
		Index:        oldBlock.Index + 1,
		Timestamp:    time.Now().String(),
		Transactions: transactions,
		PrevHash:     oldBlock.Hash,
		Nonce:        0,
	}
	mineBlock(&newBlock, difficulty)
	return newBlock
}

// Validate block
func isBlockValid(newBlock, oldBlock Block) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}
	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}
	if calculateHash(newBlock) != newBlock.Hash {
		return false
	}
	return true
}

// Print blockchain
func printBlockchain(blockchain []Block) {
	for _, block := range blockchain {
		fmt.Printf("\n================ Block %d ================\n", block.Index)
		fmt.Printf("Timestamp: %s\n", block.Timestamp)
		fmt.Printf("Nonce: %d\n", block.Nonce)
		fmt.Printf("Transactions:\n")
		for _, tx := range block.Transactions {
			fmt.Printf("  %s -> %s : %.2f\n", tx.Sender, tx.Receiver, tx.Amount)
		}
		fmt.Printf("Hash: %s\n", block.Hash)
		fmt.Printf("PrevHash: %s\n", block.PrevHash)
		fmt.Printf("=========================================\n")
	}
}

// Check wallet balance
func checkBalance(name string) float64 {
	if balance, ok := wallets[name]; ok {
		return balance
	}
	return 0
}

func main() {
	difficulty := 3
	blockchain := []Block{}

	// Genesis block
	genesisBlock := Block{0, time.Now().String(), []Transaction{}, "", "", 0}
	genesisBlock.Hash = calculateHash(genesisBlock)
	blockchain = append(blockchain, genesisBlock)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n1: Add new block with multiple transactions")
		fmt.Println("2: Print blockchain")
		fmt.Println("3: Check wallet balance")
		fmt.Println("4: Exit")
		fmt.Print("Choose an option: ")
		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			var transactions []Transaction
			for {
				fmt.Println("\nAdd a transaction or type 'done' to finish:")
				fmt.Print("Sender: ")
				sender, _ := reader.ReadString('\n')
				sender = strings.TrimSpace(sender)
				if sender == "done" {
					break
				}

				fmt.Print("Receiver: ")
				receiver, _ := reader.ReadString('\n')
				receiver = strings.TrimSpace(receiver)

				fmt.Print("Amount: ")
				amountStr, _ := reader.ReadString('\n')
				amountStr = strings.TrimSpace(amountStr)
				amount, err := strconv.ParseFloat(amountStr, 64)
				if err != nil {
					fmt.Println("Invalid amount. Try again.")
					continue
				}

				// Dinamik sender cüzdan
				balance, exists := wallets[sender]
				if !exists {
					fmt.Printf("Sender %s not found. Creating new wallet with 0 balance.\n", sender)
					wallets[sender] = 0
					balance = 0
				}

				if balance < amount {
					fmt.Println("Insufficient balance. Try again.")
					continue
				}

				// Dinamik receiver cüzdan
				if _, exists := wallets[receiver]; !exists {
					fmt.Printf("Receiver %s not found. Creating new wallet with 0 balance.\n", receiver)
					wallets[receiver] = 0
				}

				// Ýþlem bakiyelerini güncelle
				wallets[sender] -= amount
				wallets[receiver] += amount

				tx := Transaction{Sender: sender, Receiver: receiver, Amount: amount}
				transactions = append(transactions, tx)
			}

			if len(transactions) > 0 {
				lastBlock := blockchain[len(blockchain)-1]
				newBlock := generateBlock(lastBlock, transactions, difficulty)
				if isBlockValid(newBlock, lastBlock) {
					blockchain = append(blockchain, newBlock)
					fmt.Println("Block mined and added to blockchain!")
				} else {
					fmt.Println("Failed to add block.")
				}
			} else {
				fmt.Println("No transactions added.")
			}

		case "2":
			printBlockchain(blockchain)

		case "3":
			fmt.Print("Enter wallet name: ")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)
			fmt.Printf("%s balance: %.2f\n", name, checkBalance(name))

		case "4":
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}
