package controllers

import (
	"errors"
	"log"

	models "userAPI/models/user"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func InitDb() *gorm.DB {
	var pessoa = &models.Pessoa{}
	var telefone = &models.Telefone{}
	var endereco = &models.Endereco{}
	dc, err := gorm.Open("mysql", "root:1234@tcp(localhost:3306)/user?charset=utf8&parseTime=True&loc=Local")
	CheckErr(err, "Falha ao iniciar SQL")

	dc.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	dc.AutoMigrate(&pessoa, &telefone, &endereco)
	dc.LogMode(true)
	CheckErr(err, "Falha a criar Tabelas")
	return dc
}

func CheckErr(err error, msg string) error {
	if err != nil {
		log.Print(msg + " (" + err.Error() + ")")
		return errors.New(msg + " (" + err.Error() + ")")
	}
	return nil
}

// func Cors() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
// 		c.Next()
// 	}
// }
