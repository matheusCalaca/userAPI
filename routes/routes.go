package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func testeFunc(c *gin.Context) {
	fmt.Print("teste")
	log.Print("LOGGER TESTER")
	c.JSON(http.StatusOK, gin.H{"status": "success", "user": "{teste}"})
}

func StartGin() {
	router := gin.Default()
	api := router.Group("/api")
	{
		api.GET("/users", testeFunc)
	}

	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	router.Run(":8000")
}
