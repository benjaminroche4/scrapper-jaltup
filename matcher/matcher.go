package matcher

import (
	"crypto/rand"
	"log"
	"math/big"
	_db "scrapperjaltup/db"
	"scrapperjaltup/model"
	"scrapperjaltup/source"

	"github.com/liamg/loading/pkg/bar"
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

	finish, setProgression := loadingBar()
	offers, err := thiz.source.RetrieveOffers(setProgression)
	if err != nil {
		return err
	}
	finish()

	log.Printf("[Matcher]: Fetched from source %d offers\n", len(offers))

	newOffers := []model.Offer{}
	for i := range offers {
		offer := offers[i]

		if isOfferInList(offer, existing) {
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

	err = thiz.db.InsertOffers(markRandomPremium(shuffleOffers(newOffers)))
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

	finish, setProgression := loadingBar()
	categories, err := thiz.source.RetrieveCategories(setProgression)
	if err != nil {
		return err
	}
	finish()

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

func isOfferInList(offer *model.Offer, list []model.Offer) bool {
	for i := range list {
		if model.IsSame(offer, &list[i]) {
			return true
		}
	}

	return false
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

func loadingBar() (func(), func(int)) {
	loadingBar := bar.New(
		bar.OptionHideOnFinish(true),
	)
	loadingBar.SetTotal(100)
	loadingBar.SetLabel("Fetching")
	loadingBar.SetCurrent(0)

	return func() {
			loadingBar.Finish()
		}, func(p int) {
			loadingBar.SetCurrent(p)
		}
}

func shuffleOffers(offers []model.Offer) []model.Offer {
	output := []model.Offer{}

	for {
		if len(offers) == 0 {
			return output
		}
		maximum := big.NewInt(int64(len(offers)))
		random, _ := rand.Int(rand.Reader, maximum)
		index := random.Int64()
		output = append(output, offers[index])
		offers = append(offers[:index], offers[index+1:]...)
	}
}

func markRandomPremium(offers []model.Offer) []model.Offer {
	output := []model.Offer{}

	for i := range offers {
		random, _ := rand.Int(rand.Reader, big.NewInt(int64(100)))
		if random.Int64() <= 10 {
			offers[i].Premium = true
		}
		output = append(output, offers[i])
	}

	return output
}
