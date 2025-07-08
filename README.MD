Requirements
Go 1.24+

SQLite3 installed
 - sudo apt update
 - sudo apt install sqlite3 libsqlite3-dev

(Optional) PostgreSQL support via pluggable adapter (pkg/database/sqldb)


Configuration
🔑 Generate RSA Keys (for JWT)
Create the folder:
 - mkdir -p config/keys
 
Generate private and public keys:
 - openssl genrsa -out config/keys/private.pem 2048
 - openssl rsa -in config/keys/private.pem -pubout -out config/keys/public.pem

Install Dependencies
 - go mod tidy
 - go mod vendor (to get vendor)

Run the App
 - go run ./cmd/main.go



Documentation API:
 - https://documenter.getpostman.com/view/6939703/2sB34cqidy