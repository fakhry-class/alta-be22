package controllers

import (
	"be22/raw-sql/entities"
	"database/sql"
	"errors"
)

func GetAllUserController(db *sql.DB) ([]entities.User, error) {
	// variable untuk menampung seluruh data user yg dibaca di query select
	var allUsers []entities.User
	// menjalankan query select
	rows, errSelect := db.Query("SELECT id, owner_name, email, password, phone, sex, address, bank_number, store_name FROM users")
	// handling error ketika menjalankan query select
	if errSelect != nil {
		// log.Fatal("terjadi error ", errSelect)
		return nil, errSelect
	}
	defer rows.Close()
	// loop untuk membaca per baris
	// pakai loop karena ada kemungkinan ada banyak baris yg dibaca
	for rows.Next() {
		// var untuk menampung data user per baris
		var userRow entities.User
		// proses scanning dan mapping data row ke var penampung(userRow)
		errScan := rows.Scan(&userRow.ID, &userRow.OwnerName, &userRow.Email, &userRow.Password, &userRow.Phone, &userRow.Sex, &userRow.Address, &userRow.BankNumber, &userRow.StoreName)
		if errScan != nil {
			// log.Fatal("error scan ", errScan)
			return nil, errors.New("error scan " + errScan.Error())
		}

		allUsers = append(allUsers, userRow)
	}
	// for _, v := range allUsers {
	// 	fmt.Println(v.OwnerName)
	// }

	return allUsers, nil
}

func AddUserController(db *sql.DB, input entities.User) (int, error) {
	result, errInsert := db.Exec("INSERT INTO users (owner_name, email, password, phone, sex, address, bank_number, store_name) values(?,?,?,?,?,?,?,?)", input.OwnerName, input.Email, input.Password, input.Phone, input.Sex, input.Address, input.BankNumber, input.StoreName)
	if errInsert != nil {
		// log.Fatal("error insert ", errInsert)
		return -1, errInsert
	} else {
		rowAff, _ := result.RowsAffected()
		if rowAff > 0 {
			// fmt.Println("insert berhasil")
			return int(rowAff), nil
		} else {
			// fmt.Println("insert gagal")
			return 0, errors.New("error. insert failed")
		}
	}
}
