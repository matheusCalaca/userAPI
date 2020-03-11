package userDao

import (
	dbControllers "userAPI/dao/db"
	userModel "userAPI/models/user"
)

var dbmap = dbControllers.InitDb()

func InserirUserBD(user userModel.User) (userModel.User, error) {
	insert, errInsert := dbmap.Exec(`INSERT INTO user (NOME, ENDERECO, IDADE, CPF_CNPJ, CREATED_AT) VALUES (?, ?, ?, ?, ?)`, user.Nome, user.Endereco, user.Idade, user.CpfCnpj, user.CreatedAt)
	if insert != nil {
		userId, err := insert.LastInsertId()

		if err == nil {
			content := userModel.User{
				ID:        userId,
				Nome:      user.Nome,
				CpfCnpj:   user.CpfCnpj,
				Idade:     user.Idade,
				Endereco:  user.Endereco,
				CreatedAt: user.CreatedAt,
			}
			return content, nil
		} else {
			return user, dbControllers.CheckErr(err, "Falha ao inserir")
		}
	}
	if errInsert != nil {
		return user, dbControllers.CheckErr(errInsert, "Error ao inserir USER")
	}
	return userModel.User{}, nil
}

func SelectAllUser() (userModel.Users, error) {
	var users userModel.Users
	_, err := dbmap.Select(&users, "select * from user")
	if err != nil {
		return nil, dbControllers.CheckErr(err, "Erro ao Buscar Dados : ")
	}

	return users, nil
}
