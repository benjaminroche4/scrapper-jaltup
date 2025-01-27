package hellowork_test

import (
	"fmt"
	"scrapperjaltup/source/hellowork"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompanies(t *testing.T) {
	t.Parallel()

	api := hellowork.New()

	companies, err := api.Companies()
	assert.NoError(t, err)
	assert.NotEmpty(t, companies)

	fmt.Printf("nb = %d\n", len(companies))
	for _, company := range companies {
		fmt.Printf("%+v\n", company)
	}
}
