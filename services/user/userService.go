package pessoanegocio

import (
	"errors"
	"fmt"

	"strings"
	"time"
	pessoaDAO "userAPI/dao/user"
	models "userAPI/models/user"
	util "userAPI/services/util"

	"github.com/jinzhu/gorm"
)

//InsertPessoa metodo que faz a criaçao e validação das pessoas
func InsertPessoa(pessoa models.Pessoa, db *gorm.DB) (models.Pessoa, error) {

	_, errValidation := validaPessoa(&pessoa)
	if errValidation != nil {
		return pessoa, errValidation
	}

	//TODO: verifica se pessoa ja existe
	


	pessoaDB, err := pessoaDAO.InserirPessoaBD(pessoa, db)
	if err != nil {
		return pessoa, err
	}

	return pessoaDB, nil
}

// ListAllPessoa lista todas as pessoas do banco
func ListAllPessoa(db *gorm.DB) (*[]models.Pessoa, error) {
	pessoas, err := pessoaDAO.ListAllPessoa(db)
	if err != nil {
		return nil, err
	}
	fmt.Print(&pessoas)

	return pessoas, nil
}


// BuscaPessoa busca uma pessoa pelo cpf
func BuscaPessoa(cpf string, db *gorm.DB) (*models.Pessoa, error) {
	pessoas, err := pessoaDAO.BuscaPessoa(cpf , db)
	if err != nil {
		return nil, err
	}
	fmt.Print(&pessoas)

	return &pessoas, nil
}

// DeletarPessoaCpf Deletar Pessoa por ID
func DeletarPessoaCpf(cpf string, db *gorm.DB) (string, error) {
	msg, err := pessoaDAO.DeletarPessoaCpf(cpf, db)
	if err != nil {
		return "", err
	}
	fmt.Print(&msg)
	return msg, nil
}

//validaPessoa metodos com as regras de validação antes de inserir uma pessoa
func validaPessoa(pessoa *models.Pessoa) (bool, error) {
	if !util.IsCPFValido(pessoa.CPF) {
		return false, errors.New("CPF Invalido ! ")
	}

	if strings.Trim(pessoa.Nome, " ") == "" {
		return false, errors.New("Nome não pode ser vazio ! ")
	}

	if strings.Trim(pessoa.Sobrenome, " ") == "" {
		return false, errors.New("Sobrnome não pode ser vazio ! ")
	}

	if len(pessoa.Telefone) > 0 {
		for _, telefone := range pessoa.Telefone {
			isTelefoneValido, errorTelefone := util.IsTelefoneValido(telefone)
			if !isTelefoneValido {
				return false, errorTelefone
			}
		}
	}

	if pessoa.DataNascimento.UnixNano() >= time.Now().AddDate(0, 0, 1).UnixNano() {
		return false, errors.New("Data de Nascimento invalida ! ")
	}

	if util.IsEmailValido(pessoa.Email) {
		return false, errors.New("E-mail Invalido ! ")
	}

	for _, endereco := range pessoa.Endereco {
		isEnderecoValido, errorValidacaoEndereco := util.IsValidaEndereco(&endereco)
		if !isEnderecoValido {
			return false, errorValidacaoEndereco
		}
	}

	return true, nil
}
