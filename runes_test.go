package arbitrary


import "testing"

import "math/rand"
import "strings"
import "time"


func TestRunes(t *testing.T) {

	src := rand.NewSource( time.Now().UTC().UnixNano() )

	arb := New(src)

	seen := make(map[string]int)

	const limit = 50
	for i:=0; i<limit; i++ {
		rns := arb.Runes()

		s := string(rns)
		if _,ok := seen[s] ; !ok {
			seen[s] = 0
		} else {
			seen[s]++
		}
	}

	if length := len(seen) ; length != limit {
		t.Errorf("It is possible that there is nothing wrong with this test, but it just ran %d times and saw repeats such that only saw %d uniques.", limit, length)
	}
}

func TestRunesWithStringLength(t *testing.T) {

	src := rand.NewSource( time.Now().UTC().UnixNano() )

	string_length := 10 + rand.Intn(10)

	arb := New(src)

	seen := make(map[string]int)

	const limit = 50
	for i:=0; i<limit; i++ {
		rns := arb.Runes(string_length)

		if rns_length := len(rns) ; rns_length != string_length {
			t.Errorf("Expected the length of the arbitrary string to be %d but it was actually %d. Arbitrary string was %q.", string_length, rns_length, rns)
		}

		s := string(rns)
		if _,ok := seen[s] ; !ok {
			seen[s] = 0
		} else {
			seen[s]++
		}
	}

	if length := len(seen) ; length != limit {
		t.Errorf("It is possible that there is nothing wrong with this test, but it just ran %d times and saw repeats such that only saw %d uniques.", limit, length)
	}
}

func TestRunesWithStringLengthAndStringToDrawFrom(t *testing.T) {

	src := rand.NewSource( time.Now().UTC().UnixNano() )

	string_length := 10 + rand.Intn(10)
	draw_from := `!@#$%^&*()_+-={}[]|\:;"'<>,.?/`

	arb := New(src)

	seen := make(map[string]int)

	const limit = 50
	for i:=0; i<limit; i++ {
		rns := arb.Runes(string_length, draw_from)

		if rns_length := len(rns) ; rns_length != string_length {
			t.Errorf("Expected the length of the arbitrary string to be %d but it was actually %d. Arbitrary string was %q.", string_length, rns_length, rns)
		}

		for ii,r := range rns {
			if !strings.ContainsRune(draw_from, r) {
				t.Errorf("Rune #%d from arbitrary string contained rune %q but was only suppose to contain runes from %q. Arbitrary string was %q.", ii, string(r), draw_from, string(rns))
			}
		}

		s := string(rns)
		if _,ok := seen[s] ; !ok {
			seen[s] = 0
		} else {
			seen[s]++
		}
	}

	if length := len(seen) ; length != limit {
		t.Errorf("It is possible that there is nothing wrong with this test, but it just ran %d times and saw repeats such that only saw %d uniques.", limit, length)
	}
}

func TestRunesWithStringLengthAndRunesToDrawFrom(t *testing.T) {

	src := rand.NewSource( time.Now().UTC().UnixNano() )

	string_length := 10 + rand.Intn(10)
	draw_from := []rune(`!@#$%^&*()_+-={}[]|\:;"'<>,.?/`)

	arb := New(src)

	seen := make(map[string]int)

	const limit = 50
	for i:=0; i<limit; i++ {
		rns := arb.Runes(string_length, draw_from)

		if rns_length := len(rns) ; rns_length != string_length {
			t.Errorf("Expected the length of the arbitrary string to be %d but it was actually %d. Arbitrary string was %q.", string_length, rns_length, rns)
		}

		for ii,r := range rns {
			if !strings.ContainsRune(string(draw_from), r) {
				t.Errorf("Rune #%d from arbitrary string contained rune %q but was only suppose to contain runes from %q. Arbitrary string was %q.", ii, string(r), string(draw_from), string(rns))
			}
		}

		s := string(rns)
		if _,ok := seen[s] ; !ok {
			seen[s] = 0
		} else {
			seen[s]++
		}
	}

	if length := len(seen) ; length != limit {
		t.Errorf("It is possible that there is nothing wrong with this test, but it just ran %d times and saw repeats such that only saw %d uniques.", limit, length)
	}
}
