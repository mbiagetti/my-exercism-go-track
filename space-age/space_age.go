package space

type Planet interface{}

/*
var aaa = map[Planet] float64 {
	"Earth": float64(1),
}
*/
func Age(seconds float64, planet Planet) float64 {
	earth := seconds / float64(31557600)

	var m = make(map[Planet]float64)

	m["Earth"] = float64(1)
	m["Mercury"] = float64(0.2408467)
	m["Venus"] = float64(0.61519726)
	m["Mars"] = float64(1.8808158)
	m["Jupiter"] = float64(11.862615)
	m["Saturn"] = float64(29.447498)
	m["Uranus"] = float64(84.016846)
	m["Neptune"] = float64(164.79132)

	return earth / m[planet]
}
