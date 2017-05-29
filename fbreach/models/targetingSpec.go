package models

type TargetingSpec struct {
	Genders []int `json:"genders"`
	AgeMin int `json:"age_min"`
	AgeMax int `json:"age_max"`
	Flexible_spec  []FlexibleSpec `json:"flexible_spec"`
	Geo_locations  GeoLocation `json:"geo_locations"`
}



