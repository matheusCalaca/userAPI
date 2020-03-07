package user

import (
	"net/http"

	user "userAPI/models/user"

	"github.com/gin-gonic/gin"
)

const UserCollection = "user"

// var (
// 	errNotExist        = errors.New("Users are not exist")
// 	errInvalidID       = errors.New("Invalid ID")
// 	errInvalidBody     = errors.New("Invalid request body")
// 	errInsertionFailed = errors.New("Error in the user insertion")
// 	errUpdationFailed  = errors.New("Error in the user updation")
// 	errDeletionFailed  = errors.New("Error in the user deletion")
// )
var users = user.Users{
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

func ListAll(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success", "user": &users})
}
