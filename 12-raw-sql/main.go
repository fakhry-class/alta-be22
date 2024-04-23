package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// _ "github.com/go-sql-driver/mysql"
	"github.com/go-sql-driver/mysql"
)

type User struct {
	ID         uint
	OwnerName  string
	Email      string
	Password   string
	Phone      string
	Sex        string
	Address    string
	BankNumber string
	StoreName  string
}

func main() {

	// Get a database handle.
	var db *sql.DB
	var err error
	// cara 1
	// Capture connection properties.
	// cfg := mysql.Config{
	// 	User:   "root",
	// 	Passwd: "qwerty123",
	// 	Net:    "tcp",
	// 	Addr:   "127.0.0.1:3306",
	// 	DBName: "db_be22",
	// }
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   fmt.Sprintf("%s:%s", os.Getenv("DBHOST"), os.Getenv("DBPORT")),
		DBName: os.Getenv("DBNAME"),
	}
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	// cara 2
	// <username>:<password>@tcp(<hostname>:<port-db>)/<db-name>
	// db, err = sql.Open("mysql", "root:qwerty123@tcp(127.0.0.1:3306)/db_be22")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//check apakah sudah berhasil connect ke db
	pingErr := db.Ping()
	if pingErr != nil { // jika terjadi error
		log.Fatal(pingErr)
	}
	defer db.Close()
	fmt.Println("Berhasil Connect ke DB!")

	fmt.Println("Menu:")
	fmt.Println("[1]. Read data user")
	fmt.Println("[2]. Insert data user")
	fmt.Println("[3]. Delete data user")
	// TASK dikumpulkan maks 15 april 2024 jam 18.00
	// nama repo: raw-sql
	fmt.Println("[4]. Update data user by id")
	fmt.Println("[5]. Read data user by email")
	fmt.Println("[6]. Read data product")
	fmt.Println("[7]. Insert data product")
	fmt.Println("[8]. Read data product by id product")
	fmt.Println("choose your menu:")
	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		// variable untuk menampung seluruh data user yg dibaca di query select
		var allUsers []User
		// menjalankan query select
		rows, errSelect := db.Query("SELECT id, owner_name, email, password, phone, sex, address, bank_number, store_name FROM users")
		// handling error ketika menjalankan query select
		if errSelect != nil {
			log.Fatal("terjadi error ", errSelect)
		}
		defer rows.Close()
		// loop untuk membaca per baris
		// pakai loop karena ada kemungkinan ada banyak baris yg dibaca
		for rows.Next() {
			// var untuk menampung data user per baris
			var userRow User
			// proses scanning dan mapping data row ke var penampung(userRow)
			errScan := rows.Scan(&userRow.ID, &userRow.OwnerName, &userRow.Email, &userRow.Password, &userRow.Phone, &userRow.Sex, &userRow.Address, &userRow.BankNumber, &userRow.StoreName)
			if errScan != nil {
				log.Fatal("error scan ", errScan)
			}

			allUsers = append(allUsers, userRow)
		}
		for _, v := range allUsers {
			fmt.Println(v.OwnerName)
		}

	case 2:
		fmt.Println("masukkan data user:")
		var newUser = User{}
		fmt.Println("masukkan nama user:")
		fmt.Scanln(&newUser.OwnerName)
		fmt.Println("masukkan email user:")
		fmt.Scanln(&newUser.Email)
		fmt.Println("masukkan password user:")
		fmt.Scanln(&newUser.Password)
		fmt.Println("masukkan no telepon user:")
		fmt.Scanln(&newUser.Phone)
		fmt.Println("masukkan jenis kelamin user:")
		fmt.Scanln(&newUser.Sex)
		fmt.Println("masukkan alamat user:")
		fmt.Scanln(&newUser.Address)
		fmt.Println("masukkan no bank user:")
		fmt.Scanln(&newUser.BankNumber)
		fmt.Println("masukkan nama toko user:")
		fmt.Scanln(&newUser.StoreName)

		result, errInsert := db.Exec("INSERT INTO users (owner_name, email, password, phone, sex, address, bank_number, store_name) values(?,?,?,?,?,?,?,?)", newUser.OwnerName, newUser.Email, newUser.Password, newUser.Phone, newUser.Sex, newUser.Address, newUser.BankNumber, newUser.StoreName)
		if errInsert != nil {
			log.Fatal("error insert ", errInsert)
		} else {
			rowAff, _ := result.RowsAffected()
			if rowAff > 0 {
				fmt.Println("insert berhasil")
			} else {
				fmt.Println("insert gagal")
			}
		}
	case 3:
		fmt.Println("masukkan email yang akan dihapus:")
		var email string
		fmt.Scanln(&email)
		result, errDelete := db.Exec("DELETE FROM users WHERE email = ?", email)
		if errDelete != nil {
			log.Fatal("error delete ", errDelete)
		} else {
			rowAff, _ := result.RowsAffected()
			if rowAff > 0 {
				fmt.Println("delete berhasil")
			} else {
				fmt.Println("delete gagal")
			}
		}

	case 4:
		fmt.Println("update")
	case 5:
		fmt.Println("read data")
	}

}
