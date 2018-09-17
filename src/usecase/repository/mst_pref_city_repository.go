package repository

import (
	"residential_map_api/src/entity"
)

type MstPrefCityRepository interface {
	FindAll() (entity.PrefCities, error)
}
