package interactor

import (
	"residential_map_api/src/entity"
	"residential_map_api/src/usecase/interactor/param"
	"residential_map_api/src/usecase/interactor/response"
	"residential_map_api/src/usecase/repository"
	"strconv"

	"github.com/k0kubun/pp"
	geojson "github.com/paulmach/go.geojson"
)

const CityIdKey = "N03_007"

type CityDataInteractor struct {
	CityDataRepository        repository.CityDataRepository
	CityDataRankingRepository repository.CityDataRankingRepository
	//起動確認したら修正 GeojsonRepository
	GeojsonRepository repository.GeojsonRepository
}

func (cdi *CityDataInteractor) FindByCityId(param *param.GetCitiesCityIdParam) (response.GetCitiesCityIdResponse, error) {
	var citydata entity.CityDatas
	citydata, err := cdi.CityDataRepository.FindByCityId(param.CityId, param.Begin, param.End)
	res := response.GetCitiesCityIdResponse{
		Data: citydata,
	}
	if err != nil {
		return response.GetCitiesCityIdResponse{}, err
	}
	return res, nil
}

func (cdi *CityDataInteractor) FindCityRankingBuildCount(cityDataParam *param.GetCitiesRankingBuildCountParam) (response.GetCitiesRankingBuildCountResponse, error) {
	citydata, err := cdi.CityDataRepository.FindCityRankingBuildCount(cityDataParam.PrefId, cityDataParam.Begin, cityDataParam.End)
	res := response.GetCitiesRankingBuildCountResponse{
		Data: citydata,
	}
	// pp.Println(res)
	if err != nil {
		return response.GetCitiesRankingBuildCountResponse{}, err
	}
	return res, nil
}

func (cdi *CityDataInteractor) GetCityDataByTargetPeriod(param *param.GetCitiesCityIdMonthlyParam) (response.GetCitiesCityIdMonthlyResponse, error) {
	citydata, err := cdi.CityDataRepository.FindByCityIdByTargetPeriod(param.CityId, param.Begin, param.End)
	res := response.GetCitiesCityIdMonthlyResponse{
		Data: citydata,
	}
	if err != nil {
		return response.GetCitiesCityIdMonthlyResponse{}, err
	}
	return res, nil
}

func (gpi *CityDataInteractor) FindCitiesGeojsonByPrefId(param *param.GetCitiesGeojsonParam) (response.GetCitiesGeojsonResponse, error) {
	result, err := gpi.GeojsonRepository.FindByPrefId(param)
	var res response.GetCitiesGeojsonResponse
	for _, c := range result {
		fc, err := geojson.UnmarshalFeatureCollection(c.Json)
		if err != nil {
			return response.GetCitiesGeojsonResponse{}, err
		}
		for _, v := range fc.Features {
			res.AddFeature(v)
		}
	}

	if err != nil {
		return response.GetCitiesGeojsonResponse{}, err
	}
	return res, nil
}

func (gpi *CityDataInteractor) FindCitiesGeojsonWithBuildCountByPrefId(param *param.GetCitiesGeojsonBuildCountParam) (response.GetCitiesGeojsonBuildCountResponse, error) {
	// TODO: 後々に複数の都道府県に対応する
	pp.Println("hello")
	result, err := gpi.GeojsonRepository.FindBuildCountByPrefId(param.PrefIds[0], param.Weight, param.Begin, param.End)
	pp.Println(err)
	var res response.GetCitiesGeojsonBuildCountResponse
	citydata, _ := gpi.CityDataRepository.FindByPrefId(param.PrefIds[0], param.Begin, param.End)
	if err != nil {
		pp.Println(err)
		return response.GetCitiesGeojsonBuildCountResponse{}, err
	}
	// FIXME : 処理が無駄に回ってしまうので最適化
	for _, c := range result {
		fc, err := geojson.UnmarshalFeatureCollection(c.Json)
		if err != nil {
			pp.Println(err)
			return response.GetCitiesGeojsonBuildCountResponse{}, err
		}
		for _, v := range fc.Features {
			// FIXME : 処理が無駄に回ってしまうので最適化
			// N0#007にnilのvalueが存在する
			if v.Properties[CityIdKey] == nil {
				continue
			}
			cityId := v.Properties[CityIdKey].(string)
			cityIdInt, _ := strconv.Atoi(cityId)
			for _, d := range citydata {
				if d.CityId == cityIdInt {
					v.Properties["build_count"] = d.BuiltCount
					v.Properties["total_square_meter"] = d.TotalSquareMeter
					v.Properties["residential_use_type"] = d.ResidentialUseType
					v.Properties["construction_type"] = d.ConstructionType
					break
				}
			}
			res.AddFeature(v)
		}
	}

	if err != nil {
		return response.GetCitiesGeojsonBuildCountResponse{}, err
	}
	return res, nil
}
