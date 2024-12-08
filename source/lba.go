package source

import (
	"scrapperjaltup/model"
	lba "scrapperjaltup/source/labonnealternance"
	"scrapperjaltup/util"
	"time"

	"github.com/gosimple/slug"
	"github.com/liamg/loading/pkg/bar"
)

type LBA struct {
	api   *lba.LaBonneAlternance
	romes []string
}

func NewLBA() Source {
	return &LBA{
		api:   lba.New(),
		romes: lba.GetRomeCodes(),
	}
}

func (thiz *LBA) RetrieveOffers() ([]*model.Offer, error) {
	offers := []*model.Offer{}

	loadingBar := bar.New(
		bar.OptionHideOnFinish(true),
	)
	loadingBar.SetTotal(len(thiz.romes))
	loadingBar.SetLabel("progress")

	for i, rome := range thiz.romes {
		loadingBar.SetCurrent(i)

		tags := lba.GetRomeTags(rome)
		categories := transformTagsInCategories(tags)

		resp, err := thiz.api.JobFormations([]string{rome})
		if err != nil {
			return offers, err
		}
		time.Sleep(250 * time.Millisecond)

		for i := range resp.Jobs.PeJobs.Results {
			peJob := resp.Jobs.PeJobs.Results[i]
			offer := lba.TranslateOffer(&peJob)
			offer.Categories = categories
			offers = append(offers, offer)
		}
	}

	return offers, nil
}

func (thiz *LBA) RetrieveCategories() ([]model.Category, error) {
	tags := lba.GetAllRomeTags()

	return transformTagsInCategories(tags), nil
}

// Unexported function

func transformTagsInCategories(tags []string) []model.Category {
	categories := []model.Category{}

	for _, tag := range tags {
		categories = append(categories, model.Category{
			ID:       0,
			PublicID: util.GenerateUniqueID(10),
			Name:     tag,
			Slug:     slug.Make(tag),
		})
	}

	return categories
}
