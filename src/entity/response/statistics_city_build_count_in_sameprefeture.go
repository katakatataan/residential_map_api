package response

import null "gopkg.in/guregu/null.v3"

type ResStatisticsCityDatasBuildCountInSamePrefecture struct {
	// TODO ここを修正
	Data interface{} `json:"data"`
}
type ResStatisticsCityDatasBuildCountInSamePrefectureData struct {
	ResStatisticsCityDatasBuildCountInSamePrefectureCommon
	Cities ResStatisticsCityDatasBuildCountInSamePrefectureCity `json:"cities"`
}
type ResStatisticsCityDatasBuildCountInSamePrefectureCity struct {
	CityId           int         `json:"city_id"`
	CityName         null.String `json:"city_name"`
	BuiltCount       int         `json:"built_count"`
	TotalSquareMeter int         `json:"total_square_meter"`
	MonthlyRank      int         `json:"monthly_rank"`
}

type ResStatisticsCityDatasBuildCountInSamePrefectureCommon struct {
	Id                   int         `json:"id"`
	Year                 int         `json:"year"`
	Month                int         `json:"month"`
	ResidentialUseTypeId int         `json:"residential_use_type_id"`
	ConstructionTypeId   int         `json:"construction_type_id"`
	BuildTypeId          int         `json:"build_type_id"`
	ResidentialTypeId    int         `json:"residential_type_id"`
	StructureType        int         `json:"structure_type"`
	PrefId               int         `json:"pref_id"`
	PrefName             null.String `json:"pref_name"`
	BuildDate            string      `json:"build_date"`
}
