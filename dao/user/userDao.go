package userdao

import (
	"log"
	"strings"
	models "userAPI/models/user"

	"github.com/jinzhu/gorm"
)

// InserirPessoaBD insere a pessoa no banco de dados
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
	caseUpperBeforeSave(&pessoa)
	result := db.Create(&pessoa)
	log.Print(&result.Error)
	return pessoa, nil

}

// caseUpperBeforeSave coloca em upper case al guns dados antes de salvar no banco
func caseUpperBeforeSave(pessoa *models.Pessoa) {
	for index, telefone := range pessoa.Telefone {
		pessoa.Telefone[index].Tipo = strings.ToUpper(telefone.Tipo)
	}

	for index, endereco := range pessoa.Endereco {
		pessoa.Endereco[index].Tipo = strings.ToUpper(endereco.Tipo)
		pessoa.Endereco[index].Bairro = strings.ToUpper(endereco.Bairro)
		pessoa.Endereco[index].Cidade = strings.ToUpper(endereco.Cidade)
		pessoa.Endereco[index].Logradouro = strings.ToUpper(endereco.Logradouro)
		pessoa.Endereco[index].Complemento = strings.ToUpper(endereco.Complemento)
		pessoa.Endereco[index].UF = strings.ToUpper(endereco.UF)
	}

	pessoa.Nome = strings.ToUpper(pessoa.Nome)
	pessoa.Sobrenome = strings.ToUpper(pessoa.Sobrenome)
}

// func SelectAllUser() (userModel.Users, error) {
// 	var users userModel.Users
// 	_, err := dbmap.Select(&users, "select * from user")
// 	if err != nil {
// 		return nil, dbControllers.CheckErr(err, "Erro ao Buscar Dados : ")
// 	}

// 	return users, nil
// }
