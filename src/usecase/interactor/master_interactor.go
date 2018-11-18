package interactor

import (
	"residential_map_api/src/entity/response"
	"residential_map_api/src/usecase/repository"
)

type MstPrefCityInteractor struct {
	Repository repository.MstPrefCityRepository
}

func (mpci *MstPrefCityInteractor) FetchAllPrefCities() (response.ResMasterPrefCities, error) {
	prefcities, err := mpci.Repository.FindAllPrefCities()
	res := response.ResMasterPrefCities{
		prefcities,
	}
	if err != nil {
		return response.ResMasterPrefCities{}, err
	}
	return res, nil
}
