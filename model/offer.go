package model

import "time"

type Offer struct {
	ID           int64      `json:"id"`
	Company      Company    `json:"company,omitempty"`
	PublicID     string     `json:"publicId"`
	Title        string     `json:"title"`
	Place        Place      `json:"place"`
	Job          Job        `json:"job"`
	URL          string     `json:"url,omitempty"`
	Tag          []string   `json:"tag,omitempty"`
	Status       string     `json:"status"`
	CreatedAt    time.Time  `json:"createdAt"`
	EndAt        time.Time  `json:"endAt"`
	EndPremiumAt time.Time  `json:"endPremiumAt"`
	Slug         string     `json:"slug"`
	Premium      bool       `json:"premium"`
	ExternalID   string     `json:"externalId"`
	ServiceName  string     `json:"serviceName"`
	Categories   []Category `json:"categories,omitempty"`
}

func IsSame(offer1, offer2 *Offer) bool {
	return ((offer1.ServiceName == offer2.ServiceName) &&
		(offer1.ExternalID == offer2.ExternalID)) ||
		((offer1.Title != "") && (offer2.Title != "") &&
			(offer1.Title == offer2.Title) &&
			(offer1.Job.Description == offer2.Job.Description))
}
