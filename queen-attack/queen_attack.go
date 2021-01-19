package queenattack

import "errors"

var validRow = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
}
var validCol = map[string]int{
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
}

type point struct {
	x, y int
}

func (p *point) isEqual(b point) bool {
	return p.x == b.x && p.y == b.y
}

func newPoint(x, y int) *point {
	return &point{x, y}
}

func (p *point) moveUpLeft(cell int) *point {
	return newPoint(p.x+cell, p.y+cell)
}

func (p *point) moveDownLeft(cell int) *point {
	return newPoint(p.x-cell, p.y+cell)
}

var errInvalid = errors.New("invalid")

func fromDescriptiveNotation(d string) (point, error) {
	var out point
	if len(d) != 2 {
		return out, errInvalid
	}
	r, ok := validRow[d[0:1]]
	if !ok {
		return out, errInvalid
	}
	out.x = r
	r, ok = validCol[d[1:2]]
	if !ok {
		return out, errInvalid
	}
	out.y = r

	return out, nil
}

func CanQueenAttack(q1 string, q2 string) (bool, error) {
	var first, second point
	var err error
	if first, err = fromDescriptiveNotation(q1); err != nil {
		return false, err
	}
	if second, err = fromDescriptiveNotation(q2); err != nil {
		return false, err
	}
	if first.isEqual(second) {
		return false, errInvalid
	}
	// same row/col
	if first.x == second.x || first.y == second.y {
		return true, nil
	}

	if checkDiagonals(first, second) || checkDiagonals(second, first) {
		return true, nil
	}

	return false, nil
}

func checkDiagonals(first, second point) bool {
	for i := 1; (first.x+i) <= 8 && first.y+i <= 8; i++ {
		if first.moveUpLeft(i).isEqual(second) {
			return true
		}
	}
	for i := 1; (first.x-i) >= 0 && first.y+i <= 8; i++ {
		if first.moveDownLeft(i).isEqual(second) {
			return true
		}
	}

	return false
}
