package sitemap

import (
	"os"
	"strconv"
)

const xmlNS string = "http://www.sitemaps.org/schemas/sitemap/0.9"

func newXMLFile(p string) (f *os.File, err error) {
	f, err = os.OpenFile(p+".xml", os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0600)
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
