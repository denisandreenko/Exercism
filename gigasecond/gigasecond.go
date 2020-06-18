package gigasecond

import (
	"math"
	"time"
)

func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Second * time.Duration(math.Pow(10, 9)))
	return t
}
