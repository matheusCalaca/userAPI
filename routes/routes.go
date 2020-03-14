package routes

import (
	"net/http"
	user "userAPI/controlles/user"

	"github.com/gin-gonic/gin"
)

func StartGin() {
	router := gin.Default()
	api := router.Group("/api")
	{
		// api.GET("/users", user.ListAllUser)
		api.POST("/users", user.CreatePessoa)
	}

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	router.Run(":8000")
}
