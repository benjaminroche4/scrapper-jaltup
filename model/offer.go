package model

import "time"

type Offer struct {
	ID          int64      `json:"id"`
	Company     Company    `json:"company,omitempty"`
	PublicID    string     `json:"publicId"`
	Title       string     `json:"title"`
	Place       Place      `json:"place"`
	Job         Job        `json:"job"`
	URL         string     `json:"url,omitempty"`
	Tag         []string   `json:"tag,omitempty"`
	Status      string     `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	EndAt       time.Time  `json:"endAt"`
	Slug        string     `json:"slug"`
	Premium     bool       `json:"premium"`
	ExternalID  string     `json:"externalId"`
	ServiceName string     `json:"serviceName"`
	Categories  []Category `json:"categories,omitempty"`
}
