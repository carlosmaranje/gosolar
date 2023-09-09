# gosolar
A go package with info related to the sun movement

Calculations made by this package are done using NOAA's methods. You can find more information on their website: 

[Solar calculations details](https://gml.noaa.gov/grad/solcalc/calcdetails.html)

[PDF version](https://gml.noaa.gov/grad/solcalc/solareqns.PDF)

This is still a WIP. Feel free to open PRs, fork and contribute

## Usage
You should use the method `Calculator()` to return a `SolarCalculation` object that will provide access to all the other
methods in the module.

```go
package gosolar

import (
	"fmt"
	"log"
	"github.com/karlsmaranjs/gosolar"
)

func main() string {
	latitude := 35.0               // float Degrees
	longitude := -80.37486         // float Degrees
	date := "2023-06-16"           // string "YYYY-MM-DD"
	dayTime := 0.64                // float time of the day/24. 12:00:00 PM = 0.5
	timeZone := "America/New_York" // string Timezone ID

	sun, err := Calculator(latitude, longitude, dayTime, timeZone, date)
	if err != nil {
		log.Fatalf("Error calculating sun info: %v", err)
	}
	// Returns the solar declination
	declination := sun.SolarDeclination()

	// Returns the incidence angle on a surface tilted 35 degrees from the ground and pointing south
	tiltedAngle := sun.IncidenceOnTiltedSurface(35, 180)

	message := fmt.Sprintf("declination: %f, Incidence on roof: %f", declination, tiltedAngle)
	return message
}

```

Please notice that `Calculator()` expects a valid string as a `timeZone` e.g. "America/New_York". This allows 
to determine the current offset for that `timeZone` including daylight saving time (DST). `Calculator()` then will determine 
the timezone offset using `TimeZoneOffset()` and use this value (`float64`) to initialize a `SolarCalculation` object. 

### Disclaimer
This library is not associated in any way, shape or form with NOAA

