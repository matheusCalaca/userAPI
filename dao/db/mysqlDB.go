package controllers

import (
	"database/sql"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

// type Dbmap = InitDb()

func InitDb() *gorp.DbMap {
	db, err := sql.Open("mysql", "root:1234@tcp(localhost:3306)/user")
	CheckErr(err, "Falha ao iniciar SQL")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	err = dbmap.CreateTablesIfNotExists()
	CheckErr(err, "Falha a criar Tabelas")

	return dbmap
}

func CheckErr(err error, msg string) error {
	if err != nil {
		log.Print(msg + " (" + err.Error() + ")")
		return errors.New(msg + " (" + err.Error() + ")")
	}
	return nil
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}
