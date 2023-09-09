package gosolar

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Results struct {
	EOT                      float64 `json:"equation_of_time"`
	Declination              float64 `json:"declination"`
	SolarAngle               float64 `json:"solar_angle"`
	SolarZenith              float64 `json:"solar_zenith"`
	DayLength                float64 `json:"day_length"`
	Sunrise                  float64 `json:"sunrise"`
	Sunset                   float64 `json:"sunset"`
	JulianDay                float64 `json:"julian_day"`
	JulianCentury            float64 `json:"julian_century"`
	GeomMeanLongSun          float64 `json:"geom_mean_long_sun"`
	GeomMeanAnomSun          float64 `json:"geom_mean_anom_sun"`
	EccentricEarthOrbit      float64 `json:"eccentric_earth_orbit"`
	SolarNoon                float64 `json:"solar_noon"`
	SunEquationOfCenter      float64 `json:"sun_equation_of_center"`
	SunApparentLongitude     float64 `json:"sun_apparent_longitude"`
	MeanObliqueEcliptic      float64 `json:"mean_oblique_ecliptic"`
	ObliqueCorrection        float64 `json:"oblique_correction"`
	SunHourAngle             float64 `json:"sun_hour_angle"`
	TimeZoneOffset           float64 `json:"time_zone_offset"`
	SolarZenithAngle         float64 `json:"solar_zenith_angle"`
	TrueSolarTime            float64 `json:"true_solar_time"`
	SolarIncidenceAngle      float64 `json:"solar_incidence_angle"`
	SolarAzimuthAngle        float64 `json:"solar_azimuth_angle"`
	IncidenceOnTiltedSurface float64 `json:"incidence_on_tilted_surface"`
}

func TestResults(t *testing.T) {

	latitude := 35.0               // float Degrees
	longitude := -80.37486         // float Degrees
	date := "2023-06-16"           // string "YYYY-MM-DD"
	dayTime := 0.64                // float time of the day/24
	timeZone := "America/New_York" // string Timezone ID
	tz, _ := TimeZoneOffset("America/New_York")

	sun, err := Calculator(latitude, longitude, dayTime, timeZone, date)

	if err != nil {
		fmt.Println(err)
		return
	}

	sunrise, sunset := sun.SunriseAndSunset()

	results := &Results{
		EOT:                      sun.EquationOfTime(),
		Declination:              sun.SolarDeclination(),
		DayLength:                sun.DayLength(),
		Sunrise:                  sunrise,
		Sunset:                   sunset,
		JulianDay:                sun.JulianDay(),
		JulianCentury:            sun.JulianCentury(),
		GeomMeanLongSun:          sun.GeomMeanLongSun(),
		GeomMeanAnomSun:          sun.GeomMeanAnomSun(),
		EccentricEarthOrbit:      sun.EccentEarthOrbit(),
		SolarNoon:                sun.SolarNoon(),
		SunEquationOfCenter:      sun.SunEquationOfCenter(),
		SunApparentLongitude:     sun.SunApparentLongitude(),
		MeanObliqueEcliptic:      sun.MeanObliqEcliptic(),
		ObliqueCorrection:        sun.ObliqueCorrection(),
		SunHourAngle:             sun.SunHourAngle(),
		TimeZoneOffset:           float64(tz) / 3600,
		SolarZenithAngle:         sun.SolarZenithAngle(),
		TrueSolarTime:            sun.TrueSolarTime(),
		SolarIncidenceAngle:      sun.SolarIncidenceAngle(),
		SolarAzimuthAngle:        sun.SolarAzimuthAngle(),
		IncidenceOnTiltedSurface: sun.IncidenceOnTiltedSurface(45, 10),
	}

	s, _ := json.MarshalIndent(results, "", "\t")
	fmt.Print(string(s))
}
