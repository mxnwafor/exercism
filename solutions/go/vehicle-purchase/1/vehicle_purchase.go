package purchase
import "fmt"

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
    if kind == "car" || kind == "truck" {
        return true
    } else {
        return false
    }
	panic("NeedsLicense not implemented")
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
    if option1 <= option2 {
        return fmt.Sprintf("%s is clearly the better choice.", option1)
    } else {
        return fmt.Sprintf("%s is clearly the better choice.", option2)
    }
	panic("ChooseVehicle not implemented")
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
    if age < float64(3) {
        return float64(0.8 * float64(originalPrice))
    } else if age >= float64(3) && age < float64(10) {
        return float64(0.7 * float64(originalPrice))
    } else {
        return float64(0.5 * float64(originalPrice))
    }
	panic("CalculateResellPrice not implemented")
}
