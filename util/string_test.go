package util_test

import (
	"scrapperjaltup/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanCityName(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		input  string
		expect string
	}{
		{"VIC EN BIGORRE", "Vic En Bigorre"},
		{"89 - ST FARGEAU 89170", "St Fargeau"},
		{"75 - PARIS 09 75009", "Paris"},
		{"75 - PARIS 09", "Paris"},
		{"69 - Lyon 6e Arrondissement 69006", "Lyon"},
		{"69 - LYON 05 69005", "Lyon"},
		{"69", "69"},
	}

	for _, test := range tests {
		testname := test.input
		t.Run(testname, func(t *testing.T) {
			t.Parallel()
			output := util.CleanCityName(test.input)
			assert.Equal(t, test.expect, output)
		})
	}
}

func TestCleanTitle(t *testing.T) {
	t.Parallel()

	var tests = []struct {
		input  string
		expect string
	}{
		{"Conseiller de vente en alternance (H/F)", "conseiller de vente"},
		{"Conseiller de vente en alternance (f/h)", "conseiller de vente"},
		{"Manager d’unité marchande (H/F) en Apprentissage", "manager d’unité marchande"},
		{"Manager d’unité marchande (H-F)", "manager d’unité marchande"},
		{"Manager d’unité marchande ( H / F ) en Apprentissage", "manager d’unité marchande"},
		{"MANAGER D’UNITE MARCHANDE", "manager d’unite marchande"},
		{"Offre poste secrétaire dentaire en contrat d'alternance", "offre poste secrétaire dentaire"},
	}

	for _, test := range tests {
		testname := test.input
		t.Run(testname, func(t *testing.T) {
			t.Parallel()
			output := util.CleanTitle(test.input)
			assert.Equal(t, test.expect, output)
		})
	}
}
