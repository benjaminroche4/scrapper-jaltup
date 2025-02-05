package cities

import (
	"sync"
	"time"

	_slug "github.com/gosimple/slug"

	"github.com/alex-cos/geoapi"
)

type CityMapper struct {
	api    *geoapi.GeoAPI
	cities map[string]*City
	mu     sync.Mutex
}

func NewCityMapper() *CityMapper {
	return &CityMapper{
		api:    geoapi.NewWithTimeout(5 * time.Second),
		cities: cities,
		mu:     sync.Mutex{},
	}
}

func (c *CityMapper) FindCity(name string) *City {
	slug := _slug.Make(name)
	c.mu.Lock()
	city, ok := c.cities[slug]
	c.mu.Unlock()
	if ok {
		return city
	}

	city = c.FetchCity(name)
	if city != nil {
		c.mu.Lock()
		c.cities[slug] = city
		c.mu.Unlock()

		return city
	}

	return nil
}

func (c *CityMapper) FetchCity(name string) *City {
	slug := _slug.Make(name)

	items, err := c.api.SearchMunicipality(name)
	if err != nil || items == nil {
		return nil
	}
	for _, item := range items[:min(10, len(items))] {
		if slug == _slug.Make(item.Name) {
			geocity, err := c.api.GetDetailedMunicipality(item.Code)
			if err == nil {
				zipcode := ""
				if len(geocity.PostalCodes) > 0 {
					zipcode = geocity.PostalCodes[0]
				}

				return &City{
					Name:       geocity.Name,
					ZipCode:    zipcode,
					Department: geocity.Department.Name,
					Region:     geocity.Region.Name,
					Latitude:   geocity.Center.Latitude(),
					Longitude:  geocity.Center.Longitude(),
				}
			}
		}
	}

	return nil
}
