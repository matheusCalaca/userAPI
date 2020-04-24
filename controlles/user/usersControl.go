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
// ListAccounts godoc
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Pessoa
// @Param pessoa body model.Pessoa true "Pessoa"
// @Router /users [post]
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

	c.JSON(http.StatusCreated, gin.H{"status": "success", "endereco": &pessoaCreate})
}

// ListarPessoas lista todas as pessoas
// ListAccounts godoc
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Pessoa
// @Router /users [get]
func ListarPessoas(c *gin.Context) {
	pessoasFind, err := pessoaNegocio.ListAllPessoa(c.MustGet("db").(*gorm.DB))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &pessoasFind)

}

// DeletarPessoa deleta pessoa por ID
// ListAccounts godoc
// @Accept  json
// @Produce  json
// @Param cpf path string true "Cpf para deletar uma Pessoa"
// @Success 200 {string} string "ok"
// @Router /users/{cpf} [delete]
func DeletarPessoa(c *gin.Context) {

	cpf := c.Param("cpf")
	msg, err := pessoaNegocio.DeletarPessoaID(cpf, c.MustGet("db").(*gorm.DB))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &msg)

}
