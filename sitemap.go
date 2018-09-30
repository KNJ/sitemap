package sitemap

import (
	"encoding/xml"
	"os"
	"strconv"
)

const xmlNS string = "http://www.sitemaps.org/schemas/sitemap/0.9"

func writeXML(p string, xs interface{}, ugly bool) (err error) {
	var f *os.File
	f, err = os.OpenFile(p+".xml", os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0600)
	defer f.Close()
	if err != nil {
		return
	}
	b := []byte(xml.Header)
	_, err = f.Write(b)
	if err != nil {
		return
	}
	if ugly {
		b, err = xml.Marshal(xs)
	} else {
		b, err = xml.MarshalIndent(xs, "", "    ")
	}
	if err != nil {
		return
	}
	_, err = f.Write(b)
	return
}

func addNum(base string, i int) string {
	return base + "_" + strconv.Itoa(i)
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}
