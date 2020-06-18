package space

type Planet string

func Age(seconds float64, planet Planet) float64  {
	var years float64
	switch planet {
	case "Mercury":
		years = secondsToYears(seconds) / 0.2408467
	case "Venus":
		years = secondsToYears(seconds) / 0.61519726
	case "Earth":
		years = secondsToYears(seconds) / 1
	case "Mars":
		years = secondsToYears(seconds) / 1.8808158
	case "Jupiter":
		years = secondsToYears(seconds) / 11.862615
	case "Saturn":
		years = secondsToYears(seconds) / 29.447498
	case "Uranus":
		years = secondsToYears(seconds) / 84.016846
	case "Neptune":
		years = secondsToYears(seconds) / 164.79132
	}
	return years
}

func secondsToYears (seconds float64) float64 {
	return seconds / 31557600
}