package interactor

import (
	"residential_map_api/src/entity/param"
	"residential_map_api/src/entity/response"
	"residential_map_api/src/usecase/repository"
)

type MstPrefCityInteractor struct {
	Repository repository.MstPrefCityRepository
}

func (mpci *MstPrefCityInteractor) FetchAllPrefCities(param *param.GetMasterPrefCitiesParam) (response.GetMasterPrefCitiesResponse, error) {
	prefcities, err := mpci.Repository.FindAllPrefCities(param.PrefId)
	res := response.GetMasterPrefCitiesResponse{
		prefcities,
	}
	if err != nil {
		return response.GetMasterPrefCitiesResponse{}, err
	}
	return res, nil
}

func (mpci *MstPrefCityInteractor) FetchPref() (response.GetMasterPrefResponse, error) {
	prefs, err := mpci.Repository.FindPref()
	res := response.GetMasterPrefResponse{
		prefs,
	}
	if err != nil {
		return response.GetMasterPrefResponse{}, err
	}
	return res, nil
}

func (mpci *MstPrefCityInteractor) FetchCity() (response.GetMasterCityResponse, error) {
	prefs, err := mpci.Repository.FindCity()
	res := response.GetMasterCityResponse{
		prefs,
	}
	if err != nil {
		return response.GetMasterCityResponse{}, err
	}
	return res, nil
}
