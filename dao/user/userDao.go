package userdao

import (
	"log"
	"strings"
	models "userAPI/models/user"

	"github.com/jinzhu/gorm"
)

// InserirPessoaBD insere a pessoa no banco de dados
func InserirPessoaBD(pessoa models.Pessoa, db *gorm.DB) (models.Pessoa, error) {
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
