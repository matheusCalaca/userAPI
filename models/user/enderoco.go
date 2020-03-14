package model

type Endereco struct {
	ID_ENDEREOCO int64  `db:"ID_TELEFONE" json:"id"`
	CEP          int64  `db:"CEP" json:"CEP"`
	Logradouro   string `db:"LOGRADOURO" json:"logradouro"`
	Bairro       string `db:"BAIRRO" json:"bairro"`
	Cidade       string `db:"CIDADE" json:"cidade"`
	UF           string `db:"UF" json:"UF"`
	Complemento  string `db:"COMPLEMENTO" json:"complemento"`
	Numero       string `db:"NUMERO" json:"numero"`
	Tipo         string `db:"TIPO" json:"tipo"`
	CreatedAt    int64  `db:"CREATED_AT" json:"created_at"`
	UpdatedAt    int64  `db:"UPDATE_AT" json:"updated_at"`
}
