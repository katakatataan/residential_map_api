package entity

type FeatureCollection interface {
	AddFeature(feature *Feature) *FeatureCollection
}

type GeometryType string

type Feature struct {
	ID          interface{}            `json:"id,omitempty"`
	Type        string                 `json:"type"`
	BoundingBox []float64              `json:"bbox,omitempty"`
	Geometry    *Geometry              `json:"geometry"`
	Properties  map[string]interface{} `json:"properties"`
	CRS         map[string]interface{} `json:"crs,omitempty"` // Coordinate Reference System Objects are not currently supported
}

type Geometry struct {
	Type            GeometryType `json:"type"`
	BoundingBox     []float64    `json:"bbox,omitempty"`
	Point           []float64
	MultiPoint      [][]float64
	LineString      [][]float64
	MultiLineString [][][]float64
	Polygon         [][][]float64
	MultiPolygon    [][][][]float64
	Geometries      []*Geometry
	CRS             map[string]interface{} `json:"crs,omitempty"` // Coordinate Reference System Objects are not currently supported
}
