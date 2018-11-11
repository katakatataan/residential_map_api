package gateway

import (
	"residential_map_api/src/entity"
)

type PrefDataRankingGateway struct {
	SqlHandler
}

func (cdg *PrefDataRankingGateway) GetMonthlyPrefRankingOfBuildCount(begin string, end string) (entity.PrefDatasBuildCountRanking, error) {
	var prefDatas entity.PrefDatasBuildCountRanking
	err := cdg.Find(&prefDatas, "SELECT pref_name, SUM(built_count) as built_count, pref_id, to_char(build_date,'YY') as build_year, to_char(build_date,'MM') as build_month, rank() over( partition by date_trunc('month', build_date) order by SUM(built_count) desc) as monthly_rank FROM city_data WHERE build_date >= $1 AND build_date < $2 AND pref_name IS NOT NULL group by pref_id, pref_name, build_date ORDER BY date_trunc('month', build_date)", begin, end)
	if err != nil {
		return entity.PrefDatasBuildCountRanking{}, err
	}
	return prefDatas, nil
}
