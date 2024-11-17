package model

type (
	Article struct {
		ID    string `json:"id"`
		Title string `json:"title"`
		Tag   string `json:tag`
	}
)
