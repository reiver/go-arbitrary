package arbitrary

import (
	"math/rand"
	"time"
)

var (
	Default = New( rand.NewSource( time.Now().UTC().UnixNano() ) )
)

// Bool returns an arbitrary bool.
func Bool() bool {
	return Default.Bool()
}

// PhoneNumber returns an arbitrary phone‐number.
func PhoneNumber() string {
	return Default.PhoneNumber()
}

func Runes(a ...interface{}) []rune {
	return Default.Runes(a...)
}

func String(a ...interface{}) string {
	return Default.String(a...)
}
