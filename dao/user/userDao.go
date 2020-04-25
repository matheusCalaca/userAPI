package userdao

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	models "userAPI/models/user"

	"github.com/jinzhu/gorm"
)

// InserirPessoaBD insere a pessoa no banco de dados
func InserirPessoaBD(pessoa models.Pessoa, db *gorm.DB) (models.Pessoa, error) {
	caseUpperBeforeSave(&pessoa)
	result := db.Create(&pessoa)
	fmt.Print(&result.Error)
	if result.Error != nil {
		return pessoa, result.Error
	}
	return pessoa, nil

}

// DeletarPessoaCpf Deleta pessoa do banco de dados
func DeletarPessoaCpf(cpf string, db *gorm.DB) (string, error) {

	pessoa, err := BuscaPessoa(cpf, db)
	
	if err != nil {
		return "", errors.New("Erro ao localizar pessoa" + err.Error())
	}

	result := db.Delete(pessoa)
	fmt.Print(&result.Error)
	if result.Error != nil {
		return "", result.Error
	}
	return "Sucesso ao deletar pessoa: " + cpf, nil

}
// BuscaPessoa a pessoa pelo CPF
func BuscaPessoa(cpf string, db *gorm.DB) (models.Pessoa, error) {
	var pessoa models.Pessoa
	result := db.Where("CPF = ?", cpf).Find(&pessoa)

	return pessoa, result.Error

}

// ListAllPessoa lista todas as pessoas do banco de dados
func ListAllPessoa(db *gorm.DB) (*[]models.Pessoa, error) {

	var pessoas []models.Pessoa
	result := db.Find(&pessoas)
	if result.Error != nil {
		return nil, result.Error
	}

	for index, pessoa := range pessoas {
		var telefones []models.Telefone
		var enderecos []models.Endereco

		db.Find(&telefones, "TELEFONE_PESSOA = "+strconv.FormatUint(uint64(pessoa.ID), 16))
		db.Find(&enderecos, "ENDERECO_PESSOA = "+strconv.FormatUint(uint64(pessoa.ID), 16))

		pessoas[index].Telefone = telefones
		pessoas[index].Endereco = enderecos

	}

	fmt.Println(pessoas)
	return &pessoas, nil
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
