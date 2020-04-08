package routes

import (
	"net/http"
	user "userAPI/controlles/user"
	dbControler "userAPI/dao/db"

	"github.com/gin-gonic/gin"
)

// StartGin funçao para inicia a aplicação pelo gin
func StartGin() {
	router := gin.Default()

	// connnection database
	db := dbControler.InitDb()
	// provider DB to controller
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})
	router.Use(CORSMiddleware())

	api := router.Group("/api")
	{
		api.GET("/users", user.ListarPessoas)
		api.POST("/users", user.CreatePessoa)
	}

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	router.Run(":8000")
}

// CORSMiddleware remove o erro de CORS
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
