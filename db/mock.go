// nolint: misspell
package db

import (
	"scrapperjaltup/model"
	"time"

	"github.com/stretchr/testify/mock"
)

var categories = []model.Category{
	{
		ID:       1,
		PublicID: "KSCFEF7VUV",
		Name:     "Conception",
		Slug:     "conception",
	},
	{
		ID:       2,
		PublicID: "FMGL7TJYMW",
		Name:     "Production",
		Slug:     "production",
	},
	{
		ID:       3,
		PublicID: "BY6GDP1D5D",
		Name:     "Étude",
		Slug:     "etude",
	},
	{
		ID:       4,
		PublicID: "13F53S5GE3",
		Name:     "Marketing",
		Slug:     "marketing",
	},
	{
		ID:       5,
		PublicID: "ATE0K6G1S1",
		Name:     "Vente",
		Slug:     "vente",
	},
}

var companies = []model.Company{
	{
		ID:           1,
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
		ID:           1,
		PublicID:     "EW9AD5CDK0",
		Name:         "Coca-Cola",
		Siret:        "3344556677",
		ContactEmail: "contact@coca-cola.fr",
		PhoneNumber:  "0203040506",
		WebSiteURL:   "https://www.coca-cola.fr",
		Logo:         "https://www.coca-cola.fr/logo.png",
		CreatedAt:    time.Date(2021, 7, 12, 0, 0, 0, 0, time.UTC),
		Slug:         "coca-cola",
		Verified:     true,
	},
}

var offers = []model.Offer{
	{
		ID:       1,
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
		ID:       2,
		Company:  companies[1],
		PublicID: "HCY3A6OP2U",
		Title:    "Conseiller clientèle en assurances en alternance (H/F)",
		Place: model.Place{
			FullAddress: "PROPRIANO 20110",
			City:        "Propriano",
			ZipCode:     "20110",
			Latitude:    41.671527,
			Longitude:   8.906255,
		},
		Job: model.Job{
			Description:  "Lorem ipsum dolor sit amet, consectetuer adipiscing elit.",
			ContractType: "CDI",
			Duration:     10,
			Remote:       false,
			StudyLevel:   "BAC",
			StartDate:    time.Date(2024, 12, 8, 0, 0, 0, 0, time.UTC),
		},
		URL:         "https://candidat.francetravail.fr/offres/recherche/detail/184JTFT",
		Tag:         []string{},
		Status:      "published",
		CreatedAt:   time.Date(2024, 11, 2, 0, 0, 0, 0, time.UTC),
		EndAt:       time.Date(2025, 02, 2, 0, 0, 0, 0, time.UTC),
		Slug:        "conseiller-clientele-en-assurances-en-alternance-h-f",
		Premium:     false,
		ExternalID:  "184JTFT",
		ServiceName: "la-bonne-alternance",
		Categories:  []model.Category{categories[1], categories[4]},
	},
}

type Mock struct {
	mock.Mock
}

func NewMock() Database {
	m := &Mock{}

	m.On("Open").Return(nil)
	m.On("Close").Return(nil)
	m.On("Ping").Return(nil)

	m.On("CountCategories").Return(len(categories), nil)
	m.On("SelectCategories").Return(categories, nil)
	m.On("InsertCategories", mock.Anything).Return(nil)
	m.On("CleanCategories").Return(nil)

	m.On("CountCompanies").Return(len(companies), nil)
	m.On("SelectCompanies").Return(companies, nil)
	m.On("InsertCompanies", mock.Anything).Return(nil)
	m.On("CleanCompanies").Return(nil)

	m.On("CountOffers").Return(len(offers), nil)
	m.On("SelectOffers").Return(offers, nil)
	m.On("InsertOffers", mock.Anything).Return(nil)
	m.On("CleanOffers").Return(nil)

	return m
}

func (thiz *Mock) Open() error {
	args := thiz.Called()
	return args.Error(0)
}

func (thiz *Mock) Close() error {
	args := thiz.Called()
	return args.Error(0)
}

func (thiz *Mock) Ping() error {
	args := thiz.Called()
	return args.Error(0)
}

func (thiz *Mock) CountCategories() (int, error) {
	args := thiz.Called()
	return args.Get(0).(int), args.Error(1)
}

func (thiz *Mock) SelectCategories() ([]model.Category, error) {
	args := thiz.Called()
	return args.Get(0).([]model.Category), args.Error(1)
}

func (thiz *Mock) InsertCategories(categories []model.Category) error {
	args := thiz.Called(categories)
	return args.Error(0)
}

func (thiz *Mock) CleanCategories() error {
	args := thiz.Called()
	return args.Error(0)
}

func (thiz *Mock) CountCompanies() (int, error) {
	args := thiz.Called()
	return args.Get(0).(int), args.Error(1)
}

func (thiz *Mock) SelectCompanies() ([]model.Company, error) {
	args := thiz.Called()
	return args.Get(0).([]model.Company), args.Error(1)
}

func (thiz *Mock) InsertCompanies(companies []model.Company) error {
	args := thiz.Called(companies)
	return args.Error(0)
}

func (thiz *Mock) CleanCompanies() error {
	args := thiz.Called()
	return args.Error(0)
}

func (thiz *Mock) CountOffers() (int, error) {
	args := thiz.Called()
	return args.Get(0).(int), args.Error(1)
}

func (thiz *Mock) SelectOffers() ([]model.Offer, error) {
	args := thiz.Called()
	return args.Get(0).([]model.Offer), args.Error(1)
}

func (thiz *Mock) InsertOffers(offers []model.Offer) error {
	args := thiz.Called(offers)
	return args.Error(0)
}

func (thiz *Mock) CleanOffers() error {
	args := thiz.Called()
	return args.Error(0)
}
