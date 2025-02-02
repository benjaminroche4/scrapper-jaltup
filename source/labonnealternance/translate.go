package labonnealternance

import (
	"scrapperjaltup/model"
	"scrapperjaltup/util"
	"scrapperjaltup/util/cities"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	_slug "github.com/gosimple/slug"
)

const (
	ValidDuration   = 30 * 24 * time.Hour
	PremiumDuration = 7 * 24 * time.Hour
)

func TranslatePlace(in *Place) *model.Place {
	address := in.Address
	if address == "" {
		address = in.FullAddress
	}
	name := util.CleanCityName(in.City)
	city := cities.FindCity(name)
	zipCode := in.ZipCode
	latitude := in.Latitude
	longitude := in.Longitude
	if city != nil {
		if zipCode == "" {
			zipCode = city.ZipCode
		}
		if latitude == 0.0 {
			latitude = city.Latitude
		}
		if longitude == 0.0 {
			longitude = city.Longitude
		}
	}

	return &model.Place{
		FullAddress: address,
		City:        name,
		ZipCode:     zipCode,
		Latitude:    latitude,
		Longitude:   longitude,
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
		ContractType: strings.ToUpper(in.Job.ContractType),
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
		name = "Inconnu"
	}
	email := util.Truncate(util.CleanEmail(strings.TrimSpace(in.Contact.Email)), 120)
	phone := util.Truncate(strings.TrimSpace(in.Contact.Phone), 20)
	siret := strings.TrimSpace(in.Company.Siret)

	company.PublicID = util.GenerateUniqueID(10)
	company.Name = util.Truncate(c.String(name), 120)
	company.Siret = siret
	company.ContactEmail = email
	company.PhoneNumber = phone
	company.WebSiteURL = util.Truncate(util.CleanURL(in.Company.URL), 255)
	company.Logo = in.Company.Logo
	company.CreatedAt = createdAt
	company.Slug = _slug.Make(name)
	company.Verified = false

	return company
}

func TranslateOffer(in *PeJob) *model.Offer {
	var offer model.Offer

	createdAt, err := time.Parse(time.RFC3339, in.Job.CreationDate)
	if err != nil {
		createdAt = time.Now().Truncate(24 * time.Hour)
	}
	endPremiumAt := createdAt.Add(PremiumDuration)
	title := util.Truncate(strings.TrimSpace(in.Title), 120)
	url := util.Truncate(strings.TrimSpace(in.URL), 255)
	slug := _slug.Make(title)
	status := "published"
	if len(title) < 3 {
		status = "archived"
	}

	offer.ServiceName = "la-bonne-alternance"
	offer.ExternalID = in.ID
	offer.PublicID = util.GenerateUniqueID(10)
	offer.Title = util.CleanTitle(title)
	offer.Place = *TranslatePlace(&in.Place)
	offer.Job = *TranslateJob(in)
	offer.URL = util.CleanURL(url)
	offer.Status = status
	offer.CreatedAt = createdAt
	offer.EndAt = createdAt.Add(ValidDuration)
	offer.EndPremiumAt = endPremiumAt
	offer.Slug = slug
	offer.Premium = false
	if time.Now().Unix() < endPremiumAt.Unix() {
		offer.Premium = true
	}
	offer.Company = *TranslateCompany(in)

	return &offer
}
