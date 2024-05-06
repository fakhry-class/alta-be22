package service

import (
	"be22/clean-arch/app/middlewares"
	"be22/clean-arch/features/user"
	"be22/clean-arch/utils/encrypts"
	"errors"
)

type userService struct {
	userData    user.DataInterface
	hashService encrypts.HashInterface
}

func New(ud user.DataInterface, hash encrypts.HashInterface) user.ServiceInterface {
	return &userService{
		userData:    ud,
		hashService: hash,
	}

}

// Create implements user.ServiceInterface.
func (u *userService) Create(input user.Core) error {
	// validasi /logic
	if input.Name == "" || input.Email == "" || input.Password == "" {
		return errors.New("[validation] nama/email/password tidak boleh kosong")
	}

	if input.Password != "" {
		// proses hash password
		result, errHash := u.hashService.HashPassword(input.Password)
		if errHash != nil {
			return errHash
		}
		input.Password = result
	}
	// memanggil method di data layer
	err := u.userData.Insert(input)
	if err != nil {
		return err
	}
	return nil

	// return u.userData.Insert(input)
}

// GetAll implements user.ServiceInterface.
func (u *userService) GetAll() ([]user.Core, error) {
	return u.userData.SelectAll()
}

// Delete implements user.ServiceInterface.
func (u *userService) Delete(id uint) error {
	if id <= 0 {
		return errors.New("id not valid")
	}
	return u.userData.Delete(id)
}

// Login implements user.ServiceInterface.
func (u *userService) Login(email string, password string) (data *user.Core, token string, err error) {
	data, err = u.userData.SelectByEmail(email)
	if err != nil {
		return nil, "", err
	}

	isLoginValid := u.hashService.CheckPasswordHash(data.Password, password)
	// ketika isloginvalid = true, maka login berhasil
	if !isLoginValid {
		return nil, "", errors.New("[validation] password tidak sesuai.")
	}
	token, errJWT := middlewares.CreateToken(int(data.ID))
	if errJWT != nil {
		return nil, "", errJWT
	}
	return data, token, nil
}

// GetById implements user.ServiceInterface.
func (u *userService) GetById(id uint) (data *user.Core, err error) {
	if id <= 0 {
		return nil, errors.New("[validation] id not valid")
	}
	return u.userData.SelectById(id)
}
