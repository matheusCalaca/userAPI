package model

import (
	"github.com/jinzhu/gorm"
)

// Telefone a a estrutura para o obejto telefone
type Telefone struct {
	gorm.Model
	DDD            int    `gorm:"column:DDD" json:"ddd"`
	Numero         string  `gorm:"column:NUMERO" json:"numero"`
	Tipo           string `gorm:"column:TIPO" json:"tipo"`
	TelefonePessoa uint   `gorm:"column:TELEFONE_PESSOA; foreignkey:TelefonePessoa" `
}

// TableName e afunção para definir o nome da tabela
func (Telefone) TableName() string {
	return "telefone"
}
