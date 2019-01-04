package gateway

import (
	"residential_map_api/src/entity"
	"residential_map_api/src/usecase/interactor/param"
)

type GeojsonGateway struct {
	SqlHandler
}

func (cdg *GeojsonGateway) FindByPrefId(mstGeoPrefParam *param.GetCitiesGeojsonParam) ([]entity.MstPrefectureGeojson, error) {
	var mstGeojson []entity.MstPrefectureGeojson
	// csvはsharedに入れる
	q := `SELECT
			*
		FROM
			mst_prefecture_geojson
		WHERE
			pref_id IN (` + mstGeoPrefParam.ToCsvPrefIds() + `)
		AND
			weight = $1`
	err := cdg.Find(&mstGeojson, q, mstGeoPrefParam.Weight)
	if err != nil {
		return []entity.MstPrefectureGeojson{}, err
	}
	return mstGeojson, nil
}

func (cdg *GeojsonGateway) FindBuildCountByPrefId(prefId int, weight float64, begin string, end string) ([]entity.MstPrefectureGeojson, error) {
	// TODO ここはprimitiveな値を受け取るように修正
	// TODO: マスターコード変換
	var mstGeojson []entity.MstPrefectureGeojson
	q := `SELECT
			*
		FROM
			mst_prefecture_geojson
		WHERE
			pref_id = $1
		AND
			weight = $2
		ORDER BY pref_id`
	err := cdg.Find(&mstGeojson, q, prefId, weight)
	if err != nil {
		return []entity.MstPrefectureGeojson{}, err
	}
	return mstGeojson, nil
}
