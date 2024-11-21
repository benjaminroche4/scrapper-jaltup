package model

type Category struct {
	ID       int64  `json:"id"`
	PublicID string `json:"publicId"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
}
