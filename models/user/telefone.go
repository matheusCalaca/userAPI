package model

type Telefone struct {
	DDD       int64  `json:"ddd"`
	Telefone  int64  `json:"telefone"`
	Tipo      string `json:"tipo"`
	CreatedAt int64  `db:"CREATED_AT" json:"created_at"`
	UpdatedAt int64  `db:"UPDATE_AT" json:"updated_at"`
}
