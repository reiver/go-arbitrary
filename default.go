package arbitrary

import (
	"math/rand"
	"time"
)

var (
	Default = New( rand.NewSource( time.Now().UTC().UnixNano() ) )
)

func Bool() bool {
	return Default.Bool()
}

func Runes(a ...interface{}) []rune {
	return Default.Runes(a...)
}

func String(a ...interface{}) string {
	return Default.String(a...)
}
