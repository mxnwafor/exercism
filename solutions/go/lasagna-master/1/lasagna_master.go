package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, prepTimePerLayer int) int {
    var layerLength = len(layers)
    if prepTimePerLayer == 0 {
        return layerLength * 2
    }
	 
	return layerLength * prepTimePerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
    var noodCounter = 0
    var sauceCounter = 0.0

    for i := 0; i < len(layers); i++ {
        if layers[i] == "noodles" {
            noodCounter++
        } else if layers[i] == "sauce" {
            sauceCounter++
        } else {
        	continue    
        } 
    }

    return noodCounter * 50, sauceCounter * 0.2
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList []string, recipe []string) {
    
    var lastFriendsItem = friendsList[len(friendsList) - 1]

    if len(recipe) == 0 {
        panic("recipe empty")
    }
    var newSlice = recipe[:len(recipe) - 1]

    newSlice = append(newSlice, lastFriendsItem)
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
    
    var scaledQuantities = make([]float64, len(quantities))
    
    for i:= 0; i < len(quantities); i++ {
        scaledQuantities[i] = quantities[i] * float64(portions) / 2
    }

    return scaledQuantities
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
