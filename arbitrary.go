package arbitrary


import "math/rand"


// A T is a source of arbitrary data.
type T struct {
	randomness *rand.Rand
}


// New returns a new T.
func New(src rand.Source) T {

	randomness := rand.New(src)

	return T{randomness: randomness}
}
