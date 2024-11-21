package model

import "time"

type Company struct {
	ID           int64     `json:"id"`
	PublicID     string    `json:"publicId"`
	Name         string    `json:"name"`
	Siret        string    `json:"siret,omitempty"`
	ContactEmail string    `json:"contactEmail,omitempty"`
	PhoneNumber  string    `json:"phoneNumber,omitempty"`
	WebSiteURL   string    `json:"webSiteURL,omitempty"`
	Logo         string    `json:"logo,omitempty"`
	CreatedAt    time.Time `json:"createdAt"`
	Slug         string    `json:"slug"`
	Verified     bool      `json:"verified"`
}
