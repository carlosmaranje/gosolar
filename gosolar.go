package gosolar

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"time"
)

type SolarCalculation struct {
	latitude       float64 // float Degrees
	longitude      float64 // float Degrees
	date           string  // string "YYYY-MM-DD"
	dayTime        float64 // float time of the day/24
	timeZoneOffset float64 // float timezone UTC offset in seconds
}

// Calculator acts as a constructor for the module. This allows to perform some validations before implementing solarCalculation struct
func Calculator(latitude, longitude, dayTime float64, timeZone, date string) (*SolarCalculation, error) {

	tz, _ := TimeZoneOffset(timeZone)

	sc := &SolarCalculation{
		latitude:       latitude,
		longitude:      longitude,
		date:           date,
		dayTime:        dayTime,
		timeZoneOffset: float64(tz) / 3600,
	}

	if err := sc.validate(); err != nil {
		return nil, err
	}
	return sc, nil
}

// Setters

func (sc *SolarCalculation) SetLatitude(lat float64) error {
	if lat < -90 || lat > 90 {
		return errors.New("latitude must be between -90 and 90 degrees")
	}
	sc.latitude = lat
	return nil
}

func (sc *SolarCalculation) SetLongitude(lon float64) error {
	if lon < -180 || lon > 180 {
		return errors.New("longitude must be between -180 and 180 degrees")
	}
	sc.longitude = lon
	return nil
}

func (sc *SolarCalculation) SetDate(date string) error {
	if matched, _ := regexp.MatchString(`\d{4}-\d{2}-\d{2}`, date); !matched {
		return errors.New("date must be in format YYYY-MM-DD")
	}
	sc.date = date
	return nil
}

func (sc *SolarCalculation) SetDayTime(dayTime float64) error {
	if dayTime < 0 || dayTime > 1 {
		return errors.New("dayTime must be between 0 and 1")
	}
	sc.dayTime = dayTime
	return nil
}

func (sc *SolarCalculation) SetTimeZone(timeZone string) error {
	tzOffset, err := TimeZoneOffset(timeZone)
	if err != nil {
		return err
	}
	sc.timeZoneOffset = float64(tzOffset) / 3600
	return nil
}

// Getters

func (sc *SolarCalculation) GetLatitude() float64 {
	return sc.latitude
}

func (sc *SolarCalculation) GetLongitude() float64 {
	return sc.longitude
}

func (sc *SolarCalculation) GetDate() string {
	return sc.date
}

func (sc *SolarCalculation) GetTimeZone() float64 {
	return sc.timeZoneOffset
}

func (sc *SolarCalculation) GetDayTime() float64 {
	return sc.dayTime
}

func (sc *SolarCalculation) GetTimeZoneOffset() float64 {
	return sc.timeZoneOffset
}

// JulianDay returns the Julian Day number for a valid date given in format YYYY-MM-DD. The result can be corrected for
// time of the day (0 <= dayTime <=1) and timeZoneOffset (UTC)
func (sc *SolarCalculation) JulianDay() float64 {
	startEpoch := 2415020.5
	epoch := time.Date(1900, time.January, 1, 0, 0, 0, 0, time.UTC)
	parsedDate, err := time.Parse("2006-01-02", sc.date)

	if err != nil {
		return 0
	}

	// nanoseconds
	elapsed := parsedDate.Sub(epoch)
	days := elapsed.Hours() / 24

	return days + startEpoch + (sc.dayTime - float64(sc.timeZoneOffset)/24)
}

// JulianCentury returns the value of Julian Century
func (sc *SolarCalculation) JulianCentury() float64 {
	return (sc.JulianDay() - 2451545) / 36525
}

// GeomMeanLongSun returns The geometric mean longitude of the Sun, in degrees. For any given date in format YYYY-MM-DD.
// The result can be corrected for time of the day (0 <= dayTime <=1) and timeZoneOffset (UTC)
func (sc *SolarCalculation) GeomMeanLongSun() float64 {
	jCent := sc.JulianCentury()
	return math.Mod(280.46646+(jCent*(36000.76983+jCent*0.0003032)), 360)
}

// GeomMeanAnomSun returns The geometric mean anomaly of the Sun, in degrees. For any given date in format YYYY-MM-DD.
// The result can be corrected for time of the day (0 <= dayTime <=1) and timeZoneOffset (UTC)
func (sc *SolarCalculation) GeomMeanAnomSun() float64 {
	jCent := sc.JulianCentury()

	return 357.52911 + jCent*(35999.05029-0.0001537*jCent)
}

