package repository

import "residential_map_api/src/entity"

type MstPrefCityRepository interface {
	FindPref() (entity.Prefs, error)
	FindCity() (entity.Cities, error)
	FindAllPrefCities(prefId int) (entity.PrefCities, error)
}
