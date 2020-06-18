package triangle

import "math"

type Kind string

const (
    NaT = "Nat"// not a triangle
    Equ = "Equ"// equilateral
    Iso = "Iso"// isosceles
    Sca = "Sca"// scalene
)

func KindFromSides(a, b, c float64) Kind {
	var k Kind
	if math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) || math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) || a <= 0 || b <= 0 || c <= 0 || a + b < c || a + c < b || b + c < a {
		k = NaT
	} else if a == b && b == c{
		k = Equ
	} else if a == b || a == c || b == c {
		k = Iso
	} else {
		k = Sca
	}
	return k
}
