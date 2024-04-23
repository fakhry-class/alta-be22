#  Go SQL

## buat go mod
```
go mod init namaproject
```

## download driver mysql
```
go get -u github.com/go-sql-driver/mysql
```

## configurasi .env file
* buat `.env` file
* tulis env variable yang dibutuhkan
* buat `.gitignore` file, dan tulis `*.env` disana
* push code nya. gitignore akan membuat file .env tidak akan di upload ke repository

## cara load .env file
anda harus load env file terlebih dahulu di terminal. baru bisa menjalankan `go run main.go`
```bash
source .env
```
