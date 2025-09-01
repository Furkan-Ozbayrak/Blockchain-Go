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

🇹🇷 Türkçe Versiyon
Açıklama

Bu, Go dili ile yazılmış basit bir etkileşimli blockchain projesidir.
Birden fazla işlem, cüzdan yönetimi, proof-of-work mining ve dinamik cüzdan oluşturmayı destekler.
Giriş seviyesindeki kullanıcılar için blockchain mantığını öğrenmeye yönelik tasarlanmıştır.

Özellikler

Çoklu cüzdanlar (varsayılan: Furkan, Begüm, Mahmut, Beyza)

Yeni kullanıcılar için dinamik cüzdan oluşturma

Her blokta birden fazla işlem

Proof-of-Work mining

Blockchain doğrulama

Cüzdan bakiyesini kontrol etme

Etkileşimli komut satırı arayüzü (CLI)

Kurulum

Go
 yüklü olduğundan emin olun (1.20+ önerilir)

Repoyu klonlayın:

git clone <repo-link>
cd <repo-folder>


Programı çalıştırın:

go run blockchain.go

Kullanım
1: Birden fazla işlem içeren yeni blok ekle
2: Blockchain'i yazdır
3: Cüzdan bakiyesini kontrol et
4: Çıkış
Seçiminiz: 1


Örnek İşlemler

Gönderen: Furkan
Alıcı: Mahmut
Miktar: 30

Gönderen: Beyza
Alıcı: Begüm
Miktar: 50

Gönderen: done
Blok mine edildi ve blockchain'e eklendi!


Blockchain’i Yazdır

================ Block 0 ================
Zaman Damgası: 2025-09-01 15:10:00
Nonce: 0
İşlemler:
Hash: a1b2c3d4...
PrevHash: 
=========================================

================ Block 1 ================
Zaman Damgası: 2025-09-01 15:11:05
Nonce: 4021
İşlemler:
  Furkan -> Mahmut : 30.00
  Beyza -> Begüm : 50.00
Hash: 000f1a2b3c...
PrevHash: a1b2c3d4...
=========================================


Cüzdan Bakiyesini Kontrol Et

Furkan bakiye: 70.00
Mahmut bakiye: 130.00
Beyza bakiye: 50.00
Begüm bakiye: 150.00

Notlar

Yalnızca kayıtlı veya dinamik olarak oluşturulan cüzdanlar para gönderebilir.

Proof-of-Work zorluğu varsayılan olarak 3’tür.

Blockchain öğrenimi ve demo amaçlı eğitim projesidir.
