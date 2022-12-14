package arbitrary

import (
	"io/fs"
)

// RegularFile returns an arbitrary fs.File regular-file.
func (arb T) RegularFile() fs.File {

	var fns [](func()fs.File) = [](func()fs.File){
		arb.TextFile,
		arb.HTMLFile,
	}

	fn := fns[arb.randomness.Intn(len(fns))]

	return fn()

}
