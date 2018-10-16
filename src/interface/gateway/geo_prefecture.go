package gateway

import (
	"residential_map_api/src/entity"
	"residential_map_api/src/entity/param"

	"github.com/k0kubun/pp"
)

type GeoPrefecture struct {
	SqlHandler
}

func (cdg *GeoPrefecture) FindByPrefId(mstGeoPrefParam *param.GeoPrefectureDto) ([]entity.MstPrefectureGeojson, error) {
	var mstGeojson []entity.MstPrefectureGeojson
	pp.Println(mstGeoPrefParam)
	// TODO ここを配列がきた時に対応してくれるように修正
	// 1. sqlxのin演算子を使う
	// 2. goでstringに展開する
	q := "SELECT * FROM mst_prefecture_geojson WHERE pref_id IN (" + mstGeoPrefParam.ToCsvPrefIds() + ") AND weight = $1"
	pp.Println(mstGeoPrefParam.ToCsvPrefIds())
	err := cdg.Find(&mstGeojson, q, mstGeoPrefParam.Weight)
	if err != nil {
		return []entity.MstPrefectureGeojson{}, err
	}
	return mstGeojson, nil
}
