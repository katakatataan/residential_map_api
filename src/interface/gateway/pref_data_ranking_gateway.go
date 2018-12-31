package gateway

import (
	"residential_map_api/src/entity"
)

type PrefDataRankingGateway struct {
	SqlHandler
}

func (cdg *PrefDataRankingGateway) GetMonthlyPrefRankingOfBuildCount(begin string, end string) (entity.PrefDatasBuildCountRanking, error) {
	var prefDatas entity.PrefDatasBuildCountRanking
	// TODO: マスターコード変換
	err := cdg.Find(&prefDatas, "SELECT pref_name, SUM(built_count) as built_count, pref_id, to_char(build_date,'YYYY-MM') as build_date, rank() over( partition by date_trunc('month', build_date) order by SUM(built_count) desc) as monthly_rank FROM city_data WHERE build_date >= $1 AND build_date < $2 AND pref_name IS NOT NULL group by pref_id, pref_name, build_date ORDER BY date_trunc('month', build_date)", string([]rune(begin)[:10]), string([]rune(end)[:10]))
	if err != nil {
		return entity.PrefDatasBuildCountRanking{}, err
	}
	return prefDatas, nil
}
