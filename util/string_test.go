package util_test

import (
	"scrapperjaltup/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCleanTitle(t *testing.T) {
	t.Parallel()

	out := util.CleanTitle("Conseiller de vente en alternance (H/F)")

	assert.Equal(t, "conseiller de vente", out)

	out = util.CleanTitle("Conseiller de vente en alternance (f/h)")

	assert.Equal(t, "conseiller de vente", out)

	out = util.CleanTitle("Manager d’unité marchande (H/F) en Apprentissage")

	assert.Equal(t, "manager d’unité marchande", out)

	out = util.CleanTitle("Manager d’unité marchande H_F en Apprentissage")

	assert.Equal(t, "manager d’unité marchande", out)

	out = util.CleanTitle("Manager d’unité marchande (H-F)")

	assert.Equal(t, "manager d’unité marchande", out)

	out = util.CleanTitle("Manager d’unité marchande ( H / F ) en Apprentissage")

	assert.Equal(t, "manager d’unité marchande", out)

	out = util.CleanTitle("Offre poste secrétaire dentaire en contrat d'alternance")

	assert.Equal(t, "offre poste secrétaire dentaire", out)

	out = util.CleanTitle("MANAGER D’UNITE MARCHANDE")

	assert.Equal(t, "manager d’unite marchande", out)
}
