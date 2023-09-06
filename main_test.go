package gosolar

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Results struct {
	EOT                  float64 `json:"equation_of_time"`
	Declination          float64 `json:"declination"`
	SolarAngle           float64 `json:"solar_angle"`
	SolarZenith          float64 `json:"solar_zenith"`
	DayLength            float64 `json:"day_length"`
	Sunrise              float64 `json:"sunrise"`
	Sunset               float64 `json:"sunset"`
	JulianDay            float64 `json:"julian_day"`
	JulianCentury        float64 `json:"julian_century"`
	GeomMeanLongSun      float64 `json:"geom_mean_long_sun"`
	GeomMeanAnomSun      float64 `json:"geom_mean_anom_sun"`
	EccentricEarthOrbit  float64 `json:"eccentric_earth_orbit"`
	SolarNoon            float64 `json:"solar_noon"`
	SunEquationOfCenter  float64 `json:"sun_equation_of_center"`
	SunApparentLongitude float64 `json:"sun_apparent_longitude"`
	MeanObliqueEcliptic  float64 `json:"mean_oblique_ecliptic"`
	ObliqueCorrection    float64 `json:"oblique_correction"`
	SunHourAngle         float64 `json:"sun_hour_angle"`
	TimeZoneOffset       float64 `json:"time_zone_offset"`
}

func TestResults(t *testing.T) {

	latitude := 25.54821           // float Degrees
	longitude := -80.37486         // float Degrees
	date := "2023-09-05"           // string "YYYY-MM-DD"
	dayTime := 0.5                 // float time of the day/24
	timeZone := "America/New_York" // float UTC timezone in hours. GMT+12:30 = 12.5
	tz, _ := TimeZoneOffset("America/New_York")

	sun, err := Calculator(latitude, longitude, dayTime, timeZone, date, true)

	if err != nil {
		fmt.Println(err)
		return
	}

	sunrise, sunset := sun.SunriseAndSunset()

	results := &Results{
		EOT:                  sun.EquationOfTime(),
		Declination:          sun.SolarDeclination(),
		SolarAngle:           sun.SolarAltitudeAngle(),
		DayLength:            sun.DayLength(),
		Sunrise:              sunrise,
		Sunset:               sunset,
		SolarZenith:          sun.SolarZenithAngle(),
		JulianDay:            sun.JulianDay(),
		JulianCentury:        sun.JulianCentury(),
		GeomMeanLongSun:      sun.GeomMeanLongSun(),
		GeomMeanAnomSun:      sun.GeomMeanAnomSun(),
		EccentricEarthOrbit:  sun.EccentEarthOrbit(),
		SolarNoon:            sun.SolarNoon(),
		SunEquationOfCenter:  sun.SunEquationOfCenter(),
		SunApparentLongitude: sun.SunApparentLongitude(),
		MeanObliqueEcliptic:  sun.MeanObliqEcliptic(),
		ObliqueCorrection:    sun.ObliqueCorrection(),
		SunHourAngle:         sun.SunHourAngle(),
		TimeZoneOffset:       float64(tz) / 3600,
	}

	s, _ := json.MarshalIndent(results, "", "\t")
	fmt.Print(string(s))
}
