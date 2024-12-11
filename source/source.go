package source

import "scrapperjaltup/model"

type Source interface {
	RetrieveOffers(func(int)) ([]*model.Offer, error)
	RetrieveCategories(func(int)) ([]model.Category, error)
}
