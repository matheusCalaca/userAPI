package userNegocio

import ( 
	userModel "userAPI/model/user" 
	dbControllers "userAPI/controlles/db"
)

var userModel = userModel.user

func InsertUser(user *userModel){
		
	user.CreatedAt = time.Now().UnixNano()
	user.UpdatedAt = time.Now().UnixNano()
	log.Println(userJson)

	insert, errInsert := dbmap.Exec(`INSERT INTO user (NOME, ENDERECO, IDADE, CPF_CNPJ, CREATED_AT) VALUES (?, ?, ?, ?, ?)`, userJson.Nome, userJson.Endereco, userJson.Idade, userJson.CpfCnpj, userJson.CreatedAt)
	if insert != nil {
		user_id, err := insert.LastInsertId()

		if err == nil {
			content := &userModel.User{
				ID:        user_id,
				Nome:      userJson.Nome,
				CpfCnpj:   userJson.CpfCnpj,
				Idade:     userJson.Idade,
				Endereco:  userJson.Endereco,
				CreatedAt: userJson.CreatedAt,
			}
			c.JSON(201, content)
		} else {
			dbControllers.CheckErr(err, "Falha ao inserir")
		}
	}
	if errInsert != nil {
		dbControllers.CheckErr(errInsert, "Error ao inserir USER")
	}

}