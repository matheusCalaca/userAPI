package routes

import (
	"net/http"
	user "userAPI/controlles/user"
	dbControler "userAPI/dao/db"

	"github.com/gin-gonic/gin"

	"userAPI/docs"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

// StartGin funçao para inicia a aplicação pelo gin
// @title Swagger Example API
// @version 1.0
// @description api para cadastro de pessoa
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email matheusfcalaca@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /v2
func StartGin() {

	docs.SwaggerInfo.Title = "User API"
	docs.SwaggerInfo.Description = "api para cadastro de pessoa"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8000"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

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
		api.GET("/users/:cpf", user.BuscaPessoaCpf)
		api.POST("/users", user.CreatePessoa)
		api.DELETE("/users/:cpf", user.DeletarPessoa)
	}
	// use ginSwagger middleware to
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
