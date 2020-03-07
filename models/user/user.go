package model

import "time"

// User structure
type User struct {
	ID        int64     `db:"ID" json:"id"`
	Nome      string    `db:"NOME" json:"nome"`
	Endereco  string    `db:"ENDERECO" json:"endereco"`
	Idade     int       `db:"IDADE" json:"idade"`
	CreatedAt time.Time `db:"CREATED_AT" json:"created_at"`
	UpdatedAt time.Time `db:"UPDATE_AT" json:"updated_at"`
}

// Users list
type Users []User
