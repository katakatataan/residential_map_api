package interactor

import (
	"residential_map_api/src/usecase/dto"
	"residential_map_api/src/usecase/repository"
)

type MstPrefCityInteractor struct {
	Repository repository.MstPrefCityRepository
}

func (mpci *MstPrefCityInteractor) FetchAllPrefCities() (dto.PrefCities, error) {
	var prefcities dto.PrefCities
	prefcities, err := mpci.Repository.FindAll()
	if err != nil {
		return dto.PrefCities{}, err
	}
	return prefcities, nil
}
