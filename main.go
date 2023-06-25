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
	DayLength   float64 `json:"day_length"`
	Sunrise     float64 `json:"sunrise"`
	Sunset      float64 `json:"sunset"`
}

func main() {
	day := 175
	latitude := 25.4687224
	longitude := -80.37
	eot := gosolar.EquationOfTime(day)
	declination := gosolar.SolarDeclination(day)
	solarAngle := gosolar.SolarAltitudeAngle(day, latitude)
	dayLength := gosolar.DayLength(day, latitude)
	sunrise, sunset := gosolar.SunriseAndSunset(day, latitude, longitude)
	results := &Results{
		EOT:         eot,
		Declination: declination,
		SolarAngle:  solarAngle,
		DayLength:   dayLength,
		Sunrise:     sunrise,
		Sunset:      sunset,
	}

	s, _ := json.MarshalIndent(results, "", "\t")
	fmt.Print(string(s))
}
