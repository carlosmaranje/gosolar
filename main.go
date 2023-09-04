package main

import (
	"encoding/json"
	"fmt"
	"gosolar/pckg"
)

type Results struct {
	EOT         float64 `json:"equation_of_time"`
	Declination float64 `json:"declination"`
	SolarAngle  float64 `json:"solar_angle"`
	SolarZenith float64 `json:"solar_zenith"`
	DayLength   float64 `json:"day_length"`
	Sunrise     float64 `json:"sunrise"`
	Sunset      float64 `json:"sunset"`
}

func main() {
	day := 247
	latitude := 45.0
	longitude := 15.0
	eot := gosolar.EquationOfTime(day)
	declination := gosolar.SolarDeclination(day)
	solarAngle := gosolar.SolarAltitudeAngle(day, latitude)
	dayLength := gosolar.DayLength(day, latitude)
	sunrise, sunset := gosolar.SunriseAndSunset(day, latitude, longitude, true)
	solarZenith := gosolar.SolarZenithAngle(day, latitude)
	results := &Results{
		EOT:         eot,
		Declination: declination,
		SolarAngle:  solarAngle,
		DayLength:   dayLength,
		Sunrise:     sunrise,
		Sunset:      sunset,
		SolarZenith: solarZenith,
	}

	s, _ := json.MarshalIndent(results, "", "\t")
	fmt.Print(string(s))
}
