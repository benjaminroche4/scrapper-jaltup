package matcher

import (
	"log"
	_db "scrapperjaltup/db"
	"scrapperjaltup/model"
	"scrapperjaltup/source"
)

type Matcher struct {
	db     _db.Database
	source source.Source
}

func New(db _db.Database, source source.Source) *Matcher {
	return &Matcher{
		db:     db,
		source: source,
	}
}

func (thiz *Matcher) Execute() error {
	log.Println("[Matcher]: Execute")

	err := thiz.syncCategories()
	if err != nil {
		return err
	}

	categories, err := thiz.db.SelectCategories()
	if err != nil {
		return err
	}

	log.Printf("[Matcher]: Retrieved from db %d existing categories\n", len(categories))

	companies, err := thiz.db.SelectCompanies()
	if err != nil {
		return err
	}

	log.Printf("[Matcher]: Retrieved from db %d existing companies\n", len(companies))

	existing, err := thiz.db.SelectOffers()
	if err != nil {
		return err
	}

	log.Printf("[Matcher]: Retrieved from db %d existing offers\n", len(existing))

	offers, err := thiz.source.RetrieveOffers()
	if err != nil {
		return err
	}

	log.Printf("[Matcher]: Fetched from source %d offers\n", len(offers))

	newOffers := []model.Offer{}
	for i := range offers {
		offer := offers[i]

		if ok, _ := isOfferInList(offer, existing); ok {
			continue
		}

		if ok, company := isCompanyInList(&offer.Company, companies); ok && company != nil {
			offer.Company.ID = company.ID
		} else {
			newCompany := []model.Company{offer.Company}
			err = thiz.db.InsertCompanies(newCompany)
			if err == nil {
				companies = append(companies, newCompany[0])
				offer.Company = newCompany[0]
			}
		}

		for i := range offer.Categories {
			if ok, category := isCategoryInList(&offer.Categories[i], categories); ok && category != nil {
				offer.Categories[i] = *category
			}
		}

		newOffers = append(newOffers, *offer)
	}

	err = thiz.db.InsertOffers(newOffers)
	if err != nil {
		return err
	}

	log.Printf("[Matcher]: Inserted %d new offers\n", len(newOffers))

	return nil
}

// Unexported functions

func (thiz *Matcher) syncCategories() error {
	existing, err := thiz.db.SelectCategories()
	if err != nil {
		return err
	}
	categories, err := thiz.source.RetrieveCategories()
	if err != nil {
		return err
	}
	newCategories := []model.Category{}
	for _, category := range categories {
		if ok, _ := isCategoryInList(&category, existing); ok {
			continue
		}
		newCategories = append(newCategories, category)
	}

	err = thiz.db.InsertCategories(newCategories)
	if err != nil {
		return err
	}

	log.Printf("[Matcher]: Inserted %d new categrories\n", len(newCategories))

	return nil
}

func isOfferInList(offer *model.Offer, list []model.Offer) (bool, *model.Offer) {
	for i := range list {
		if (offer.ServiceName == list[i].ServiceName) &&
			(offer.ExternalID == list[i].ExternalID) {
			return true, &list[i]
		}
	}

	return false, nil
}

func isCompanyInList(company *model.Company, list []model.Company) (bool, *model.Company) {
	for i := range list {
		if (company.Siret != "" && company.Siret == list[i].Siret) ||
			(company.Slug != "" && company.Slug == list[i].Slug) {
			return true, &list[i]
		}
	}

	return false, nil
}

func isCategoryInList(category *model.Category, list []model.Category) (bool, *model.Category) {
	for i := range list {
		if category.Slug == list[i].Slug {
			return true, &list[i]
		}
	}

	return false, nil
}
