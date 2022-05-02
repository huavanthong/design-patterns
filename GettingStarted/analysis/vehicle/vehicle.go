package vehicle

// Step 1: Define a interface
type Vehicle interface {
	Structure() []string // Common Method
	Speed() string
}

// A global variable at the top of program
type Car string

// Step 2: Implement interface inheritance from Vehicle()
func (c Car) Structure() []string {
	var parts = []string{"ECU", "Engine", "Air Filters", "Wipers", "Gas Task"}
	return parts
}

func (c Car) Speed() string {
	return "200 Km/Hrs"
}
