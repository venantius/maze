package util

import (
	"math/rand"
	"time"
)

var RANDOM *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()));