// EccentEarthOrbit returns The eccentricity of the Earth's orbit, in degrees. For any given date in format YYYY-MM-DD.
// The result can be corrected for time of the day (0 <= dayTime <=1) and timeZoneOffset (UTC)
func (sc *SolarCalculation) EccentEarthOrbit() float64 {
	jCent := sc.JulianCentury()

	return 0.016708634 - jCent*(0.000042037+0.0000001267*jCent)
}

// EquationOfTime Calculates the value for the equation of time for any given day of the year as found
// in the official NOAA website
func (sc *SolarCalculation) EquationOfTime() float64 {

	geomMeanLongSun := sc.GeomMeanLongSun()
	eccentEarthOrbit := sc.EccentEarthOrbit()
	varY := 0.043031509

	gmlRad := 2 * sc.toRadians(geomMeanLongSun)
	gmaRad := sc.toRadians(sc.GeomMeanAnomSun())

	// Mean longitude
	gmlComp := varY * math.Sin(gmlRad)
	// Geometric mean anomaly
	gmaComp := 2 * eccentEarthOrbit * math.Sin(gmaRad)
	// Eccentricity
	eccComp := 4 * eccentEarthOrbit * varY * math.Sin(gmaRad) * math.Cos(gmlRad)

	varYComp := 0.5 * math.Pow(varY, 2) * math.Sin(4*sc.toRadians(geomMeanLongSun))
	eccSqComp := 1.25 * math.Pow(eccentEarthOrbit, 2) * math.Sin(2*gmaRad)

	formula := 4 * sc.toDegrees(gmlComp-gmaComp+eccComp-varYComp-eccSqComp)

	return formula
}

// SolarNoon returns the solar noon time in hours
func (sc *SolarCalculation) SolarNoon() float64 {
	return (720 - 4*sc.longitude - sc.EquationOfTime() + float64(sc.timeZoneOffset)*60) / 1440
}

// SunEquationOfCenter returns the angular difference between the actual position
// of the sun in its elliptical orbit and the position it would occupy if its motion were uniform
func (sc *SolarCalculation) SunEquationOfCenter() float64 {
	meanAnomaly := sc.GeomMeanAnomSun()
	jC := sc.JulianCentury()

	term1 := math.Sin(sc.toRadians(meanAnomaly)) * (1.914602 - jC*(0.004817+0.000014*jC))
	term2 := math.Sin(sc.toRadians(2*meanAnomaly)) * (0.019993 - 0.000101*jC)
	term3 := math.Sin(sc.toRadians(3*meanAnomaly)) * 0.000289

	return term1 + term2 + term3
}

// SunTrueLongitude returns the Sun's true longitude, in degrees
func (sc *SolarCalculation) SunTrueLongitude() float64 {
	return sc.GeomMeanLongSun() + sc.SunEquationOfCenter()
}

// TrueSolarTime returns the true solar time in minutes
func (sc *SolarCalculation) TrueSolarTime() float64 {
	result := sc.dayTime*1440 + sc.EquationOfTime() + 4*sc.longitude - 60*sc.timeZoneOffset
	return math.Mod(result, 1440)
}

// SunApparentLongitude returns the Sun's apparent longitude, in degrees
func (sc *SolarCalculation) SunApparentLongitude() float64 {
	sunTrueLongitude := sc.SunTrueLongitude()
	jC := sc.JulianCentury()
	return sunTrueLongitude - 0.00569 - 0.00478*math.Sin(sc.toRadians(125.04-1934.136*jC))
}

// MeanObliqEcliptic returns the mean inclination of Earth's equator with respect to the ecliptic
func (sc *SolarCalculation) MeanObliqEcliptic() float64 {
	jC := sc.JulianCentury()
	term2 := 26.0 + ((21.448 - jC*(46.815+jC*(0.00059-jC*0.001813))) / 60.0)

	return 23.0 + term2/60.0
}

// ObliqueCorrection returns the oblique correction
func (sc *SolarCalculation) ObliqueCorrection() float64 {
	jC := sc.JulianCentury()
	moe := sc.MeanObliqEcliptic()
	angle := 125.04 - 1934.136*jC

	return moe + 0.00256*math.Cos(sc.toRadians(angle))
}

