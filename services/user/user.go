package userNegocio

import (
	"errors"
	"time"
	dbControllers "userAPI/controlles/db"
	userModel "userAPI/models/user"
)

var dbmap = dbControllers.InitDb()

func InsertUser(user userModel.User) (userModel.User, error) {

	_, errValidation := validaInsertUser(user)
	if errValidation != nil {
		return user, errValidation
	}

	user.CreatedAt = time.Now().UnixNano()
	user.UpdatedAt = time.Now().UnixNano()

	insert, errInsert := dbmap.Exec(`INSERT INTO user (NOME, ENDERECO, IDADE, CPF_CNPJ, CREATED_AT) VALUES (?, ?, ?, ?, ?)`, user.Nome, user.Endereco, user.Idade, user.CpfCnpj, user.CreatedAt)
	if insert != nil {
		user_id, err := insert.LastInsertId()

		if err == nil {
			content := userModel.User{
				ID:        user_id,
				Nome:      user.Nome,
				CpfCnpj:   user.CpfCnpj,
				Idade:     user.Idade,
				Endereco:  user.Endereco,
				CreatedAt: user.CreatedAt,
			}
			return content, nil
		} else {
			dbControllers.CheckErr(err, "Falha ao inserir")
			return user, errors.New("Falha ao inserir")
		}
	}
	if errInsert != nil {
		dbControllers.CheckErr(errInsert, "Error ao inserir USER")
		return user, errors.New("Error ao inserir USER")
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
