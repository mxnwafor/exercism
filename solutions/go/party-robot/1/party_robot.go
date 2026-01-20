package partyrobot
import "fmt"

// Welcome greets a person by name.
func Welcome(name string) string {
    return fmt.Sprintf("Welcome to my party, %s!", name)
	panic("Please implement the Welcome function")
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
    return fmt.Sprintf("Happy birthday %s! You are now %d years old!", name, age)
	panic("Please implement the HappyBirthday function")
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor string, direction string, distance float64) string {
    width := 3
    formattedTable := fmt.Sprintf("%0*d", width, table)
    return fmt.Sprintf("%s\nYou have been assigned to table %s. Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.", Welcome(name), formattedTable, direction, distance, neighbor)
	panic("Please implement the AssignTable function")
}
