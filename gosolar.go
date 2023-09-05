package gosolar

import (
	"math"
)

// EquationOfTime Calculates the value for the equation of time for any given day of the year
func EquationOfTime(day int) float64 {
	p := float64(day-81) * (360.0 / 364)

	// Convert p to radians
	p = toRadians(p)

	return 9.87*math.Sin(2*p) - 7.53*math.Cos(p) - 1.5*math.Sin(p)
}

func SolarNoon(day int, longitude float64) float64 {
	lsm := float64(standardMeridian(longitude))
	return (12 - EquationOfTime(day)) - 4*(lsm-longitude)
}

// SolarDeclination returns the declination in degrees for any given day of the year
func SolarDeclination(day int) float64 {
	p := (360 / 365.0) * float64(284+day)

	return 23.45 * math.Sin(toRadians(p))
}

func SolarAltitudeAngle(day int, latitude float64) float64 {
	declination := toRadians(SolarDeclination(day))
	hourAngle := toRadians(float64(30))
	latitude = toRadians(latitude)
	eq := math.Sin(latitude)*math.Sin(declination) + math.Cos(latitude)*math.Cos(declination)*math.Cos(hourAngle)
	return toDegrees(math.Asin(eq))
}

// SolarZenithAngle returns the Solar zenith angle in degrees calculated from the solar altitude angle,
// assuming (sin altitude = cos zenith) then (zenith = Acos(sin altitude))
func SolarZenithAngle(day int, latitude float64) float64 {
	sinAltitudeAngle := math.Sin(toRadians(SolarAltitudeAngle(day, latitude)))
	return toDegrees(math.Acos(sinAltitudeAngle))
}

func SolarAzimuthAngle(day int, latitude float64, hourAngle float64) float64 {
	declination := toRadians(SolarDeclination(day))
	solarAltitude := toRadians(SolarAltitudeAngle(day, latitude))
	hourAngle = toRadians(hourAngle)
	sinAz := (math.Cos(declination) * math.Sin(hourAngle)) / math.Cos(solarAltitude)
	return math.Asin(sinAz)
}

func SolarIncidenceAngle(day int, latitude float64, tiltAngle float64, surfaceAzimuth float64, hourAngleDeg float64) float64 {
	return 0
}

// DayLength returns the length of a day in hours
func DayLength(day int, latitude float64) float64 {
	declination := toRadians(SolarDeclination(day))
	latitude = toRadians(latitude)

	// Corrected angle for Sun's angular diameter and atmospheric refraction
	solarDepressionAngle := toRadians(-1.15)

	cosH := (math.Sin(solarDepressionAngle) - math.Sin(latitude)*math.Sin(declination)) / (math.Cos(latitude) * math.Cos(declination))

	if cosH > 1 {
		cosH = 1
	} else if cosH < -1 {
		cosH = -1
	}

	// Compute the hour angle at sunrise/sunset.
	hourAngle := math.Acos(cosH)

	dayLength := (2.0 * toDegrees(hourAngle)) / 15.0

	return dayLength
}

// SunriseAndSunset returns the sunrise and sunset time as hours in solar time. ACCURACY +- 5 min.
func SunriseAndSunset(day int, latitude, longitude float64, dst bool) (sunrise, sunset float64) {

	dsTime := 0
	if dst {
		dsTime = 1
	}
	// Calculate the day length in hours.
	dayLength := DayLength(day, latitude)
	equationOfTime := EquationOfTime(day)
	// Longitude of the standard meridian for the local time zone.
	sMeridian := float64(standardMeridian(longitude))

	// Calculate the solar noon in hours, considering equation of time and longitude correction.
	solarNoon := 12.0 + (((sMeridian-longitude)*4)+equationOfTime)/60.0

	// Calculate the sunrise and sunset times in hours.
	sunrise = (solarNoon - (dayLength / 2.0)) + float64(dsTime)
	sunset = (solarNoon + (dayLength / 2.0)) + float64(dsTime)

	return sunrise, sunset
}

func SunriseTime(day int, latitude float64, longitude float64) float64 {
	sunrise, _ := SunriseAndSunset(day, latitude, longitude, true)
	return sunrise
}

func SunsetTime(day int, latitude float64, longitude float64) float64 {
	_, sunset := SunriseAndSunset(day, latitude, longitude, true)
	return sunset
}

// toRadians converts an angle in degrees to radians
func toRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180.0)
}

// toDegrees converts an angle in radians to degrees
func toDegrees(radians float64) float64 {
	return radians * (180.0 / math.Pi)
}

func standardMeridian(longitude float64) int {
	ceil := math.Ceil(longitude/15) * 15
	floor := math.Floor(longitude/15) * 15

	if math.Abs(longitude-ceil) < math.Abs(longitude-floor) {
		return int(ceil)
	}
	return int(floor)
}