// SolarDeclination returns the declination in degrees
func (sc *SolarCalculation) SolarDeclination() float64 {
	oblCorr := sc.toRadians(sc.ObliqueCorrection())
	sunAppLon := sc.toRadians(sc.SunApparentLongitude())
	declination := math.Asin(math.Sin(oblCorr) * math.Sin(sunAppLon))

	return sc.toDegrees(declination)
}

// SunHourAngle returns the hour angle of the sun in degrees
func (sc *SolarCalculation) SunHourAngle() float64 {
	return (sc.TrueSolarTime() / 4) - 180
}

// HourAngleSunrise returns the hour angle of the sun at sunrise in degrees
func (sc *SolarCalculation) HourAngleSunrise() float64 {
	declination := sc.toRadians(sc.SolarDeclination())
	latitude := sc.toRadians(sc.latitude)

	num := math.Cos(sc.toRadians(90.833))
	cos := math.Cos(latitude) * math.Cos(declination)
	tang := math.Tan(latitude) * math.Tan(declination)
	hourAngle := math.Acos(num/cos - tang)

	return sc.toDegrees(hourAngle)
}

// SolarZenithAngle returns the Solar zenith angle in degrees calculated from the solar altitude angle,
// assuming (sin altitude = cos zenith) then (zenith = Acos(sin altitude))
func (sc *SolarCalculation) SolarZenithAngle() float64 {
	declination := sc.toRadians(sc.SolarDeclination())
	latitude := sc.toRadians(sc.latitude)
	hourAngle := sc.toRadians(sc.SunHourAngle())

	sin := math.Sin(latitude) * math.Sin(declination)
	cos := math.Cos(latitude) * math.Cos(declination) * math.Cos(hourAngle)

	return sc.toDegrees(math.Acos(sin + cos))
}

// SolarAzimuthAngle returns the Solar azimuth angle in degrees calculated from the solar altitude angle,
// assuming (sin altitude = cos zenith) then (zenith = Acos(sin altitude))
func (sc *SolarCalculation) SolarAzimuthAngle() float64 {
	var mod float64
	hourAngle := sc.SunHourAngle()
	latitude := sc.toRadians(sc.latitude)
	zenithAngle := sc.toRadians(sc.SolarZenithAngle())
	declination := sc.toRadians(sc.SolarDeclination())

	num := (math.Sin(latitude) * math.Cos(zenithAngle)) - math.Sin(declination)
	cosSin := math.Cos(latitude) * math.Sin(zenithAngle)
	formula := sc.toDegrees(math.Acos(num / cosSin))

	if hourAngle > 0 {
		mod = formula + 180
	} else {
		mod = 540 - formula
	}

	return math.Mod(mod, 360)
}

// Returns the SolarIncidenceAngle in degrees
func (sc *SolarCalculation) SolarIncidenceAngle() float64 {

	return 90 - sc.SolarZenithAngle()
}

// IncidenceOnTiltedSurface returns the Solar incidence angle on a tilted surface in degrees
func (sc *SolarCalculation) IncidenceOnTiltedSurface(surfaceAngle, surfaceAzimuth float64) float64 {
	latitude := sc.toRadians(sc.latitude)
	declination := sc.toRadians(sc.SolarDeclination())
	azimuth := sc.toRadians(surfaceAzimuth)
	surfaceAngle = sc.toRadians(surfaceAngle)
	hourAngle := sc.toRadians(sc.SunHourAngle())

	seasonalTilt := math.Sin(latitude) * math.Sin(declination) * math.Cos(surfaceAngle)
	azmTerm := math.Cos(latitude) * math.Sin(declination) * math.Cos(azimuth) * math.Sin(surfaceAngle)
	hourTerm := math.Cos(latitude) * math.Cos(declination) * math.Cos(hourAngle) * math.Cos(surfaceAngle)
	hourAzim := math.Sin(latitude) * math.Cos(declination) * math.Cos(hourAngle) * math.Sin(surfaceAngle) * math.Cos(azimuth)
	declAzim := math.Cos(declination) * math.Sin(hourAngle) * math.Sin(surfaceAngle) * math.Sin(azimuth)

	cosAng := seasonalTilt - azmTerm + hourTerm + hourAzim + declAzim
	angle := math.Acos(cosAng)
	return sc.toDegrees(angle)
}

