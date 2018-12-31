package gateway

import (
	"residential_map_api/src/entity"
)

type CityDataRankingGateway struct {
	SqlHandler
}

func (cdg *CityDataRankingGateway) GetMonthlyCityRankingOfBuildCount(prefId int, begin string, end string) (entity.CityDatasBuildCountRanking, error) {
	var cityDatas entity.CityDatasBuildCountRanking
	// TODO: マスターコード変換
	err := cdg.Find(&cityDatas, "SELECT *,rank() over( partition by date_trunc('month',build_date) order by built_count desc) as monthly_rank FROM city_data WHERE pref_id = $1 AND  build_date >= $2 AND build_date < $3  ORDER BY date_trunc('month', build_date)", prefId, string([]rune(begin)[:10]), string([]rune(end)[:10]))
	if err != nil {
		return entity.CityDatasBuildCountRanking{}, err
	}
	return cityDatas, nil
}
