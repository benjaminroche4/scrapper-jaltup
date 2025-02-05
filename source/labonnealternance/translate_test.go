package labonnealternance_test

import (
	"scrapperjaltup/model"
	"scrapperjaltup/source/labonnealternance"
	"testing"
	"time"

	"github.com/alex-cos/lbaapi"
	"github.com/stretchr/testify/assert"
)

var peJob = &lbaapi.PeJob{
	ID:            "184BFDV",
	Title:         "Salarié agricole (H/F)",
	IdeaType:      "peJob",
	URL:           "https://candidat.francetravail.fr/offres/recherche/detail/184BFDV",
	DetailsLoaded: false,
	Contact: lbaapi.Contact{
		Name:  "LETANG HERME SOURDUN - Mme PATRICIA SIMOND",
		Email: "Pour postuler: p.simond@etang-herme.fr",
		Phone: "0102030405",
		Info:  "Pour postuler: https://candidat.francetravail.fr/offres/recherche/detail/184BFDV",
		IV:    "",
	},
	Place: lbaapi.Place{
		Distance:          0,
		FullAddress:       "77 - HERME 77114",
		Latitude:          48.483968,
		Longitude:         3.34641,
		City:              "77 - HERME",
		Address:           "",
		Cedex:             "",
		ZipCode:           "77114",
		Insee:             "77227",
		DepartementNumber: "77",
		Region:            "Île-de-France",
		RemoteOnly:        false,
	},
	Company: lbaapi.Company{
		ID:            "5fd56d6d83a02f00081ef8b1",
		UAI:           "0623280D",
		Name:          "LETANG HERME SOURDUN",
		Siret:         "78362626000013",
		Size:          "",
		Logo:          "https://entreprise.francetravail.fr/static/img/logos/fT1mSwFTF51hAOPe3RI6Tb84RFTQBwLU.png",
		Description:   "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		SocialNetwork: "",
		URL:           "https://www.etang-herme.fr",
		Mandataire:    false,
		CreationDate:  "",
		Place:         lbaapi.Place{},
		Headquarter:   lbaapi.Headquarter{},
	},
	Job: lbaapi.Job{
		ID:                   "184BFDV",
		Description:          "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		EmployeurDescription: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
		CreationDate:         "2024-11-13T14:07:42.585Z",
		ContractType:         "cdd",
		ContractDescription:  "CDD - 24 Mois",
		Duration:             "35H",
		JobStartDate:         "2025-01-02T00:00:00.000Z",
		RythmeAlternance:     "",
		ElligibleHandicap:    false,
		DureeContrat:         "24 Mois",
		QuantiteContrat:      0,
	},
	Romes: []lbaapi.Rome{
		{
			Code: "A1101",
		},
	},
	Nafs: []lbaapi.Naf{},
}

func TestTranslatePlace(t *testing.T) {
	t.Parallel()

	out := labonnealternance.TranslatePlace(&peJob.Place)

	assert.Equal(t, model.Place{
		FullAddress: "77 - HERME 77114",
		City:        "Herme",
		ZipCode:     "77114",
		Latitude:    48.483968,
		Longitude:   3.34641,
	}, *out)
}

func TestTranslateOffer(t *testing.T) {
	t.Parallel()

	out := labonnealternance.TranslateOffer(peJob)

	assert.NotEmpty(t, out.PublicID)
	assert.NotEmpty(t, out.Company.PublicID)
	out.PublicID = ""
	out.Company.PublicID = ""

	assert.Equal(t, model.Offer{
		ID: 0,
		Company: model.Company{
			ID:           0,
			PublicID:     "",
			Name:         "Letang Herme Sourdun",
			Siret:        "78362626000013",
			ContactEmail: "p.simond@etang-herme.fr",
			PhoneNumber:  "0102030405",
			WebSiteURL:   "https://www.etang-herme.fr",
			Logo:         "https://entreprise.francetravail.fr/static/img/logos/fT1mSwFTF51hAOPe3RI6Tb84RFTQBwLU.png",
			CreatedAt:    time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
			Slug:         "letang-herme-sourdun",
			Verified:     false,
		},
		PublicID: "",
		Title:    "salarié agricole",
		Place: model.Place{
			FullAddress: "77 - HERME 77114",
			City:        "Herme",
			ZipCode:     "77114",
			Latitude:    48.483968,
			Longitude:   3.34641},
		Job: model.Job{
			Description:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			ContractType: "CDD",
			Duration:     0,
			Remote:       false,
			StudyLevel:   "",
			StartDate:    time.Date(2025, time.January, 2, 0, 0, 0, 0, time.UTC)},
		URL:          "https://candidat.francetravail.fr/offres/recherche/detail/184BFDV",
		Tag:          []string(nil),
		Status:       "published",
		CreatedAt:    time.Date(2024, time.November, 13, 14, 7, 42, 585000000, time.UTC),
		EndAt:        time.Date(2024, time.December, 13, 14, 7, 42, 585000000, time.UTC),
		EndPremiumAt: time.Date(2024, time.November, 20, 14, 7, 42, 585000000, time.UTC),
		Slug:         "salarie-agricole-h-f",
		Premium:      false,
		ExternalID:   "184BFDV",
		ServiceName:  "la-bonne-alternance",
		Categories:   []model.Category(nil),
	}, *out)
}

func TestTranslateOfferNoTitle(t *testing.T) {
	t.Parallel()

	job := *peJob
	job.Title = "  "
	out := labonnealternance.TranslateOffer(&job)

	assert.NotEmpty(t, out.PublicID)
	assert.NotEmpty(t, out.Company.PublicID)
	out.PublicID = ""
	out.Company.PublicID = ""

	assert.Equal(t, model.Offer{
		ID: 0,
		Company: model.Company{
			ID:           0,
			PublicID:     "",
			Name:         "Letang Herme Sourdun",
			Siret:        "78362626000013",
			ContactEmail: "p.simond@etang-herme.fr",
			PhoneNumber:  "0102030405",
			WebSiteURL:   "https://www.etang-herme.fr",
			Logo:         "https://entreprise.francetravail.fr/static/img/logos/fT1mSwFTF51hAOPe3RI6Tb84RFTQBwLU.png",
			CreatedAt:    time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC),
			Slug:         "letang-herme-sourdun",
			Verified:     false,
		},
		PublicID: "",
		Title:    "",
		Place: model.Place{
			FullAddress: "77 - HERME 77114",
			City:        "Herme",
			ZipCode:     "77114",
			Latitude:    48.483968,
			Longitude:   3.34641},
		Job: model.Job{
			Description:  "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			ContractType: "CDD",
			Duration:     0,
			Remote:       false,
			StudyLevel:   "",
			StartDate:    time.Date(2025, time.January, 2, 0, 0, 0, 0, time.UTC)},
		URL:          "https://candidat.francetravail.fr/offres/recherche/detail/184BFDV",
		Tag:          []string(nil),
		Status:       "archived",
		CreatedAt:    time.Date(2024, time.November, 13, 14, 7, 42, 585000000, time.UTC),
		EndAt:        time.Date(2024, time.December, 13, 14, 7, 42, 585000000, time.UTC),
		EndPremiumAt: time.Date(2024, time.November, 20, 14, 7, 42, 585000000, time.UTC),
		Slug:         "",
		Premium:      false,
		ExternalID:   "184BFDV",
		ServiceName:  "la-bonne-alternance",
		Categories:   []model.Category(nil),
	}, *out)
}
