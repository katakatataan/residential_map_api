package gateway

import (
	"residential_map_api/src/entity"
)

type PrefDataGateway struct {
	SqlHandler
}

func (cdg *PrefDataGateway) FindByPrefId(prefId int, begin string, end string) (entity.PrefDatas, error) {
	var prefDatas entity.PrefDatas
	// TODO: マスターコード変換
	err := cdg.Find(&prefDatas, "SELECT pref_id, pref_name, SUM(built_count) as built_count, SUM(total_square_meter) AS total_square_meter,to_char(build_date,'YYYY-MM') as build_date FROM city_data WHERE pref_id = $1 AND build_date >= $2 AND build_date < $3 GROUP BY pref_id, pref_name,  build_date ORDER BY  build_date ASC", prefId, string([]rune(begin)[:10]), string([]rune(end)[:10]))
	if err != nil {
		return entity.PrefDatas{}, err
	}
	return prefDatas, nil
}
