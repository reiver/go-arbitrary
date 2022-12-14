package arbitrary

import (
	"github.com/reiver/go-strfs"

	"io/fs"
	"time"
)

var (
	htmlFiles []string = []string{
		"",

		"Hello world!",
		"Hello world!" +"\n",
		"Hello world!" +"\r",
		"Hello world!" +"\r\n",
		"Hello world!" +"\u0085",
		"Hello world!" +"\u2028",
		"Hello world!" +"\u2029",

		"<!DOCTYPE html>",
		"<!DOCTYPE html>" +"\n",
		"<!DOCTYPE html>" +"\r",
		"<!DOCTYPE html>" +"\r\n",
		"<!DOCTYPE html>" +"\u0085",
		"<!DOCTYPE html>" +"\u2028",
		"<!DOCTYPE html>" +"\u2029",

		"<div></div>",
		"<div></div>" +"\n",
		"<div></div>" +"\r",
		"<div></div>" +"\r\n",
		"<div></div>" +"\u0085",
		"<div></div>" +"\u2028",
		"<div></div>" +"\u2029",

		"<!DOCTYPE html>" +"\n"    + "<html></html>",
		"<!DOCTYPE html>" +"\r"    + "<html></html>",
		"<!DOCTYPE html>" +"\r\n"  + "<html></html>",
		"<!DOCTYPE html>" +"\u0085"+ "<html></html>",
		"<!DOCTYPE html>" +"\u2028"+ "<html></html>",
		"<!DOCTYPE html>" +"\u2029"+ "<html></html>",

		"<!DOCTYPE html>" +"\n"    + "<html></html>" +"\n",
		"<!DOCTYPE html>" +"\r"    + "<html></html>" +"\r",
		"<!DOCTYPE html>" +"\r\n"  + "<html></html>" +"\r\n",
		"<!DOCTYPE html>" +"\u0085"+ "<html></html>" +"\u0085",
		"<!DOCTYPE html>" +"\u2028"+ "<html></html>" +"\u2028",
		"<!DOCTYPE html>" +"\u2029"+ "<html></html>" +"\u2029",
	}
)

func (arb T) HTMLFile() fs.File {

	var filecontent strfs.Content
	{
		var s string = textFiles[arb.randomness.Intn(len(textFiles))]

		filecontent = strfs.CreateContent(s)
	}

	var filename string
	{
		switch arb.randomness.Intn(10) {
		case 0:
			filename = "index.html"
		case 1:
			filename = "index.htm"
		case 2:
			filename = "INDEX.HTML"
		case 3:
			filename = "INDEX.HTM"
		case 4:
			filename = "about.html"
		case 5:
			filename = "about.htm"
		case 6:
			filename = "something.HoTMeTaL"
		default:
			filename = arb.String()
			if arb.Bool() {
				filename += ".html"
			} else if arb.Bool() {
				filename += ".htm"
			} else if arb.Bool() {
				filename += ".HTML"
			} else if arb.Bool() {
				filename += ".HTM"
			} else if arb.Bool() {
				filename += ".HoTMeTaL"
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
