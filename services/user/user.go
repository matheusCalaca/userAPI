package userNegocio

import (
	"errors"
	"time"
	userDao "userAPI/dao/user"
	userModel "userAPI/models/user"
)

func InsertUser(user userModel.User) (userModel.User, error) {

	_, errValidation := validaInsertUser(user)
	if errValidation != nil {
		return user, errValidation
	}

	user.CreatedAt = time.Now().UnixNano()
	user.UpdatedAt = time.Now().UnixNano()

	userBD, err := userDao.InserirUserBD(user)
	if err != nil {
		return user, err
	}

	return userBD, err
}

func validaInsertUser(userVariavel userModel.User) (bool, error) {
	if len(userVariavel.CpfCnpj) <= 0 {
		return false, errors.New("CPF CNPJ não pode ser vazio!")
	}
	return true, nil
}

func FindAllUser() (userModel.Users, error) {
	return userDao.SelectAllUser()
}
