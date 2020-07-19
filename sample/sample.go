package sample

import "time"

type Sample struct {
	MessagePerSecond float64
	BytesPerSecond   float64
	Duration         time.Duration
	TripTime         time.Duration
}
