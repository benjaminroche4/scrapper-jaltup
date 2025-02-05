package db_test

import (
	"fmt"
	_db "scrapperjaltup/db"
	"scrapperjaltup/model"
	"scrapperjaltup/util"
	"testing"
	"time"

	lorem "github.com/derektata/lorem/ipsum"
	"github.com/google/uuid"
	"github.com/gosimple/slug"
	"github.com/stretchr/testify/assert"
)

const LOGOURL = "https://www.coca-cola.fr/content/dam/onexp/fr/fr/lead/le-logo-coca-cola-huit-lettres-un-trait-dunion.jpg"

func openDB(t *testing.T) _db.Database {
	t.Helper()

	db := _db.New("localhost", 3306, "root", "root", "jaltup")
	err := db.Open()
	assert.NoError(t, err)

	return db
}

func TestOpen(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	db := openDB(t)
	defer db.Close()
}

func TestCountCategories(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	db := openDB(t)
	defer db.Close()

	count, err := db.CountCategories()
	assert.NoError(t, err)

	fmt.Printf("count = %v\n", count)
}

func TestSelectCategories(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	db := openDB(t)
	defer db.Close()

	categories, err := db.SelectCategories()
	assert.NoError(t, err)

	fmt.Printf("categories = %v\n", categories)
}

func TestInsertCategories(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	db := openDB(t)
	defer db.Close()

	categories := []model.Category{}

	categories = append(categories, model.Category{
		PublicID: util.GenerateUniqueID(10),
		Name:     "Développement Web",
		Slug:     slug.Make("Développement Web"),
	}, model.Category{
		PublicID: util.GenerateUniqueID(10),
		Name:     "Gestion de Projet",
		Slug:     slug.Make("Gestion de Projet"),
	}, model.Category{
		PublicID: util.GenerateUniqueID(10),
		Name:     "Transformation Digitale",
		Slug:     slug.Make("Transformation Digitale"),
	})

	err := db.InsertCategories(categories)
	assert.NoError(t, err)

	fmt.Printf("categories = %v\n", categories)
}

func TestCountCompanies(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	db := openDB(t)
	defer db.Close()

	count, err := db.CountCompanies()
	assert.NoError(t, err)

	fmt.Printf("count = %v\n", count)
}

func TestSelectCompanies(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	db := openDB(t)
	defer db.Close()

	companies, err := db.SelectCompanies()
	assert.NoError(t, err)

	fmt.Printf("companies = %v\n", companies)
}

func TestInsertCompanies(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	db := openDB(t)
	defer db.Close()

	companies := []model.Company{
		{
			ID:           0,
			PublicID:     util.GenerateUniqueID(10),
			Name:         "Coca-Cola",
			Siret:        "123456789",
			ContactEmail: "jfoe@cocacola.fr",
			PhoneNumber:  "0102030405",
			WebSiteURL:   "https://www.coca-cola.fr",
			Logo:         LOGOURL,
			CreatedAt:    time.Now().UTC(),
			Slug:         slug.Make("Coca-Cola"),
			Verified:     true,
		},
	}

	err := db.InsertCompanies(companies)
	assert.NoError(t, err)

	fmt.Printf("companies = %v\n", companies)
}

func TestCountOffers(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	db := openDB(t)
	defer db.Close()

	count, err := db.CountOffers()
	assert.NoError(t, err)

	fmt.Printf("count = %v\n", count)
}

func TestSelectOffers(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	db := openDB(t)
	defer db.Close()

	offers, err := db.SelectOffers()
	assert.NoError(t, err)

	fmt.Printf("offers = %v\n", offers)
}

func TestInsertOffers(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	db := openDB(t)
	defer db.Close()

	g := lorem.NewGenerator()
	g.WordsPerSentence = 10
	g.SentencesPerParagraph = 5
	g.CommaAddChance = 3

	companies, err := db.SelectCompanies()
	assert.NoError(t, err)
	assert.NotEmpty(t, companies)

	categories, err := db.SelectCategories()
	assert.NoError(t, err)
	assert.NotEmpty(t, categories)

	id := uuid.New()
	offer := []model.Offer{
		{
			ID:       0,
			Company:  companies[0],
			PublicID: util.GenerateUniqueID(10),
			Title:    "Chef de Projet Informatique",
			Place: model.Place{
				FullAddress: "23 Av. des Champs-Élysées",
				City:        "Paris",
				ZipCode:     "75016",
				Latitude:    48.8692752,
				Longitude:   2.3053147,
			},
			Job: model.Job{
				Description:  g.Generate(100),
				ContractType: "CDD",
				Duration:     8,
				Remote:       false,
				StudyLevel:   "Bac",
				StartDate:    time.Now(),
			},
			URL:          "https://www.coca-cola.fr",
			Tag:          []string{"Gestion de Projet", "Transformation Digitale", "Cloud Computing"},
			Status:       "published",
			CreatedAt:    time.Now(),
			EndAt:        time.Now().AddDate(0, 0, 30),
			EndPremiumAt: time.Now().AddDate(0, 0, 7),
			Slug:         slug.Make("Chef de Projet Informatique"),
			Premium:      false,
			ExternalID:   id.String(),
			ServiceName:  "labonnealternance",
			Categories:   []model.Category{categories[0], categories[2]},
		},
	}

	err = db.InsertOffers(offer)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, offer[0].ID)

	fmt.Printf("offer = %v\n", offer)
}
