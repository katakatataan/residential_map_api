package repository

import (
	"residential_map_api/src/entity"
)

type repo struct {
	handler SqlHandler
}

//ここで依存性の逆転のためにRepositoryのinterfaceを作成する
type MstPrefCityRepository interface {
	FindAll() (entity.PrefCities, error)
}

func NewMstPrefCityRepository(handler SqlHandler) MstPrefCityRepository {
	return repo{
		handler: handler,
	}
}

func (r repo) FindAll() (entity.PrefCities, error) {
	var prefcities entity.PrefCities
	err := r.handler.Find(&prefcities, "SELECT * FROM mst_pref_city ORDER BY id ASC")
	if err != nil {
		return entity.PrefCities{}, err
	}
	return prefcities, nil
}
