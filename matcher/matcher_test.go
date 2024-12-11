package matcher_test

import (
	"errors"
	_db "scrapperjaltup/db"
	"scrapperjaltup/model"
	"scrapperjaltup/source"

	"scrapperjaltup/matcher"
	"testing"

	"github.com/stretchr/testify/assert"
)

func prepare() (*_db.Mock, *source.Mock) {
	db := _db.NewMock().(*_db.Mock)
	source := source.NewMock().(*source.Mock)

	return db, source
}

func TestMatcher(t *testing.T) {
	t.Parallel()

	db, source := prepare()

	m := matcher.New(db, source)
	err := m.Execute()
	assert.NoError(t, err)

	assert.Len(t, db.Companies, 3)
	assert.Len(t, db.Offers, 3)

	db.AssertNumberOfCalls(t, "SelectCategories", 2)
	db.AssertNumberOfCalls(t, "InsertCategories", 1)

	db.AssertNumberOfCalls(t, "SelectCompanies", 1)
	db.AssertNumberOfCalls(t, "InsertCompanies", 1)

	db.AssertNumberOfCalls(t, "SelectOffers", 1)
	db.AssertNumberOfCalls(t, "InsertOffers", 1)

	source.AssertNumberOfCalls(t, "RetrieveOffers", 1)
	source.AssertNumberOfCalls(t, "RetrieveCategories", 1)
}

// Test Error on RetrieveOffers.
func TestErrorRetrieveOffers(t *testing.T) {
	t.Parallel()

	db, source := prepare()

	source.On("RetrieveOffers").Unset()
	source.On("RetrieveOffers").Return([]*model.Offer{}, errors.New("failed to fetch offers"))

	m := matcher.New(db, source)
	err := m.Execute()

	assert.Error(t, err)

	assert.Len(t, db.Companies, 2)
	assert.Len(t, db.Offers, 2)

	db.AssertNumberOfCalls(t, "SelectCategories", 2)
	db.AssertNumberOfCalls(t, "InsertCategories", 1)

	db.AssertNumberOfCalls(t, "SelectCompanies", 1)
	db.AssertNumberOfCalls(t, "InsertCompanies", 0)

	db.AssertNumberOfCalls(t, "SelectOffers", 1)
	db.AssertNumberOfCalls(t, "InsertOffers", 0)

	source.AssertNumberOfCalls(t, "RetrieveOffers", 1)
	source.AssertNumberOfCalls(t, "RetrieveCategories", 1)
}

func TestErrorRetrieveCategories(t *testing.T) {
	t.Parallel()

	db, source := prepare()

	source.On("RetrieveCategories").Unset()
	source.On("RetrieveCategories").Return([]model.Category{}, errors.New("failed to fetch offers"))

	m := matcher.New(db, source)
	err := m.Execute()

	assert.Error(t, err)

	assert.Len(t, db.Companies, 2)
	assert.Len(t, db.Offers, 2)

	db.AssertNumberOfCalls(t, "SelectCategories", 1)
	db.AssertNumberOfCalls(t, "InsertCategories", 0)

	db.AssertNumberOfCalls(t, "SelectCompanies", 0)
	db.AssertNumberOfCalls(t, "InsertCompanies", 0)

	db.AssertNumberOfCalls(t, "SelectOffers", 0)
	db.AssertNumberOfCalls(t, "InsertOffers", 0)

	source.AssertNumberOfCalls(t, "RetrieveOffers", 0)
	source.AssertNumberOfCalls(t, "RetrieveCategories", 1)
}

// Test Error on SelectCategories.
func TestErrorSelectCategories(t *testing.T) {
	t.Parallel()

	db, source := prepare()

	db.On("SelectCategories").Unset()
	db.On("SelectCategories").Return([]model.Category{}, errors.New("failed to retrieve categories"))

	m := matcher.New(db, source)
	err := m.Execute()

	assert.Error(t, err)

	assert.Len(t, db.Companies, 2)
	assert.Len(t, db.Offers, 2)

	db.AssertNumberOfCalls(t, "SelectCategories", 1)
	db.AssertNumberOfCalls(t, "InsertCategories", 0)

	db.AssertNumberOfCalls(t, "SelectCompanies", 0)
	db.AssertNumberOfCalls(t, "InsertCompanies", 0)

	db.AssertNumberOfCalls(t, "SelectOffers", 0)
	db.AssertNumberOfCalls(t, "InsertOffers", 0)

	source.AssertNumberOfCalls(t, "RetrieveOffers", 0)
	source.AssertNumberOfCalls(t, "RetrieveCategories", 0)
}

// Test Error on SelectCompanies.
func TestErrorSelectCompanies(t *testing.T) {
	t.Parallel()

	db, source := prepare()

	db.On("SelectCompanies").Unset()
	db.On("SelectCompanies").Return([]model.Company{}, errors.New("failed to retrieve companies"))

	m := matcher.New(db, source)
	err := m.Execute()

	assert.Error(t, err)

	assert.Len(t, db.Companies, 2)
	assert.Len(t, db.Offers, 2)

	db.AssertNumberOfCalls(t, "SelectCategories", 2)
	db.AssertNumberOfCalls(t, "InsertCategories", 1)

	db.AssertNumberOfCalls(t, "SelectCompanies", 1)
	db.AssertNumberOfCalls(t, "InsertCompanies", 0)

	db.AssertNumberOfCalls(t, "SelectOffers", 0)
	db.AssertNumberOfCalls(t, "InsertOffers", 0)

	source.AssertNumberOfCalls(t, "RetrieveOffers", 0)
	source.AssertNumberOfCalls(t, "RetrieveCategories", 1)
}

// Test Error on SelectOffers.
func TestErrorSelectOffers(t *testing.T) {
	t.Parallel()

	db, source := prepare()

	db.On("SelectOffers").Unset()
	db.On("SelectOffers").Return([]model.Offer{}, errors.New("failed to retrieve offers"))

	m := matcher.New(db, source)
	err := m.Execute()

	assert.Error(t, err)

	assert.Len(t, db.Companies, 2)
	assert.Len(t, db.Offers, 2)

	db.AssertNumberOfCalls(t, "SelectCategories", 2)
	db.AssertNumberOfCalls(t, "InsertCategories", 1)

	db.AssertNumberOfCalls(t, "SelectCompanies", 1)
	db.AssertNumberOfCalls(t, "InsertCompanies", 0)

	db.AssertNumberOfCalls(t, "SelectOffers", 1)
	db.AssertNumberOfCalls(t, "InsertOffers", 0)

	source.AssertNumberOfCalls(t, "RetrieveOffers", 0)
	source.AssertNumberOfCalls(t, "RetrieveCategories", 1)
}
