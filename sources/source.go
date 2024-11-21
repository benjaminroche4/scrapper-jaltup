package sources

import "scrapperjaltup/model"

type Source interface {
	RetrieveOffers() ([]*model.Offer, error)
	RetrieveCategories() []model.Category
}
