package models

type TemperatureInfo struct {
	City        string  `json:"city"`
	Unit        string  `json:"unit"`
	Temperature float32 `json:"temperature"`
}
