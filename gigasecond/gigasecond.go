// Package gigasecond expose a function that calculate the gigasecond starting from a date
package gigasecond

import "time"

const (
	Nanosecond  time.Duration = 1
	Microsecond               = 1000 * Nanosecond
	Millisecond               = 1000 * Microsecond
	Second                    = 1000 * Millisecond
	GigaSecond 	          = 1000000000 * Second

)
// Calculate the moment when someone has lived for 10^9 seconds.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(GigaSecond)
}
