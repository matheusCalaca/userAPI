package user

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// User structure
type User struct {
	// ID        bson.ObjectId `bson:"_id"`
	Nome      string    `bson:"nome"`
	Endereço  string    `bson:"endereço"`
	idade     int       `bson:"idade"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

const UserCollection = "user"

// var (
// 	errNotExist        = errors.New("Users are not exist")
// 	errInvalidID       = errors.New("Invalid ID")
// 	errInvalidBody     = errors.New("Invalid request body")
// 	errInsertionFailed = errors.New("Error in the user insertion")
// 	errUpdationFailed  = errors.New("Error in the user updation")
// 	errDeletionFailed  = errors.New("Error in the user deletion")
// )

// Users list
var users = []User{
	{
		Nome:     "Matheus Calaça",
		Endereço: "Endereço matheus calaça",
		idade:    24,
	},
	{
		Nome:     "Ana julia calaça",
		Endereço: "Endereço ana julia",
		idade:    15,
	},
	{
		Nome:     "Juliana Calaça",
		Endereço: "Endereço da Juliana",
		idade:    15,
	},
}

func ListAll(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"status": "success", "user": &users})
}
