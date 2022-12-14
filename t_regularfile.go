package arbitrary

import (
	"io/fs"
)

func (arb T) RegularFile() fs.File {

	var fns [](func()fs.File) = [](func()fs.File){
		arb.TextFile,
		arb.HTMLFile,
	}

	fn := fns[arb.randomness.Intn(len(fns))]

	return fn()

}
