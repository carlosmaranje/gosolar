package gosolar

import (
	"os"
	"testing"
)

var sc *SolarCalculation
var err error

func TestMain(m *testing.M) {

	latitude := 23.0975036         // float Degrees
	longitude := -82.4206579       // float Degrees
	date := "2023-01-01"           // string "YYYY-MM-DD"
	dayTime := 0.5                 // float time of the day/24
	timeZone := "America/New_York" // string TimezoneID

	sc, err = Calculator(latitude, longitude, dayTime, timeZone, date)

	if err != nil {
		os.Exit(1)
	}
	os.Exit(m.Run())
}
