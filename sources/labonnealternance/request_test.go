package labonnealternance_test

import (
	"fmt"
	lba "scrapperjaltup/sources/labonnealternance"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJobFormations(t *testing.T) {
	t.Parallel()

	api := lba.New()

	romes := lba.GetRomeCodes()

	resp, err := api.JobFormations([]string{romes[0], romes[10], romes[20], romes[50], romes[100]})
	assert.NoError(t, err)

	fmt.Printf("resp = %+v\n", resp)
}
