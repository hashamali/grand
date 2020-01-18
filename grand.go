package grand 

import (
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

// RandomStringWithSeed generates a random alpha-numeric string of length n, seeded by seed value.
func RandomStringWithSeed(n int, seed int64) string {
	rand.Seed(seed)

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

// RandomString generates a random alpha-numeric string of length n.
func RandomString(n int) string {
	return RandomStringWithSeed(n, TimeSeed())
}

// TimeSeed uses the current unix timestamp.
func TimeSeed() int64 {
	return time.Now().UnixNano()
}
