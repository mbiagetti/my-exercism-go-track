package clock

import "fmt"

// Clock struct represents a clock
type Clock struct {
	m int
}

// OneHourInMinute represents an hour in minutes
const OneHourInMinute = 60

func (c Clock) String() string {
	return fmt.Sprintf(
		"%02d:%02d",
		c.m/OneHourInMinute,
		c.m%OneHourInMinute,
	)
}

// Add sum minutes to the current time
func (c Clock) Add(minutes int) Clock {
	return New(0, c.m+minutes)
}

// Subtract remove minutes to the current time
func (c Clock) Subtract(minutes int) Clock {
	return New(0, c.m-minutes)
}

// New create a new clock with the provided hours and minutes
func New(h int, m int) Clock {
	totMin := m + h*60

	for totMin < 0 {
		totMin += 1440
	}

	return Clock{totMin % 1440}
}
