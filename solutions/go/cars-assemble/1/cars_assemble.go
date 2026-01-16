package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    return (successRate / float64(100)) * float64(productionRate)
	panic("CalculateWorkingCarsPerHour not implemented")
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
    return int(float64(productionRate / 60) * (successRate / float64(100)))
	panic("CalculateWorkingCarsPerMinute not implemented")
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
    if carsCount >= 10 {
        var remainder = carsCount % 10
        var group = carsCount / 10
        return uint((group * 95000) + (remainder * 10000))
    } else {
        return uint(carsCount * 10000)
    }
	panic("CalculateCost not implemented")
}
