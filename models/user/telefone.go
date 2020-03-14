package model

type Telefone struct {
	DDD       int64  `db:"DDD" json:"ddd"`
	Numero    int64  `db:"NUMERO" json:"numero"`
	Tipo      string `db:"ITPO" json:"tipo"`
	CreatedAt int64  `db:"CREATED_AT" json:"created_at"`
	UpdatedAt int64  `db:"UPDATE_AT" json:"updated_at"`
}
