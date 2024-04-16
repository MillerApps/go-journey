package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return successRate / 100 * float64(productionRate)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	// 1105 / 60 * .90
	// perMin := productionRate / 60
	// perMinSuccess := float64(perMin) * (successRate / 100) 
	// return int(perMinSuccess) 

	perHourSuccess := CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(perHourSuccess) / 60
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	costPerCar := 10000
	groupCost := 95000

	groups := carsCount / 10
	remainingCars := carsCount % 10

	return uint(groups) * uint(groupCost) + uint(remainingCars) * uint(costPerCar) 
}
