# Interactive Blockchain in Go

## 🌐 English Version

### Description
This is a simple interactive blockchain project written in Go.  
It supports multiple transactions, wallet management, proof-of-work mining, and dynamic wallet creation.  
Designed for beginner-level blockchain demonstration and learning purposes.

### Features
- Multiple wallets (default: Furkan, Begüm, Mahmut, Beyza)  
- Dynamic wallet creation for new users  
- Multiple transactions per block  
- Proof-of-Work mining  
- Blockchain validation  
- Check wallet balance  
- Interactive command-line interface (CLI)

### Installation
1. Make sure you have [Go](https://golang.org/dl/) installed (version 1.20+ recommended)  
2. Clone the repository:  
   ```bash
   git clone <(https://github.com/Furkan-Ozbayrak)>
   cd <(https://github.com/Furkan-Ozbayrak/Blockchain-Go)>

Run the program:

go run blockchain.go

Usage:

1: Add new block with multiple transactions
2: Print blockchain
3: Check wallet balance
4: Exit
Choose an option: 1

Example Transactions : 

================ Block 0 ================
Timestamp: 2025-09-01 15:10:00
Nonce: 0
Transactions:
Hash: a1b2c3d4...
PrevHash: 
=========================================

================ Block 1 ================
Timestamp: 2025-09-01 15:11:05
Nonce: 4021
Transactions:
  Furkan -> Mahmut : 30.00
  Beyza -> Begüm : 50.00
Hash: 000f1a2b3c...
PrevHash: a1b2c3d4...
=========================================
Print Blockchain : 

================ Block 0 ================
Timestamp: 2025-09-01 15:10:00
Nonce: 0
Transactions:
Hash: a1b2c3d4...
PrevHash: 
=========================================

================ Block 1 ================
Timestamp: 2025-09-01 15:11:05
Nonce: 4021
Transactions:
  Furkan -> Mahmut : 30.00
  Beyza -> Begüm : 50.00
Hash: 000f1a2b3c...
PrevHash: a1b2c3d4...
=========================================

Check Wallet Balance : 

Furkan balance: 70.00
Mahmut balance: 130.00
Beyza balance: 50.00
Begüm balance: 150.00

Notes

Only registered or dynamically created wallets can send funds.

Proof-of-Work difficulty is set to 3 by default.

Educational project for blockchain learning and demonstration.

