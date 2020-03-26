package pessoaNegocio

import (
	"errors"
	"strconv"
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
	if util.IsCPFValido(strconv.FormatInt(pessoa.CPF, 10)) {
		return false, errors.New("CPF CNPJ não pode ser vazio!")
	}
	return true, nil
}

// func FindAllUser() (userModel.Users, error) {
// 	return userDao.SelectAllUser()
// }
