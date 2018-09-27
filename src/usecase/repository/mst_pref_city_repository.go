package repository

import (
	"residential_map_api/src/usecase/dto"
)

type MstPrefCityRepository interface {
	FindAll() (dto.PrefCities, error)
}
