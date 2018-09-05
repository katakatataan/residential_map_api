package repository

import (
	"residential_map_api/src/entity"

	"github.com/jmoiron/sqlx"
)

type repo struct {
	conn *sqlx.DB
}

//ここで依存性の逆転のためにRepositoryのinterfaceを作成する
type MstPrefCityRepository interface {
	FindAll() (entity.PrefCities, error)
}

func NewMstPrefCityRepository(connection *sqlx.DB) MstPrefCityRepository {
	return &repo{
		conn: connection,
	}
}

func (r *repo) FindAll() (entity.PrefCities, error) {
	var prefcities entity.PrefCities
	err := r.conn.Select(&prefcities, "SELECT * FROM mst_pref_city ORDER BY id ASC")
	if err != nil {
		return entity.PrefCities{}, err
	}
	return prefcities, nil
}
