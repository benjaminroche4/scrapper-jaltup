package labonnealternance

import (
	"regexp"
	"scrapperjaltup/model"
	"scrapperjaltup/util"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/gosimple/slug"
)

func cleanCity(in string) string {
	parts := strings.Split(in, "-")
	c := cases.Title(language.French)
	r := regexp.MustCompile("^[^0-9]+$")

	for _, part := range parts {
		field := strings.TrimSpace(part)

		if r.MatchString(field) {
			return c.String(field)
		}
	}

	return c.String(in)
}

func cleanEmail(in string) string {
	r := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}`)

	return r.FindString(in)
}

func cleanURL(in string) string {
	r := regexp.MustCompile(
		`(http:\/\/www\.|https:\/\/www\.|http:\/\/|https:\/\/|\/|\/\/)?[A-z0-9_-]*?[:]?[A-z0-9_-]*?[@]?[A-z0-9]+([\-\.]{1}[a-z0-9]+)*\.[a-z]{2,5}(:[0-9]{1,5})?(\/.*)?`, // nolint: lll
	)

	return r.FindString(in)
}

func TranslatePlace(in *Place) *model.Place {
	address := in.Address
	if address == "" {
		address = in.FullAddress
	}

	return &model.Place{
		FullAddress: address,
		City:        cleanCity(in.City),
		ZipCode:     in.ZipCode,
		Latitude:    in.Latitude,
		Longitude:   in.Longitude,
	}
}

func TranslateJob(in *PeJob) *model.Job {
	duration, err := strconv.ParseInt(in.Job.Duration, 10, 16)
	if err != nil {
		duration = 0
	}
	startDate, err := time.Parse(time.RFC3339, in.Job.JobStartDate)
	if err != nil {
		startDate = time.Now().Truncate(24 * time.Hour)
	}

	return &model.Job{
		Description:  in.Job.Description,
		ContractType: in.Job.ContractType,
		Duration:     int16(duration),
		Remote:       in.Place.RemoteOnly,
		StudyLevel:   "",
		StartDate:    startDate,
	}
}

func TranslateCompany(in *PeJob) *model.Company {
	c := cases.Title(language.French)
	company := &model.Company{}

	createdAt, err := time.Parse(time.RFC3339, in.Company.CreationDate)
	if err != nil {
		createdAt = time.Unix(0, 0).UTC()
	}
	name := strings.TrimSpace(in.Company.Name)
	if name == "" {
		name = "<Vide>"
	}
	siret := strings.TrimSpace(in.Company.Siret)

	company.PublicID = util.GenerateUniqueID(10)
	company.Name = c.String(name)
	company.Siret = siret
	company.ContactEmail = cleanEmail(strings.TrimSpace(in.Contact.Email))
	company.PhoneNumber = strings.TrimSpace(in.Contact.Phone)
	company.WebSiteURL = cleanURL(in.Company.URL)
	company.Logo = in.Company.Logo
	company.CreatedAt = createdAt
	company.Slug = slug.Make(name)
	company.Verified = false

	return company
}

func TranslateOffer(in *PeJob) *model.Offer {
	var offer model.Offer

	createdAt, err := time.Parse(time.RFC3339, in.Job.CreationDate)
	if err != nil {
		createdAt = time.Now().Truncate(24 * time.Hour)
	}
	title := strings.TrimSpace(in.Title)
	url := strings.TrimSpace(in.URL)

	offer.ServiceName = "la-bonne-alternance"
	offer.ExternalID = in.ID
	offer.PublicID = util.GenerateUniqueID(10)
	offer.Title = title
	offer.Place = *TranslatePlace(&in.Place)
	offer.Job = *TranslateJob(in)
	offer.URL = cleanURL(url)
	offer.Status = "published"
	offer.CreatedAt = createdAt
	offer.EndAt = createdAt.AddDate(0, 3, 0)
	offer.Slug = slug.Make(title)
	offer.Premium = false
	offer.Company = *TranslateCompany(in)

	return &offer
}
