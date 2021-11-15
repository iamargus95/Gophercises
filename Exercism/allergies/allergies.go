package allergies

var allergenes = map[string]uint{
	"eggs":         1,
	"peanuts":      2,
	"shellfish":    4,
	"strawberries": 8,
	"tomatoes":     16,
	"chocolate":    32,
	"pollen":       64,
	"cats":         128,
}

func Allergies(allergies uint) (result []string) {
	for allergen, value := range allergenes {
		if allergies&value == value {
			result = append(result, allergen)
		}
	}
	return result
}
func AllergicTo(allergies uint, allergen string) bool {
	return allergies&allergenes[allergen] == allergenes[allergen]
}
