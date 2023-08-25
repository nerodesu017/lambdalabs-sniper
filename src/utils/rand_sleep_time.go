package utils

import (
	"math/rand"
	"time"
)

var rng *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GetRandSleepTimeInMs(min int, max int) int {
	return rng.Intn(max - min + 1) + min
}