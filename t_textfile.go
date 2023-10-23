package arbitrary

import (
	"sourcecode.social/reiver/go-strfs"

	"io/fs"
	"time"
)

var (
	textFiles []string = []string{
		"",

		"Hello world!",
		"Hello world!" +"\n",
		"Hello world!" +"\r",
		"Hello world!" +"\r\n",
		"Hello world!" +"\u0085",
		"Hello world!" +"\u2028",
		"Hello world!" +"\u2029",

		"To Whom It May Concern:" +"\n"+"\n"+         "We were not alone." +"\n",
		"To Whom It May Concern:" +"\r"+"\r"+         "We were not alone." +"\r",
		"To Whom It May Concern:" +"\r\n"+"\r\n"+     "We were not alone." +"\r\n",
		"To Whom It May Concern:" +"\u0085"+"\u0085"+ "We were not alone." +"\u0085",
		"To Whom It May Concern:" +"\u2029"+          "We were not alone." +"\u2028",
	}
)

// TextFile returns an arbitrary fs.File regular-file whose content is a text-file.
func (arb T) TextFile() fs.File {

	var filecontent strfs.Content
	{
		var s string = textFiles[arb.randomness.Intn(len(textFiles))]

		filecontent = strfs.CreateContent(s)
	}

	var filename string
	{
		switch arb.randomness.Intn(10) {
		case 0:
			filename = ".plan"
		case 1:
			filename = ".project"
		case 2:
			filename = "message.txt"
		case 3:
			filename = "style.txt"
		case 4:
			filename = "style.txt"
		case 5:
			filename = "DOCUMENT.txt"
		case 6:
			filename = "file (2).txt"
		default:
			filename = arb.String()
			if arb.Bool() {
				filename += ".txt"
			}
		}
	}

	var filemodtime time.Time = time.Now()
	func(){
		if 0 == arb.randomness.Intn(10) {
			filemodtime = time.Date(1970,1,1,0,0,0,0,time.UTC)
			return
		}

		if 0 == arb.randomness.Intn(10) {
			filemodtime = time.Date(1924,6,7,20,6,3,11,time.UTC)
			return
		}

		if 0 == arb.randomness.Intn(10) {
			filemodtime = time.Date(2022,12,14,20,12,23,17,time.UTC)
			return
		}

		if 0 == 3 {
			year       := arb.randomness.Intn(3000)
			month      := time.Month(1+arb.randomness.Intn(12))
			day        := 1+arb.randomness.Intn(31)
			hour       := arb.randomness.Intn(24)
			minute     := arb.randomness.Intn(60)
			second     := arb.randomness.Intn(60)
			nanosecond := arb.randomness.Intn(60)

			filemodtime = time.Date(year, month, day, hour, minute, second, nanosecond, time.UTC)
			return
		}
	}()

	var regularfile strfs.RegularFile = strfs.RegularFile{
		FileContent: filecontent,
		FileName:    filename,
		FileModTime: filemodtime,
	}

	var file fs.File = &regularfile

	return file
}
