package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Pessoa struct {
	gorm.Model
	CPF            int64      `gorm:"column:CPF" json:"CPF"`
	Nome           string     `gorm:"column:NOME" json:"nome"`
	Sobrenome      string     `gorm:"column:SOBRENOME" json:"sobrenome"`
	RG             string     `gorm:"column:RG" json:"RG"`
	DataNascimento time.Time  `gorm:"column:DATA_NASCIMETO" json:"dataNascimento"`
	Email          string     `gorm:"column:EMAIL" json:"email"`
	// Telefone       []Telefone ` json:"telefones"`
	// Endereco       []Endereco ` json:"enderecos"`
}


func (Pessoa) TableName() string {
	return "PESSOA"
}
