package controllers

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

// type Dbmap = InitDb()

func InitDb() *gorp.DbMap {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/golang")
	CheckErr(err, "Falha ao iniciar SQL")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	err = dbmap.CreateTablesIfNotExists()
	CheckErr(err, "Falha a criar Tabelas")

	return dbmap
}

func CheckErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
