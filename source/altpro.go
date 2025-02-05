package source

import (
	"scrapperjaltup/model"
	alternancepro "scrapperjaltup/source/alternance-pro"
)

type AltPro struct {
	api *alternancepro.AlternancePro
}

func NewAltPro() Source {
	return &AltPro{
		api: alternancepro.New(),
	}
}

func (thiz *AltPro) RetrieveOffers(setProgression func(int)) ([]*model.Offer, error) {
	var offers []*model.Offer

	setProgression(1)
	companies, err := thiz.api.Companies()
	setProgression(10)
	if err != nil {
		return nil, err
	}

	page := 1
	for {
		results, maxPage, err := thiz.api.Jobs(page, "")
		setProgression(10 + page/maxPage*80)
		if err != nil {
			return nil, err
		}
		offers = append(offers, results...)
		if page >= maxPage {
			break
		}
		page++
	}

	setProgression(90)
	for i := range offers {
		slug := offers[i].Company.Slug
		company := findCompanyBySlug(slug, companies)
		if company != nil {
			offers[i].Company = *company
		}
	}
	setProgression(100)

	return offers, nil
}

func (thiz *AltPro) RetrieveCategories(setProgression func(int)) ([]model.Category, error) {
	return []model.Category{}, nil
}

// Unexported function

func findCompanyBySlug(slug string, companies []*model.Company) *model.Company {
	for _, company := range companies {
		if company.Slug == slug {
			return company
		}
	}

	return nil
}
