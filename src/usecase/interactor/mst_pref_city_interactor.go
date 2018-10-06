package interactor

import (
	"residential_map_api/src/entity"
	"residential_map_api/src/usecase/repository"
)

type MstPrefCityInteractor struct {
	Repository repository.MstPrefCityRepository
}

func (mpci *MstPrefCityInteractor) FetchAllPrefCities() (entity.PrefCities, error) {
	var prefcities entity.PrefCities
	prefcities, err := mpci.Repository.FindAll()
	if err != nil {
		return entity.PrefCities{}, err
	}
	return prefcities, nil
}
