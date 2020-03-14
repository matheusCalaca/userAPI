package model

type Telefone struct {
	ID_TELEFONE int64  `db:"ID_TELEFONE" json:"id"`
	DDD         int64  `db:"DDD" json:"ddd"`
	Numero      int64  `db:"NUMERO" json:"numero"`
	Tipo        string `db:"TIPO" json:"tipo"`
	CreatedAt   int64  `db:"CREATED_AT" json:"created_at"`
	UpdatedAt   int64  `db:"UPDATE_AT" json:"updated_at"`
}
