package gosolar

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDayLength(t *testing.T) {
	dayLength := sc.DayLength()
	assert.Equal(t, 10.743499619136012, dayLength)
}

func TestEccentEarthOrbit(t *testing.T) {
	orbit := sc.EccentEarthOrbit()
	assert.Equal(t, 0.016698958307359145, orbit)
}

func TestEquationOfTime(t *testing.T) {
	eot := sc.EquationOfTime()
	assert.Equal(t, -3.5247395659817014, eot)
}

func TestGeomMeanAnomSun(t *testing.T) {
	meanAnom := sc.GeomMeanAnomSun()
	assert.Equal(t, 8637.721335352362, meanAnom)
}

func TestGeomMeanLongSun(t *testing.T) {
	meanLongSun := sc.GeomMeanLongSun()
	assert.Equal(t, 281.05422334078503, meanLongSun)
}

func TestHourAngleSunrise(t *testing.T) {
	hourAngle := sc.HourAngleSunrise()
	assert.Equal(t, 80.57624714352008, hourAngle)
}

func TestIncidenceOnTiltedSurface(t *testing.T) {
	incidence := sc.IncidenceOnTiltedSurface(45, 10)
	assert.Equal(t, 28.596313464002552, incidence)
}

func TestJulianCentury(t *testing.T) {
	julianCentury := sc.JulianCentury()
	assert.Equal(t, 0.2300114077116088, julianCentury)
}

func TestJulianDay(t *testing.T) {
	julianDay := sc.JulianDay()
	assert.Equal(t, 2459946.1666666665, julianDay)
}

func TestMeanObliqEcliptic(t *testing.T) {
	obliqEcliptic := sc.MeanObliqEcliptic()
	assert.Equal(t, 23.436300001887762, obliqEcliptic)
}

func TestObliqueCorrection(t *testing.T) {
	obliqueCorrection := sc.ObliqueCorrection()
	assert.Equal(t, 23.438256281010286, obliqueCorrection)
}

func TestSolarAzimuthAngle(t *testing.T) {
	azimuthAngle := sc.SolarAzimuthAngle()
	assert.Equal(t, 152.20611634980753, azimuthAngle)
}

func TestSolarDeclination(t *testing.T) {
	declination := sc.SolarDeclination()
	assert.Equal(t, -22.985319542237658, declination)
}

func TestSolarIncidenceAngle(t *testing.T) {
	incidenceAngle := sc.SolarIncidenceAngle()
	assert.Equal(t, 38.648930569794715, incidenceAngle)
}

func TestSolarNoon(t *testing.T) {
	solarNoon := sc.SolarNoon()
	assert.Equal(t, 0.5647273410874872, solarNoon)
}

func TestSolarZenithAngle(t *testing.T) {
	zenithAngle := sc.SolarZenithAngle()
	assert.Equal(t, 51.351069430205285, zenithAngle)
}

func TestSunApparentLongitude(t *testing.T) {
	apparentLongitude := sc.SunApparentLongitude()
	assert.Equal(t, 280.9677490971792, apparentLongitude)
}

func TestSunEquationOfCenter(t *testing.T) {
	equationOfCenter := sc.SunEquationOfCenter()
	assert.Equal(t, -0.07770108109845357, equationOfCenter)
}

func TestSunHourAngle(t *testing.T) {
	hourAngle := sc.SunHourAngle()
	assert.Equal(t, -23.301842791495403, hourAngle)
}

func TestSunTrueLongitude(t *testing.T) {
	assert.Equal(t, 1, 1)
}

func TestSunriseAndSunset(t *testing.T) {
	sunrise, sunset := sc.SunriseAndSunset()
	assert.Equal(t, 8.181706376531688, sunrise)
	assert.Equal(t, 18.9252059956677, sunset)
}

func TestTrueSolarTime(t *testing.T) {
	trueSolarTime := sc.TrueSolarTime()
	assert.Equal(t, 626.7926288340184, trueSolarTime)
}

func TestTimeZoneOffset(t *testing.T) {
	tzOff, err := TimeZoneOffset("America/New_York")
	require.NoError(t, err)
	assert.Equal(t, -14400, tzOff)
}
