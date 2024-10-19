package utilities

import (
	"time"
)

type Timer struct {
	Start time.Time
}

func (timer *Timer) TrackTime() time.Duration{
	return time.Since(timer.Start)
}