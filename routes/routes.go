package routes

import (
	"net/http"
	user "userAPI/controlles/user"
	dbControler "userAPI/dao/db"

	"github.com/gin-gonic/gin"
)

func StartGin() {
	router := gin.Default()

	// connnection database
	db := dbControler.InitDb()
	// provider DB to controller
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

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
