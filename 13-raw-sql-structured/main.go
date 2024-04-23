package main

import (
	"be22/raw-sql/controllers"
	"be22/raw-sql/entities"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// Get a database handle.
	var db *sql.DB
	var err error
	// cara 1
	// Capture connection properties.
	//
	//	cfg := mysql.Config{
	//		User:   "root",
	//		Passwd: "qwerty123",
	//		Net:    "tcp",
	//		Addr:   "127.0.0.1:3306",
	//		DBName: "db_be22",
	//	}
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

	// check apakah sudah berhasil connect ke db
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
		result, err := controllers.GetAllUserController(db)
		if err != nil {
			fmt.Println("error read users ", err.Error())
		} else {
			for _, v := range result {
				fmt.Println("nama:", v.OwnerName)
			}
		}

	case 2:
		fmt.Println("masukkan data user:")
		var newUser = entities.User{}
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

		row, err := controllers.AddUserController(db, newUser)
		if err != nil {
			fmt.Println("error insert users ", err.Error())
		} else {
			fmt.Println("insert success. row affected: ", row)
		}
	}
}
