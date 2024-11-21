// nolint: misspell
package sources

import (
	"scrapperjaltup/model"
	"time"

	"github.com/stretchr/testify/mock"
)

var categories = []model.Category{
	{
		ID:       0,
		PublicID: "KSCFEF7VUV",
		Name:     "Conception",
		Slug:     "conception",
	},
	{
		ID:       0,
		PublicID: "FMGL7TJYMW",
		Name:     "Production",
		Slug:     "production",
	},
	{
		ID:       0,
		PublicID: "BY6GDP1D5D",
		Name:     "Étude",
		Slug:     "etude",
	},
	{
		ID:       0,
		PublicID: "13F53S5GE3",
		Name:     "Marketing",
		Slug:     "marketing",
	},
	{
		ID:       0,
		PublicID: "ATE0K6G1S1",
		Name:     "Vente",
		Slug:     "vente",
	},
	{
		ID:       0,
		PublicID: "WCAPKE32QX",
		Name:     "Support",
		Slug:     "support",
	},
	{
		ID:       0,
		PublicID: "CW3M6CZZP6",
		Name:     "Achat",
		Slug:     "achat",
	},
}

var companies = []model.Company{
	{
		ID:           0,
		PublicID:     "EW9AD5CDK0",
		Name:         "Carrefour",
		Siret:        "1234567890",
		ContactEmail: "contact@carrefour.fr",
		PhoneNumber:  "0102030405",
		WebSiteURL:   "https://www.carrefour.fr",
		Logo:         "https://www.carrefour.fr/logo.png",
		CreatedAt:    time.Date(2022, 12, 12, 0, 0, 0, 0, time.UTC),
		Slug:         "carrefour",
		Verified:     true,
	},
	{
		ID:           0,
		PublicID:     "CL6XGZWB95",
		Name:         "La Poste",
		Siret:        "9988776655",
		ContactEmail: "contact@la-poste.fr",
		PhoneNumber:  "0304050607",
		WebSiteURL:   "https://www.la-poste.fr",
		Logo:         "https://www.la-poste.fr/logo.png",
		CreatedAt:    time.Date(2018, 1, 4, 0, 0, 0, 0, time.UTC),
		Slug:         "la-poste",
		Verified:     false,
	},
}

var offers = []*model.Offer{
	{
		ID:       0,
		Company:  companies[0],
		PublicID: "KT9BXAK2NK",
		Title:    "Assistant / Assistante de communication en Alternance (H/F)",
		Place: model.Place{
			FullAddress: "DIJON 21000",
			City:        "Dijon",
			ZipCode:     "21000",
			Latitude:    47.331848,
			Longitude:   5.033601,
		},
		Job: model.Job{
			Description:  "Lorem ipsum dolor sit amet, consectetuer adipiscing elit.",
			ContractType: "CDD",
			Duration:     8,
			Remote:       false,
			StudyLevel:   "BAC+2",
			StartDate:    time.Date(2024, 12, 8, 0, 0, 0, 0, time.UTC),
		},
		URL:         "https://candidat.francetravail.fr/offres/recherche/detail/184LMCH",
		Tag:         []string{},
		Status:      "published",
		CreatedAt:   time.Date(2024, 12, 8, 0, 0, 0, 0, time.UTC),
		EndAt:       time.Date(2025, 03, 8, 0, 0, 0, 0, time.UTC),
		Slug:        "assistant-assistante-de-communication-en-alternance-h-f",
		Premium:     false,
		ExternalID:  "184LMCH",
		ServiceName: "la-bonne-alternance",
		Categories:  []model.Category{categories[3], categories[4]},
	},
	{
		ID:       0,
		Company:  companies[1],
		PublicID: "AIZSQVFRGY",
		Title:    "Alternance Facteur (H/F)",
		Place: model.Place{
			FullAddress: "BOIS GUILLAUME 76230",
			City:        "Bois Guillaume",
			ZipCode:     "76230",
			Latitude:    49.461399,
			Longitude:   1.1092,
		},
		Job: model.Job{
			Description:  "Lorem ipsum dolor sit amet, consectetuer adipiscing elit.",
			ContractType: "CDD",
			Duration:     9,
			Remote:       false,
			StudyLevel:   "BAC+3",
			StartDate:    time.Date(2024, 12, 8, 0, 0, 0, 0, time.UTC),
		},
		URL:         "https://candidat.francetravail.fr/offres/recherche/detail/8870347",
		Tag:         []string{},
		Status:      "published",
		CreatedAt:   time.Date(2024, 12, 8, 0, 0, 0, 0, time.UTC),
		EndAt:       time.Date(2025, 03, 8, 0, 0, 0, 0, time.UTC),
		Slug:        "alternance-facteur-h-f",
		Premium:     false,
		ExternalID:  "8870347",
		ServiceName: "la-bonne-alternance",
		Categories:  []model.Category{categories[5]},
	},
}

type Mock struct {
	mock.Mock
}

func NewMock() Source {
	m := &Mock{}

	m.On("RetrieveOffers").Return(offers, nil)
	m.On("RetrieveCategories").Return(categories)

	return m
}

func (thiz *Mock) RetrieveOffers() ([]*model.Offer, error) {
	args := thiz.Called()
	return args.Get(0).([]*model.Offer), args.Error(1)
}

func (thiz *Mock) RetrieveCategories() []model.Category {
	args := thiz.Called()
	return args.Get(0).([]model.Category)
}
