package user

import (
	"errors"
	"net/http"

	models "userAPI/models/user"
	pessoaNegocio "userAPI/services/user"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)


var (
	errNotExist        = errors.New("User não existe")
	errInvalidID       = errors.New("ID Invalido")
	errInvalidBody     = errors.New("Corpo Json Invalido")
	errInsertionFailed = errors.New("Erro ao Inserir usuario")
	errUpdationFailed  = errors.New("Erro ao Atualizar usuario")
	errDeletionFailed  = errors.New("Erro ao Deletar usuario")
)

// CreatePessoa cria uma pessoa e retorna
func CreatePessoa(c *gin.Context) {
	// pessoaJSON objetos com os dadios de pessoa que recebe do Json
	var pessoaJSON models.Pessoa

	err := c.Bind(&pessoaJSON)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
		return
	}

	pessoaCreate, err := pessoaNegocio.InsertPessoa(pessoaJSON, c.MustGet("db").(*gorm.DB))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "endereco": &pessoaCreate})
}
