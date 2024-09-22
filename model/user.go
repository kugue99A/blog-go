package model

import "time"

type (
	User struct {
		ID        string    `json:"id"`
		Name      string    `json:"name"`
		Birthday  time.Time `json:"birthday"`
		Tel       string    `json:"tel"`
		PostCode  string    `json:"post_code"`
		Address   string    `json:"address"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}
)
