package user

import "time"

type Core struct {
	ID        uint
	Name      string
	Email     string
	Password  string
	Phone     string
	Address   string
	StoreName string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type DataInterface interface {
	Insert(input Core) error
	SelectAll() ([]Core, error)
	Delete(id uint) error
}

type ServiceInterface interface {
	Create(input Core) error
	GetAll() ([]Core, error)
	Delete(id uint) error
}
