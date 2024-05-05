package lasagna

// Define the 'PreparationTime()' function
func PreparationTime(layers []string, prepLayer int) (prepTime int) {
	if prepLayer == 0 {
		prepTime = 2 * len(layers)
	} else {
		prepTime = prepLayer * len(layers)
	}
	return prepTime
}

// Define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	literCount := 0
	gramCount := 0
	for _, layer := range layers {
		switch layer {
		case "sauce":
			literCount++
			sauce = float64(literCount) * 0.2
		case "noodles":
			gramCount++
			noodles = gramCount * 50
		}
	}
	return noodles, sauce
}

// Define the 'AddSecretIngredient()' function
func AddSecretIngredient(sercertList, myList []string) {
	if len(sercertList) > 0 && len(myList) > 0 {
		(myList)[len(myList)-1] = (sercertList)[len(sercertList)-1]
	}
}

// Define the 'ScaleRecipe()' function
func ScaleRecipe(quanitities []float64, portions int) (scaledPortions []float64) {
	for _, portion := range quanitities {
		scaledPortions = append(scaledPortions, float64(portions)/2.0*portion)
	}
	return scaledPortions
}
