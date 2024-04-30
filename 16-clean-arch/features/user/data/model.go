package data

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// ID        string `gorm:"primarykey"`
	// CreatedAt time.Time
	// UpdatedAt time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string
	Email     string `gorm:"unique"`
	Password  string
	Phone     string
	Address   string
	StoreName string
}
