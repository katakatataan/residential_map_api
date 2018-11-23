package gateway

import (
	"residential_map_api/src/entity"
)

type MstPrefCityGateway struct {
	SqlHandler
}

func (pcg *MstPrefCityGateway) FindAllPrefCities() (entity.PrefCities, error) {
	var prefCities entity.PrefCities
	err := pcg.Find(&prefCities, "SELECT * FROM mst_pref_city ORDER BY id ASC")
	if err != nil {
		return entity.PrefCities{}, err
	}
	return prefCities, nil
}

func (pcg *MstPrefCityGateway) FindPref() (entity.Prefs, error) {
	var prefs entity.Prefs
	err := pcg.Find(&prefs, "SELECT * FROM mst_pref ORDER BY id ASC")
	if err != nil {
		return entity.Prefs{}, err
	}
	return prefs, nil
}

func (pcg *MstPrefCityGateway) FindCity() (entity.Cities, error) {
	var cities entity.Cities
	err := pcg.Find(&cities, "SELECT * FROM mst_city ORDER BY id ASC")
	if err != nil {
		return entity.Cities{}, err
	}
	return cities, nil
}
