package model

import "github.com/jinzhu/gorm"

type Endereco struct {
	gorm.Model
	CEP         int64  `gorm:"column:CEP" json:"CEP"`
	Logradouro  string `gorm:"column:LOGRADOURO" json:"logradouro"`
	Bairro      string `gorm:"column:BAIRRO" json:"bairro"`
	Cidade      string `gorm:"column:CIDADE" json:"cidade"`
	UF          string `gorm:"column:UF" json:"UF"`
	Complemento string `gorm:"column:COMPLEMENTO" json:"complemento"`
	Numero      string `gorm:"column:NUMERO" json:"numero"`
	Tipo        string `gorm:"column:TIPO" json:"tipo"`
}


func (Endereco) TableName() string {
	return "ENDERECO"
}
