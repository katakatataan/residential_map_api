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

func (mpci *MstPrefCityInteractor) FetchPref() (response.ResMasterPref, error) {
	prefs, err := mpci.Repository.FindPref()
	res := response.ResMasterPref{
		prefs,
	}
	if err != nil {
		return response.ResMasterPref{}, err
	}
	return res, nil
}

func (mpci *MstPrefCityInteractor) FetchCity() (response.ResMasterCity, error) {
	prefs, err := mpci.Repository.FindCity()
	res := response.ResMasterCity{
		prefs,
	}
	if err != nil {
		return response.ResMasterCity{}, err
	}
	return res, nil
}
