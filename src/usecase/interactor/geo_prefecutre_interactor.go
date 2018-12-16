package interactor

import (
	"residential_map_api/src/entity/param"
	"residential_map_api/src/entity/response"
	"residential_map_api/src/usecase/repository"
	"strconv"

	"github.com/k0kubun/pp"
	geojson "github.com/paulmach/go.geojson"
)

type GeoPrefectureInteractor struct {
	Repository         repository.GeoPrefectureRepository
	CityDataRepository repository.CityDataRepository
}

func (gpi *GeoPrefectureInteractor) FindByPrefId(geoPrefParam *param.GeoPrefectureDto) (response.ResGeojsonFeatureCollection, error) {
	result, err := gpi.Repository.FindByPrefId(geoPrefParam)
	var res response.ResGeojsonFeatureCollection
	for _, c := range result {
		fc, err := geojson.UnmarshalFeatureCollection(c.Json)
		if err != nil {
			return response.ResGeojsonFeatureCollection{}, err
		}
		for _, v := range fc.Features {
			// TODO: ここのloopのレベルでデータを結合
			pp.Println(v.Properties["N03_007"])
			res.AddFeature(v)
		}
	}

	if err != nil {
		return response.ResGeojsonFeatureCollection{}, err
	}
	return res, nil
}

func (gpi *GeoPrefectureInteractor) FindBuildCountByPrefId(geoPrefParam *param.GeoPrefectureWithPeriodDto) (response.ResGeojsonFeatureCollection, error) {
	// TODO: 後々に複数の都道府県に対応する
	result, err := gpi.Repository.FindBuildCountByPrefId(geoPrefParam.PrefIds[0], geoPrefParam.Weight, geoPrefParam.Begin, geoPrefParam.End)
	var res response.ResGeojsonFeatureCollection
	citydata, _ := gpi.CityDataRepository.FindByPrefId(geoPrefParam.PrefIds[0], geoPrefParam.Begin, geoPrefParam.End)
	if err != nil {
		return response.ResGeojsonFeatureCollection{}, err
	}
	// TODO: ここは複数の都道府県を想定してしまっている
	// TODO: pythonで言う所のこの処理をやりたい。 data converterとしてpythonを用意する?
	// pgd_df_pref_geojson = gpd.read_file('http://54.249.162.108/geojson/pref?pref_id=13&weight=0.0001')
	// pgd_df_pref_geojson.rename(columns={
	// 													 'N03_007': 'city_id', 'N03_001': 'pref_name', 'N03_004': 'city_name'}, inplace=True)
	// df_citydata = read_frame(citydata)
	// print(df_citydata)
	// df_citydata['city_id'] = df_citydata['city_id'].astype(str).str.zfill(5)
	// merged = pgd_df_pref_geojson.merge(df_citydata, on='city_id')
	for _, c := range result {
		fc, err := geojson.UnmarshalFeatureCollection(c.Json)
		if err != nil {
			return response.ResGeojsonFeatureCollection{}, err
		}
		for _, v := range fc.Features {
			// FIXME :ここでこけるのでプログラムでいい感じに判定
			cityId := v.Properties["N03_007"].(string)
			cityIdInt, _ := strconv.Atoi(cityId)
			for _, d := range citydata {
				if d.CityId == cityIdInt {
					v.Properties["statistics"] = d
					break
				}
			}
			res.AddFeature(v)
		}
	}

	if err != nil {
		return response.ResGeojsonFeatureCollection{}, err
	}
	return res, nil
}
