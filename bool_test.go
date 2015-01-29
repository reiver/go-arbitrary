package arbitrary


import "testing"

import "math/rand"
import "time"


func TestBool(t *testing.T) {

	src := rand.NewSource( time.Now().UTC().UnixNano() )

	arb := New(src)

	seen_false := false
	seen_true := false

	const limit = 50
	for i:=0; i<limit; i++ {
		bl := arb.Bool()

		switch bl {
			case false:
				seen_false = true
			case true:
				seen_true = true
		}
	}

	if !seen_false {
		t.Errorf("It is possible that there is nothing wrong with this test, but it just ran %d times and did not see the value false.", limit)
	}
	if !seen_true {
		t.Errorf("It is possible that there is nothing wrong with this test, but it just ran %d times and did not see the value true.", limit)
	}
}
