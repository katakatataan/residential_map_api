package interactor

import (
	"residential_map_api/src/entity"
	"residential_map_api/src/usecase/repository"

	"github.com/jmoiron/sqlx"
)

type interactor struct {
	conn *sqlx.DB
}

type MstPrefCityInteractor interface {
	FetchAllPrefCities() (entity.PrefCities, error)
}

func NewMstPrefCityInteractor(conn *sqlx.DB) MstPrefCityInteractor {
	return &interactor{
		conn: conn,
	}
}

func (mpci *interactor) FetchAllPrefCities() (entity.PrefCities, error) {
	// 理屈としてはこのusecaseレベルでconnecionを使いまわしたい
	mpc := repository.NewMstPrefCityRepository(mpci.conn)
	prefcities, err := mpc.FindAll()
	if err != nil {
		return entity.PrefCities{}, err
	}
	return prefcities, nil
}
