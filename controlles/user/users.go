package user

import (
	"errors"
	"net/http"

	dbControllers "userAPI/dao/db"
	userModel "userAPI/models/user"
	userNegocio "userAPI/services/user"

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

var dbmap = dbControllers.InitDb()

func ListAllUser(c *gin.Context) {
	users, err := userNegocio.FindAllUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, &users)
}

func CreateUser(c *gin.Context) {
	var userJson userModel.User
	err := c.Bind(&userJson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": errInvalidBody.Error()})
		return
	}

	userCreate, err := userNegocio.InsertUser(userJson)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "failed", "message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "user": &userCreate})
}
