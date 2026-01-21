package interest
// import "math"

// InterestRate returns the interest rate for the provided balance.
func InterestRate(balance float64) float32 {
    if balance < 0 {
        return float32(3.213)
    } else if balance >= 0 && balance < 1000 {
        return float32(0.5)
    } else if balance >= 1000 && balance < 5000 {
        return float32(1.621)
    } else {
        return float32(2.475)
    }
}

// Interest calculates the interest for the provided balance.
func Interest(balance float64) float64 {
    return float64(InterestRate(balance)) / 100 * balance
}

// AnnualBalanceUpdate calculates the annual balance update, taking into account the interest rate.
func AnnualBalanceUpdate(balance float64) float64 {
    return Interest(balance) + balance
}

// YearsBeforeDesiredBalance calculates the minimum number of years required to reach the desired balance.
func YearsBeforeDesiredBalance(balance, targetBalance float64) int {
    var years = 0
    
    if balance == targetBalance {
        return years
    }
    
    for i := 0; balance < targetBalance; i++ {
        balance = AnnualBalanceUpdate(balance)
        years++
    }
    
    return years
}
