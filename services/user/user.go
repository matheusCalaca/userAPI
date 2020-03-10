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
	if err == nil {
		return userBD, err
	}

	//todo: refatorar esta parte
	return user, errors.New("Error no metodo")
}

func validaInsertUser(userVariavel userModel.User) (bool, error) {
	if len(userVariavel.CpfCnpj) <= 0 {
		return false, errors.New("CPF CNPJ não pode ser vazio!")
	}
	return true, nil
}
