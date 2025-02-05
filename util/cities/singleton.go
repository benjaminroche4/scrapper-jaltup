package cities

import "sync"

var (
	lock     = &sync.Mutex{}
	instance *CityMapper
)

func GetCityMapper() *CityMapper {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = NewCityMapper()
		}
	}

	return instance
}
