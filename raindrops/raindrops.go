package raindrops

import (
	"math"
	"strconv"
)

func Convert(num int) string {
	ret := ""

	if math.Mod(float64(num), float64(3)) == 0 {
		ret += "Pling"
	}

	if math.Mod(float64(num), float64(5)) == 0 {
		ret += "Plang"
	}

	if math.Mod(float64(num), float64(7)) == 0 {
		ret += "Plong"
	}

	if len(ret) == 0 {
		return strconv.Itoa(num)
	}

	return ret
}
