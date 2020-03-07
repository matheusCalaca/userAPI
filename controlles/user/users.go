package user

import (
	"errors"
	"log"
	"net/http"
	"time"

	userModel "userAPI/models/user"

	"github.com/gin-gonic/gin"
)

const UserCollection = "user"

var (
	errNotExist        = errors.New("User não existe")
	errInvalidID       = errors.New("ID Invalido")
	errInvalidBody     = errors.New("Corpo Json Invalido")
	errInsertionFailed = errors.New("Erro ao Inserir usuario")
	errUpdationFailed  = errors.New("Erro ao Atualizar usuario")
	errDeletionFailed  = errors.New("Erro ao Deletar usuario")
)

var users []userModel.User = userModel.Users{
	{
		Nome:     "Matheus Calaça",
		Endereço: "Endereço matheus calaça",
		Idade:    23,
	},
	{
		Nome:     "Ana julia calaça",
		Endereço: "Endereço ana julia",
		Idade:    15,
	},
	{
		Nome:     "Juliana Calaça",
		Endereço: "Endereço da Juliana",
		Idade:    15,
	},
}

func ListAllUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success", "user": &users})
}

func CreateUser(c *gin.Context) {
	var userJson userModel.User
	err := c.Bind(&userJson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
		return
	}
	userJson.CreatedAt = time.Now()
	userJson.UpdatedAt = time.Now()
	log.Printf("[ LOGGER ]", userJson)

	c.JSON(http.StatusOK, gin.H{"status": "success", "user": &userJson})
}
