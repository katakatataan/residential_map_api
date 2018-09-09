package interactor

import (
	"residential_map_api/src/entity"
	"residential_map_api/src/usecase/repository"
)

type interactor struct {
	Repository repository.MstPrefCityRepository
}

type MstPrefCityInteractor interface {
	FetchAllPrefCities() (entity.PrefCities, error)
}

func NewMstPrefCityInteractor(repo repository.MstPrefCityRepository) MstPrefCityInteractor {
	return &interactor{
		Repository: repo,
	}
}

func (mpci *interactor) FetchAllPrefCities() (entity.PrefCities, error) {
	// 理屈としてはこのusecaseレベルでconnecionを使いまわしたい
	prefcities, err := mpci.Repository.FindAll()
	if err != nil {
		return entity.PrefCities{}, err
	}
	return prefcities, nil
}
