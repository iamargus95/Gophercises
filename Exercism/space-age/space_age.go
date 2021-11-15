package space

type Planet string

const Earth = 31557600.0

var PlanetSeconds = map[Planet]float64{
	"Earth":   Earth,
	"Mercury": Earth * 0.2408467,
	"Venus":   Earth * 0.61519726,
	"Mars":    Earth * 1.8808158,
	"Jupiter": Earth * 11.862615,
	"Saturn":  Earth * 29.447498,
	"Uranus":  Earth * 84.016846,
	"Neptune": Earth * 164.79132,
}

func Age(seconds float64, planet Planet) float64 {
	return seconds / PlanetSeconds[planet]
}
