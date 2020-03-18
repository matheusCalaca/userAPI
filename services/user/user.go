package pessoaNegocio

import (
	pessoaDAO "userAPI/dao/user"
	models "userAPI/models/user"

	"github.com/jinzhu/gorm"
)

func InsertPessoa(pessoa models.Pessoa, db *gorm.DB) (models.Pessoa, error) {

	// 	_, errValidation := validaInsertUser(user)
	// 	if errValidation != nil {
	// 		return user, errValidation
	// 	}

	pessoaDB, err := pessoaDAO.InserirPessoaBD(pessoa, db)
	if err != nil {
		return pessoa, err
	}

	return pessoaDB, nil
}

// func validaInsertUser(userVariavel userModel.User) (bool, error) {
// 	if len(userVariavel.CpfCnpj) <= 0 {
// 		return false, errors.New("CPF CNPJ não pode ser vazio!")
// 	}
// 	return true, nil
// }

// func FindAllUser() (userModel.Users, error) {
// 	return userDao.SelectAllUser()
// }
