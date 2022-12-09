package units

const (
	FromGstoLbs = 0.00220462
	FromLbstoGs = 1 / FromGstoLbs
)

type weightUnit int

const (
	Pounds weightUnit = iota
	Kilograms
	Ounces
	Grams
)
