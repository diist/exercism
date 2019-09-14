// package gigagsecond is the Exercism package for exercise gigasecond
package gigasecond

import (
	"time"
)

// AggGigasecond adds a gigasecond to a given date
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1e9)
}
