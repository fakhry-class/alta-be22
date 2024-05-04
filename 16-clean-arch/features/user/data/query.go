package data

import (
	"be22/clean-arch/features/user"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.DataInterface {
	return &userQuery{
		db: db,
	}
}

// Insert implements user.DataInterface.
func (u *userQuery) Insert(input user.Core) error {
	var userGorm User

	userGorm = User{
		Name:      input.Name,
		Email:     input.Email,
		Password:  input.Password,
		Phone:     input.Phone,
		Address:   input.Address,
		StoreName: input.StoreName,
	}
	tx := u.db.Create(&userGorm)

	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// SelectAll implements user.DataInterface.
func (u *userQuery) SelectAll() ([]user.Core, error) {
	var allUsers []User // var penampung data yg dibaca dari db
	tx := u.db.Find(&allUsers)
	if tx.Error != nil {
		return nil, tx.Error
	}
	//mapping
	var allUserCore []user.Core
	for _, v := range allUsers {
		allUserCore = append(allUserCore, user.Core{
			ID:        v.ID,
			Name:      v.Name,
			Email:     v.Email,
			Password:  v.Password,
			Phone:     v.Phone,
			Address:   v.Address,
			StoreName: v.StoreName,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return allUserCore, nil
}

// Delete implements user.DataInterface.
func (u *userQuery) Delete(id uint) error {
	tx := u.db.Delete(&User{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// Login implements user.DataInterface.
func (u *userQuery) SelectByEmail(email string) (*user.Core, error) {
	// select id, name, phone from users where email = xxxx and password = xxxxxx
	// select id, name, phone from users where email = xxxx

	// variable penampung datanya
	var userData User
	tx := u.db.Where("email = ?", email).First(&userData)
	if tx.Error != nil {
		return nil, tx.Error
	}
	// mapping
	var usercore = user.Core{
		ID:        userData.ID,
		Name:      userData.Name,
		Email:     userData.Email,
		Password:  userData.Password,
		Phone:     userData.Phone,
		Address:   userData.Address,
		StoreName: userData.StoreName,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
	}

	return &usercore, nil
}

// SelectById implements user.DataInterface.
func (u *userQuery) SelectById(id uint) (*user.Core, error) {
	// variable penampung datanya
	var userData User
	tx := u.db.First(&userData, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	// mapping
	var usercore = user.Core{
		ID:        userData.ID,
		Name:      userData.Name,
		Email:     userData.Email,
		Password:  userData.Password,
		Phone:     userData.Phone,
		Address:   userData.Address,
		StoreName: userData.StoreName,
		CreatedAt: userData.CreatedAt,
		UpdatedAt: userData.UpdatedAt,
	}

	return &usercore, nil
}
