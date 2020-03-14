package pessoaNegocio

import (
	"time"
	pessoaDAO "userAPI/dao/user"
	models "userAPI/models/user"
)

func InsertPessoa(pessoa models.Pessoa) (models.Pessoa, error) {

	// 	_, errValidation := validaInsertUser(user)
	// 	if errValidation != nil {
	// 		return user, errValidation
	// 	}

	pessoa.CreatedAt = time.Now()
	pessoa.UpdatedAt = time.Now()

	pessoaDB, err := pessoaDAO.InserirPessoaBD(pessoa)
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
