package source

import (
	"scrapperjaltup/model"
	lba "scrapperjaltup/source/labonnealternance"
	"scrapperjaltup/util"
	"time"

	"github.com/gosimple/slug"
)

type LBA struct {
	api           *lba.LaBonneAlternance
	romes         []string
	sleepDuration time.Duration
}

func NewLBA() Source {
	return &LBA{
		api:           lba.New(),
		romes:         lba.GetRomeCodes(),
		sleepDuration: 250 * time.Millisecond,
	}
}

func (thiz *LBA) RetrieveOffers(setProgression func(int)) ([]*model.Offer, error) {
	offers := []*model.Offer{}

	for i, rome := range thiz.romes {
		if setProgression != nil {
			setProgression(i * 100 / len(thiz.romes))
		}

		tags := lba.GetRomeTags(rome)
		categories := transformTagsInCategories(tags, nil)

		resp, err := thiz.api.JobFormations([]string{rome})
		if err != nil {
			return offers, err
		}
		time.Sleep(thiz.sleepDuration)

		for i := range resp.Jobs.PeJobs.Results {
			peJob := resp.Jobs.PeJobs.Results[i]
			offer := lba.TranslateOffer(&peJob)
			offer.Categories = categories
			offers = append(offers, offer)
		}
	}

	return offers, nil
}

func (thiz *LBA) RetrieveCategories(setProgression func(int)) ([]model.Category, error) {
	tags := lba.GetAllRomeTags()

	return transformTagsInCategories(tags, setProgression), nil
}

// Unexported function

func transformTagsInCategories(tags []string, setProgression func(int)) []model.Category {
	categories := []model.Category{}

	for i, tag := range tags {
		if setProgression != nil {
			setProgression(i * 100 / len(tags))
		}
		categories = append(categories, model.Category{
			ID:       0,
			PublicID: util.GenerateUniqueID(10),
			Name:     tag,
			Slug:     slug.Make(tag),
		})
	}

	return categories
}
