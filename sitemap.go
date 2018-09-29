package sitemap

import (
	"encoding/xml"
	"log"
	"os"
	"strconv"
)

const xmlDec string = "<?xml version=\"1.0\" encoding=\"UTF-8\" ?>\n"
const xmlNS string = "http://www.sitemaps.org/schemas/sitemap/0.9"

// URLSet は <urlset> の構造定義です.
type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	XMLNS   string   `xml:"xmlns,attr"`
	URLs    []URL    `xml:"url"`
	limit   int
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
	us.limit = 3
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
	for i := 0; i <= len(urls)/us.limit; i++ {
		us.outputSingleFile(p, i, urls[i*us.limit:min((i+1)*us.limit, len(urls))])
	}
}

func (us *URLSet) outputSingleFile(p string, i int, urls []URL) {
	var f *os.File
	var err error
	name := addExt(addNum(p, i))
	f, err = os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	b := []byte(xmlDec)
	_, err = f.Write(b)
	if err != nil {
		log.Fatal(err)
	}
	us.URLs = urls
	b, err = xml.MarshalIndent(us, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write(b)
	if err != nil {
		log.Fatal(err)
	}
}

func addNum(base string, i int) string {
	return base + "_" + strconv.Itoa(i)
}

func addExt(base string) string {
	return base + ".xml"
}

func min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}
