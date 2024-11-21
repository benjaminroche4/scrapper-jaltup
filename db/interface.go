package db

import "scrapperjaltup/model"

type Database interface {
	Open() error
	Close() error
	Ping() error

	CountCategories() (int, error)
	SelectCategories() ([]model.Category, error)
	InsertCategories([]model.Category) error
	CleanCategories() error

	CountCompanies() (int, error)
	SelectCompanies() ([]model.Company, error)
	InsertCompanies([]model.Company) error
	CleanCompanies() error

	CountOffers() (int, error)
	SelectOffers() ([]model.Offer, error)
	InsertOffers([]model.Offer) error
	CleanOffers() error
}
