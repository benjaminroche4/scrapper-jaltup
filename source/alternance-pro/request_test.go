package alternancepro_test

import (
	"fmt"
	"scrapperjaltup/model"
	ap "scrapperjaltup/source/alternance-pro"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompanies(t *testing.T) {
	t.Parallel()

	api := ap.New()

	companies, err := api.Companies()
	assert.NoError(t, err)
	assert.NotEmpty(t, companies)

	fmt.Printf("nb = %d\n", len(companies))
	for _, company := range companies {
		fmt.Printf("%+v\n", company)
	}
}

func TestJobs(t *testing.T) {
	t.Parallel()

	var allOffers []*model.Offer

	api := ap.New()

	page := 1
	for {
		offers, maxPage, err := api.Jobs(page, "")
		assert.NoError(t, err)
		assert.NotEmpty(t, offers)
		allOffers = append(allOffers, offers...)
		if page >= maxPage {
			break
		}
		page++
	}

	fmt.Printf("nb = %d\n", len(allOffers))
	for _, offer := range allOffers {
		fmt.Printf("%+v\n", offer)
	}
}
