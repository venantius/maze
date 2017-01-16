package model

import (
	"math/rand"
	"time"
	"math"
)

var RANDOM *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()));

func round(f float64) int {
	if math.Abs(f) < 0.5 {
		return 0
	}
	return int(f + math.Copysign(0.5, f))
}
