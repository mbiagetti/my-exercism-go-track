// Package triangle should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package triangle

import (
	"math"
)

type Kind int

const (
	_   = iota
	NaT // not a triangle
	Equ // equilateral
	Iso // isosceles
	Sca // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	switch {
	case !isValid(a) || !isValid(b) || !isValid(c):
		return NaT

	case (a+b) < c || (b+c) < a || (a+c) < b:
		return NaT

	case a == b && a == c:
		return Equ

	case a == b || a == c || b == c:
		return Iso

	default:
		return Sca
	}
}

// isValid return true if the float is a positive and valid value
func isValid(a float64) bool {
	return a > 0 && !math.IsNaN(a) && !math.IsInf(a, 1) && !math.IsInf(a, -1)
}