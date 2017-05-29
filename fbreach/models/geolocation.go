package models

type GeoLocation struct {
	Location_types []string  `json:"location_types"`
	Cities         []Location `json:"cities"`
}

func ( geoLocation *GeoLocation) SetLocation_Types(locationType []string) {
	geoLocation.Location_types =  locationType
}

func ( geoLocation *GeoLocation) SetCities(cities []Location) {
	
	geoLocation.Cities =cities
}