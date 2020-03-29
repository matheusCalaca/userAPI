package pessoanegocio

import (
	"errors"
	"strconv"
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

	pessoaDB, err := pessoaDAO.InserirPessoaBD(pessoa, db)
	if err != nil {
		return pessoa, err
	}

	return pessoaDB, nil
}

//validaPessoa metodos com as regras de validação antes de inserir uma pessoa
func validaPessoa(pessoa *models.Pessoa) (bool, error) {
	if !util.IsCPFValido(strconv.FormatInt(pessoa.CPF, 10)) {
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
			isTelefoneValido, errorTelefone :=  util.IsTelefoneValido(telefone)
			if !isTelefoneValido {
				return false, errorTelefone
			}
		}
	}

	if pessoa.DataNascimento.UnixNano() >= time.Now().AddDate(0, 0, 1).UnixNano(){
return false, errors.New("Data de Nascimento invalida")
	}

	return true, nil
}

// func FindAllUser() (userModel.Users, error) {
// 	return userDao.SelectAllUser()
// }
