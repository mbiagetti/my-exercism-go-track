package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

// Robot Manage robot factory settings
type Robot struct {
	name string
}

const maxCombination = 26 * 26 * 10 * 10 * 10

var nameAlreadySeen = map[string]bool{}

// Name return the name of the robot
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	if len(nameAlreadySeen) >= maxCombination {
		return "", errors.New("max combination limit reach")
	}
	r.name = newName()
	for nameAlreadySeen[r.name] {
		r.name = newName()
	}
	nameAlreadySeen[r.name] = true
	return r.name, nil
}

// Reset reset a robot to its factory settings
func (r *Robot) Reset() {
	r.name = ""
}
func newName() string {
	r1 := rand.Intn(26) + 'A'
	r2 := rand.Intn(26) + 'A'
	num := rand.Intn(1000)
	return fmt.Sprintf("%c%c%03d", r1, r2, num)
}
