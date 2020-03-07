package user

import (
	"errors"
	"log"
	"net/http"
	"time"

	dbControllers "userAPI/controlles/db"
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

// var users = userModel.Users
var dbmap = dbControllers.InitDb()

func ListAllUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success", "user": "teste"})
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
	log.Println(userJson)

	insert, errInsert := dbmap.Exec(`INSERT INTO user (NOME, ENDERECO, IDADE, CREATED_AT) VALUES (?, ?, ?, ?)`, userJson.Nome, userJson.Endereco, userJson.Idade, userJson.CreatedAt)
	if insert != nil {
		user_id, err := insert.LastInsertId()

		if err == nil {
			content := &userModel.User{
				ID:        user_id,
				Nome:      userJson.Nome,
				Endereco:  userJson.Endereco,
				Idade:     userJson.Idade,
				CreatedAt: userJson.CreatedAt,
			}
			c.JSON(201, content)
		} else {
			dbControllers.CheckErr(err, "Falha ao inserir")
		}
	}
	if errInsert != nil {
		dbControllers.CheckErr(errInsert, "Error ao inserir USER")
	}

	// c.JSON(http.StatusOK, gin.H{"status": "success", "user": &userJson})
}
