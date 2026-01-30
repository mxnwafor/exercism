package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
    gross_map := make(map[string]int)
    gross_map["quarter_of_a_dozen"] = 3
    gross_map["half_of_a_dozen"] = 6
    gross_map["dozen"] = 12
    gross_map["small_gross"] = 120
    gross_map["gross"] = 144
    gross_map["great_gross"] = 1728

    return gross_map
}

// NewBill creates a new bill.
func NewBill() map[string]int {
    newBill := make(map[string]int)
    return newBill
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
    value, exists := units[unit]
    
    if !exists {
        return false
    } else {
        bill[item] += value
        return true
    }
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
    _, itemExists := bill[item]
    _, unitExists := units[unit]

    if !itemExists {
        return false
    }
    if !unitExists {
        return false
    }
    
    var new_quantity = bill[item] - units[unit]
    bill[item] -= units[item]
    
    if new_quantity < 0 {
        return false
    } else if new_quantity == 0 {
        delete(bill, item)
        return true
    }
    
    bill[item] = new_quantity
    return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
    value, exists := bill[item]
    if !exists {
        return 0, false
    } else {
        return value, exists
    }
}
