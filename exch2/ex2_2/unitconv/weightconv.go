package unitconv

import "fmt"

type Pounds float64
type Kilograms float64

func (p Pounds) String() string {
	return fmt.Sprintf("%f lbs", p)
}

func (k Kilograms) String() string {
	return fmt.Sprintf("%f Kg", k)
}

// PToK converts Pounds to Kilograms
func PToK(p Pounds) Kilograms {
	return Kilograms(p * 0.45359237)
}

// KToP converts Kilograms to Pounds
func KToP(k Kilograms) Pounds {
	return Pounds(k / 0.45359237)
}
