package repository

import "residential_map_api/src/entity"

type MstPrefCityRepository interface {
	FindAllPrefCities() (entity.PrefCities, error)
}
