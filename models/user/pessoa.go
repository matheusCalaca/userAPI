package model

import "time"

type Pessoa struct {
	ID_PESSOA      int64      `db:"ID_PESSOA" json:"id"`
	CPF            int64      `db:"CPF" json:"CPF"`
	Nome           string     `db:"NOME" json:"nome"`
	Sobrenome      string     `db:"SOBRENOME" json:"sobrenome"`
	RG             string     `db:"RG" json:"RG"`
	DataNascimento time.Time  `db:"DATA_NASCIMETO" json:"dataNascimento"`
	Email          string     `db:"EMAIL" json:"email"`
	Telefone       []Telefone ` json:"telefones"`
	Endereco       []Endereco ` json:"enderecos"`
	CreatedAt      int64      `db:"CREATED_AT" json:"created_at"`
	UpdatedAt      int64      `db:"UPDATE_AT" json:"updated_at"`
}
