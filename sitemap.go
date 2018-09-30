package sitemap

import (
	"encoding/xml"
	"log"
	"os"
	"strconv"
)

const xmlNS string = "http://www.sitemaps.org/schemas/sitemap/0.9"

// URLSet は <urlset> の構造定義です.
type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	XMLNS   string   `xml:"xmlns,attr"`
	URLs    []URL    `xml:"url"`
	Limit   int      `xml:"-"`
	Prefix  string   `xml:"-"`
}

// URL は <url> の構造定義です.
type URL struct {
	Loc        string  `xml:"loc"`
	Priority   float64 `xml:"priority"`
	ChangeFreq string  `xml:"changefreq"`
}

// NewURLSet は URLSet を初期化してそのポインタを返します.
func NewURLSet() *URLSet {
	us := &URLSet{}
	us.XMLNS = xmlNS
	us.Limit = 50000
	return us
}

// Configure は URLSet のプロパティを書き換えるために使われます.
func (us *URLSet) Configure(options ...func(*URLSet)) *URLSet {
	for _, option := range options {
		option(us)
	}
	return us
}

// AddURL は URLSet に URL を追加します.
func (us *URLSet) AddURL(url URL) *URLSet {
	us.URLs = append(us.URLs, url)
	return us
}

// Output はファイルを書き出します.
func (us *URLSet) Output(p string) {
	urls := us.URLs
	for i := 0; i <= len(urls)/us.Limit; i++ {
		us.outputSingleFile(p, i, urls[i*us.Limit:min((i+1)*us.Limit, len(urls))])
	}
}

func (us *URLSet) outputSingleFile(p string, i int, urls []URL) {
	p = addNum(p, i)
	if us.Prefix != "" {
		for i, u := range urls {
			urls[i].Loc = us.Prefix + u.Loc
		}
	}
	us.URLs = urls
	err := writeXML(p, *us)
	if err != nil {
		log.Fatal(err)
	}
}

func writeXML(p string, xs interface{}) (err error) {
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
	b, err = xml.MarshalIndent(xs, "", "    ")
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
