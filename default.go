package arbitrary

import (
	"math/rand"
	"io/fs"
	"net"
	"time"
)

var (
	Default = New( rand.NewSource( time.Now().UTC().UnixNano() ) )
)

// Bool returns an arbitrary bool.
func Bool() bool {
	return Default.Bool()
}

// NetAddr returns an arbitrary net.Addr.
func NetAddr() net.Addr {
	return Default.NetAddr()
}

// Password returns an arbitrary password.
func Password() string {
	return Default.Password()
}

// PhoneNumber returns an arbitrary phone‐number.
//
// Some example phone‐number include:
//
//	289-5625421
//	(604) 585-1234
//	٠٩٦-٣٦١-٦٦٧٦
//	613/978-1920
//	+20-13-004-6453
//	(782) 440 6832
//	+٢٠-٩٣-٧٤٨-٧٩٨١
//	093-340-0709
//	+1-403-870-6449
//	+18250827424
//
// (Note that there are other phone‐number formats too, in addition to the formats that are implied by the examples provided here.)
func PhoneNumber() string {
	return Default.PhoneNumber()
}

// RegularFile returns an arbitrary fs.File regular-file.
func RegularFile() fs.File {
	return Default.RegularFile()
}

func Runes(a ...interface{}) []rune {
	return Default.Runes(a...)
}

func String(a ...interface{}) string {
	return Default.String(a...)
}
