package gateway

import (
	"residential_map_api/src/entity"
	"residential_map_api/src/entity/param"
)

type GeoPrefecture struct {
	SqlHandler
}

func (cdg *GeoPrefecture) FindByPrefId(mstGeoPrefParam *param.GeoPrefectureDto) ([]entity.MstPrefectureGeojson, error) {
	var mstGeojson []entity.MstPrefectureGeojson
	q := "SELECT * FROM mst_prefecture_geojson WHERE pref_id IN (" + mstGeoPrefParam.ToCsvPrefIds() + ") AND weight = $1"
	err := cdg.Find(&mstGeojson, q, mstGeoPrefParam.Weight)
	if err != nil {
		return []entity.MstPrefectureGeojson{}, err
	}
	return mstGeojson, nil
}

func (cdg *GeoPrefecture) FindBuildCountByPrefId(prefId int, weight float64, begin string, end string) ([]entity.MstPrefectureGeojson, error) {
	// TODO ここはprimitiveな値を受け取るように修正
	var mstGeojson []entity.MstPrefectureGeojson
	q := "SELECT * FROM mst_prefecture_geojson WHERE pref_id = $1 AND weight = $2 ORDER BY pref_id"
	err := cdg.Find(&mstGeojson, q, prefId, weight)
	if err != nil {
		return []entity.MstPrefectureGeojson{}, err
	}
	return mstGeojson, nil
}
