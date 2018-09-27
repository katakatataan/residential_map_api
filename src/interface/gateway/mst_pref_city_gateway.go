package gateway

import "residential_map_api/src/usecase/dto"

type MstPrefCityGateway struct {
	SqlHandler
}

func (pcg *MstPrefCityGateway) FindAll() (dto.PrefCities, error) {
	var prefCities dto.PrefCities
	err := pcg.Find(&prefCities, "SELECT * FROM mst_pref_city ORDER BY id ASC limit 10")
	if err != nil {
		return dto.PrefCities{}, err
	}
	return prefCities, nil
}
