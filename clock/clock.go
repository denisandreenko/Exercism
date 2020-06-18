package clock

import (
	"fmt"
)

type Clock int
const dayMinutes = 24 * 60

func New(hours int, minutes int) Clock  {
	m := (hours * 60 + minutes) % dayMinutes
	if m < 0 {
		m += dayMinutes
	}
	return Clock(m)
}

func (c Clock) Add(m int) Clock {
	return New(0, int(c) + m)
}

func (c Clock) Subtract(m int) Clock {
	return New(0, int(c) - m)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c / 60, c % 60)
}

