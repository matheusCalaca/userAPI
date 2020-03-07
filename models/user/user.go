package model

// User structure
type User struct {
	ID        int64  `db:"ID" json:"id"`
	Nome      string `db:"NOME" json:"nome"`
	Endereco  string `db:"ENDERECO" json:"endereco"`
	Idade     int    `db:"IDADE" json:"idade"`
	CreatedAt int64  `db:"CREATED_AT" json:"created_at"`
	UpdatedAt int64  `db:"UPDATE_AT" json:"updated_at"`
}

// Users list
type Users []User
