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

// HTMLFile returns an arbitrary fs.File regular-file whose content is an HTML file.
func HTMLFile() fs.File {
	return Default.HTMLFile()
}

// NetAddr returns an arbitrary net.Addr.
func NetAddr() net.Addr {
	return Default.NetAddr()
}

// NetPort returns an arbitrary network port.
func NetPort() uint16 {
	return Default.NetPort()
}

// Password returns an arbitrary password.
func Password() string {
	return Default.Password()
}

// Password returns an arbitrary pct-encoded.
func PctEncoded() string {
	return Default.PctEncoded()
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

// TextFile returns an arbitrary fs.File regular-file whose content is a text-file.
func TextFile() fs.File {
	return Default.TextFile()
}
