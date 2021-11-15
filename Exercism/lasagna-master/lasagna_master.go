package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, minutes int) int {
	if minutes != 0 {
		return len(layers) * minutes
	}
	return len(layers) * 2
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	noodles := 0
	sauce := 0.0
	for _, v := range layers {
		if v == "noodles" {
			noodles++
		}
		if v == "sauce" {
			sauce++
		}
	}
	return noodles * 50, 0.2 * sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(frientsList, myList []string) []string {
	ingredient := frientsList[len(frientsList)-1]
	result := myList
	result = append(result, ingredient)
	return result
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {

	var scaledQuantities []float64
	for _, v := range quantities {
		scaledQuantities = append(scaledQuantities, v*(float64(portions)/2.0))
	}
	return scaledQuantities
}
