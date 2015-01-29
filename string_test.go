package arbitrary


import "testing"

import "math/rand"
import "strings"
import "time"


func TestString(t *testing.T) {

	src := rand.NewSource( time.Now().UTC().UnixNano() )

	arb := New(src)

	seen := make(map[string]int)

	const limit = 50
	for i:=0; i<limit; i++ {
		s := arb.String()

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

func TestStringWithStringLength(t *testing.T) {

	src := rand.NewSource( time.Now().UTC().UnixNano() )

	string_length := 10 + rand.Intn(10)

	arb := New(src)

	seen := make(map[string]int)

	const limit = 50
	for i:=0; i<limit; i++ {
		s := arb.String(string_length)

		if s_length := len(s) ; s_length != string_length {
			t.Errorf("Expected the length of the arbitrary string to be %d but it was actually %d. Arbitrary string was %q.", string_length, s_length, s)
		}

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

func TestStringWithStringLengthAndStringToDrawFrom(t *testing.T) {

	src := rand.NewSource( time.Now().UTC().UnixNano() )

	string_length := 10 + rand.Intn(10)
	draw_from := `!@#$%^&*()_+-={}[]|\:;"'<>,.?/`

	arb := New(src)

	seen := make(map[string]int)

	const limit = 50
	for i:=0; i<limit; i++ {
		s := arb.String(string_length, draw_from)

		if s_length := len(s) ; s_length != string_length {
			t.Errorf("Expected the length of the arbitrary string to be %d but it was actually %d. Arbitrary string was %q.", string_length, s_length, s)
		}

		for ii,r := range []rune(s) {
			if !strings.ContainsRune(draw_from, r) {
				t.Errorf("Rune #%d from arbitrary string contained rune %q but was only suppose to contain runes from %q. Arbitrary string was %q.", ii, string(r), draw_from, s)
			}
		}

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

func TestStringWithStringLengthAndRunesToDrawFrom(t *testing.T) {

	src := rand.NewSource( time.Now().UTC().UnixNano() )

	string_length := 10 + rand.Intn(10)
	draw_from := []rune(`!@#$%^&*()_+-={}[]|\:;"'<>,.?/`)

	arb := New(src)

	seen := make(map[string]int)

	const limit = 50
	for i:=0; i<limit; i++ {
		s := arb.String(string_length, draw_from)

		if s_length := len(s) ; s_length != string_length {
			t.Errorf("Expected the length of the arbitrary string to be %d but it was actually %d. Arbitrary string was %q.", string_length, s_length, s)
		}

		for ii,r := range []rune(s) {
			if !strings.ContainsRune(string(draw_from), r) {
				t.Errorf("Rune #%d from arbitrary string contained rune %q but was only suppose to contain runes from %q. Arbitrary string was %q.", ii, string(r), string(draw_from), s)
			}
		}

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
