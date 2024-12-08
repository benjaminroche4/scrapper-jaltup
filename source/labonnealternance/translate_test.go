package labonnealternance_test

import (
	"fmt"
	"scrapperjaltup/model"
	"scrapperjaltup/source/labonnealternance"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslatePlace(t *testing.T) {
	t.Parallel()

	out := labonnealternance.TranslatePlace(&labonnealternance.Place{
		Distance:          0,
		FullAddress:       "59 - BORRE 59190",
		Latitude:          50.735966,
		Longitude:         2.580917,
		City:              "59 - BORRE ",
		Address:           "",
		Cedex:             "",
		ZipCode:           "59190",
		Insee:             "59091",
		DepartementNumber: "59",
		Region:            "",
		RemoteOnly:        false,
	})

	assert.Equal(t, model.Place{
		FullAddress: "59 - BORRE 59190",
		City:        "Borre",
		ZipCode:     "59190",
		Latitude:    50.735966,
		Longitude:   2.580917,
	}, *out)
}

func TestTranslateOffer(t *testing.T) {
	t.Parallel()

	out := labonnealternance.TranslateOffer(&labonnealternance.PeJob{
		ID:            "184BFDV",
		Title:         "Salarié agricole (H/F)",
		IdeaType:      "peJob",
		URL:           "https://candidat.francetravail.fr/offres/recherche/detail/184BFDV",
		DetailsLoaded: false,
		Contact: labonnealternance.Contact{
			Name:  "LETANG HERME SOURDUN - Mme PATRICIA SIMOND",
			Email: "Pour postuler: p.simond@etang-herme.fr",
			Phone: "0102030405",
			Info:  "Pour postuler: https://candidat.francetravail.fr/offres/recherche/detail/184BFDV",
			IV:    "",
		},
		Place: labonnealternance.Place{
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
		Company: labonnealternance.Company{
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
			Place:         labonnealternance.Place{},
			Headquarter:   labonnealternance.Headquarter{},
		},
		Job: labonnealternance.Job{
			ID:                   "184BFDV",
			Description:          "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			EmployeurDescription: "Lorem ipsum dolor sit amet, consectetur adipiscing elit.",
			CreationDate:         "2024-11-13T14:07:42.585Z",
			ContractType:         "CDD",
			ContractDescription:  "CDD - 24 Mois",
			Duration:             "35H",
			JobStartDate:         "2025-01-02T00:00:00.000Z",
			RythmeAlternance:     "",
			ElligibleHandicap:    false,
			DureeContrat:         "24 Mois",
			QuantiteContrat:      0,
		},
		Romes: []labonnealternance.Rome{
			{
				Code: "A1101",
			},
		},
		Nafs: []labonnealternance.Naf{
			{
				Label: "Autre mise à disposition de ressources humaines",
				Code:  "78",
			},
		},
	})

	fmt.Printf("out = %+v\n", out)
}
