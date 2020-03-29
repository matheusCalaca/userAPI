package model

import "github.com/jinzhu/gorm"

// Endereco e a strutura de objeto para endereço
type Endereco struct {
	gorm.Model
	CEP            string  `gorm:"column:CEP" json:"CEP"`
	Logradouro     string `gorm:"column:LOGRADOURO" json:"logradouro"`
	Bairro         string `gorm:"column:BAIRRO" json:"bairro"`
	Cidade         string `gorm:"column:CIDADE" json:"cidade"`
	UF             string `gorm:"column:UF" json:"UF"`
	Complemento    string `gorm:"column:COMPLEMENTO" json:"complemento"`
	Numero         string `gorm:"column:NUMERO" json:"numero"`
	Tipo           string `gorm:"column:TIPO" json:"tipo"`
	EnderecoPessoa uint   `gorm:"column:ENDERECO_PESSOA"`
}

// TableName e afunção para definir o nome da tabela
func (Endereco) TableName() string {
	return "ENDERECO"
}
