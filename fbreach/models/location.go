package models

type Location struct {
	Distance_unit string `json:"distance_unit"`
	Key string `json:"key"`
	Name string `json:"name"`
	Region string `json:"region"`
	Region_id string `json:"region_id"`
	Country string `json:"country"`
	Radius int `json:"radius"`
}