// SunriseAndSunset returns the sunrise and sunset time as hours in solar time.
func (sc *SolarCalculation) SunriseAndSunset() (sunrise, sunset float64) {
	solarNoon := sc.SolarNoon()
	hourAngle := sc.HourAngleSunrise()

	sunrise = (solarNoon*360 - hourAngle) / 15
	sunset = (solarNoon*360 + hourAngle) / 15

	return sunrise, sunset
}

// DayLength returns the length of a day in hours
func (sc *SolarCalculation) DayLength() float64 {
	sunrise, sunset := sc.SunriseAndSunset()
	dayLength := sunset - sunrise

	return dayLength
}

// SunriseTime returns the sunrise time in hours
func (sc *SolarCalculation) SunriseTime() float64 {
	sunrise, _ := sc.SunriseAndSunset()
	return sunrise
}

// SunsetTime returns the sunset time in hours
func (sc *SolarCalculation) SunsetTime() float64 {
	_, sunset := sc.SunriseAndSunset()
	return sunset
}

// TimeZoneOffset returns the offset in seconds for a given TimeZone id
func TimeZoneOffset(timeZoneId string) (int, error) {
	location, err := time.LoadLocation(timeZoneId)
	if err != nil {
		fmt.Println("error:", err)
		return -1, err
	}

	t := time.Now().In(location)
	_, offset := t.Zone()

	return offset, nil
}

// toRadians converts an angle in degrees to radians
func (sc *SolarCalculation) toRadians(degrees float64) float64 {
	return degrees * (math.Pi / 180.0)
}

// toDegrees converts an angle in radians to degrees
func (sc *SolarCalculation) toDegrees(radians float64) float64 {
	return radians * (180.0 / math.Pi)
}

// standardMeridian returns the standard meridian for a given longitude
func (sc *SolarCalculation) standardMeridian(longitude float64) int {
	ceil := math.Ceil(longitude/15) * 15
	floor := math.Floor(longitude/15) * 15

	if math.Abs(longitude-ceil) < math.Abs(longitude-floor) {
		return int(ceil)
	}
	return int(floor)
}

// dayOfYear returns the year day, out of 365, from a properly formatted string date
func (sc *SolarCalculation) dayOfYear(date string) int {
	const dateFormat = "2006-01-02"
	t, err := time.Parse(dateFormat, date)
	if err != nil {
		return -1
	}

	day := t.Day()
	month := t.Month()
	year := t.Year()

	timeDate := time.Date(year, month, day, 0, 0, 0, 0, time.UTC)

	return timeDate.YearDay()
}

// toDateFormatted returns a formatted string date from a year day. If -1 is used as argument, then the current day
// and/or year will be used
func (sc *SolarCalculation) toDateFormatted(day int, year int) string {
	const dateFormat = "2006-01-02"

	if day == -1 || day > 366 {
		day = time.Now().YearDay()
	}
	if year == -1 {
		year = time.Now().Year()
	}

	startOfYear := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
	date := startOfYear.AddDate(0, 0, day-1)
	fmt.Println(date.Format(dateFormat))

	return date.Format(dateFormat)
}

func (sc *SolarCalculation) roundTo(value float64, decimals int) float64 {
	pow := math.Pow(10, float64(decimals))
	return math.Round(value*pow) / pow
}

// validate performs some validations on the SolarCalculation struct
// validations are: latitude between -90 and 90, longitude between -180 and 180, date in format YYYY-MM-DD,
func (sc *SolarCalculation) validate() error {
	// Validate latitude
	if sc.latitude < -90 || sc.latitude > 90 {
		return errors.New("invalid latitude: must be between -90 and 90")
	}

	// Validate longitude
	if sc.longitude < -180 || sc.longitude > 180 {
		return errors.New("invalid longitude: must be between -180 and 180")
	}

	// Validate date
	dateMatch, _ := regexp.MatchString(`\d{4}-\d{2}-\d{2}`, sc.date)
	if !dateMatch {
		return errors.New("invalid date format: must be YYYY-MM-DD")
	}
	_, err := time.Parse("2006-01-02", sc.date)
	if err != nil {
		return errors.New("invalid date: " + err.Error())
	}

	// Validate timeZoneOffset
	if sc.timeZoneOffset < -12 || sc.timeZoneOffset > 14 {
		return errors.New("invalid time zone: must be between -12 and 14")
	}

	// Validate dayTime
	if sc.dayTime < 0 || sc.dayTime > 1 {
		return errors.New("invalid dayTime: must be between 0 and 1")
	}

	return nil
}
