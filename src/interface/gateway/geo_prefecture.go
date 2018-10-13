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
	q := "SELECT * FROM mst_prefecture_geojson WHERE pref_id in ($1, $2) AND weight = $3"
	err := cdg.Find(&mstGeojson, q, mstGeoPrefParam.PrefIds[0], mstGeoPrefParam.PrefIds[1], mstGeoPrefParam.Weight)
	if err != nil {
		return []entity.MstPrefectureGeojson{}, err
	}
	return mstGeojson, nil
}
