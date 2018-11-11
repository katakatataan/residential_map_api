package gateway

import (
	"residential_map_api/src/entity"
)

type MstPrefCityGateway struct {
	SqlHandler
}

func (pcg *MstPrefCityGateway) FindAll() (entity.PrefCities, error) {
	var prefCities entity.PrefCities
	err := pcg.Find(&prefCities, "SELECT * FROM mst_pref_city ORDER BY id ASC")
	if err != nil {
		return entity.PrefCities{}, err
	}
	return prefCities, nil
}
