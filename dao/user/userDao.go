package userDao

import (
	"log"
	models "userAPI/models/user"

	"github.com/jinzhu/gorm"
)

func InserirPessoaBD(pessoa models.Pessoa, db *gorm.DB) (models.Pessoa, error) {
	// result, err := dbmap.Exec(`INSERT INTO pessoa(NOME, SOBRENOME, DATA_NASCIMENTO, CPF, RG, EMAIL, CREATED_AT) VALUES (?, ?, ?, ?, ?, ?, ?)`,
	// 	pessoa.Nome, pessoa.Sobrenome, pessoa.DataNascimento, pessoa.CPF, pessoa.RG, pessoa.Email, pessoa.CreatedAt)
	// if err != nil {
	// 	return pessoa, dbControllers.CheckErr(err, "Falha ao inserir")
	// }
	// pessoaId, errInsert := result.LastInsertId()
	// if errInsert != nil {
	// 	return pessoa, dbControllers.CheckErr(err, "Falha ao inserir")
	// }
	// pessoa.ID = uint(pessoaId)
	// telefone := &models.Telefone{
	// 	DDD:    62,
	// 	Numero: 62999627272,
	// 	Tipo:   "CELULAR",
	// }
	// pessoa.CreatedAt = time.Now()
	result := db.Create(&pessoa)
	log.Print(&result.Error)
	return pessoa, nil

}

// func SelectAllUser() (userModel.Users, error) {
// 	var users userModel.Users
// 	_, err := dbmap.Select(&users, "select * from user")
// 	if err != nil {
// 		return nil, dbControllers.CheckErr(err, "Erro ao Buscar Dados : ")
// 	}

// 	return users, nil
// }
