package model

import (
	"github.com/jinzhu/gorm"
)

type Telefone struct {
	gorm.Model
	DDD    int64  `gorm:"column:DDD" json:"ddd"`
	Numero int64  `gorm:"column:NUMERO" json:"numero"`
	Tipo   string `gorm:"column:TIPO" json:"tipo"`
}

func (Telefone) TableName() string {
	return "telefone"
}
