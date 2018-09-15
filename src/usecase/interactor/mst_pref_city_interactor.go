package interactor

import (
	"residential_map_api/src/entity"
	"residential_map_api/src/usecase/repository"
)

type mstprefcityinteractor struct {
	Repository repository.MstPrefCityRepository
}

type MstPrefCityInteractor interface {
	FetchAllPrefCities() (entity.PrefCities, error)
}

func NewMstPrefCityInteractor(repo repository.MstPrefCityRepository) MstPrefCityInteractor {
	return &mstprefcityinteractor{
		Repository: repo,
	}
}

func (mpci *mstprefcityinteractor) FetchAllPrefCities() (entity.PrefCities, error) {
	prefcities, err := mpci.Repository.FindAll()
	if err != nil {
		return entity.PrefCities{}, err
	}
	return prefcities, nil
}
