package space

type Planet interface{}

func Age(seconds float64, planet Planet) float64 {
	var age float64

	switch planet {
	case "Earth":
		age = float64(1)
	case "Mercury":
		age = float64(0.2408467)
	case "Venus":
		age = float64(0.61519726)
	case "Mars":
		age = float64(1.8808158)
	case "Jupiter":
		age = float64(11.862615)
	case "Saturn":
		age = float64(29.447498)
	case "Uranus":
		age = float64(84.016846)
	case "Neptune":
		age = float64(164.79132)

	}

	return seconds / float64(31557600) / age
}
