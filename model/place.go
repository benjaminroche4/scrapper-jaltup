package model

type Place struct {
	FullAddress string  `json:"fullAddress,omitempty"`
	City        string  `json:"city,omitempty"`
	ZipCode     string  `json:"zipCode,omitempty"`
	Latitude    float64 `json:"latitude,omitempty"`
	Longitude   float64 `json:"longitude,omitempty"`
}
