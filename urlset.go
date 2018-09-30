package sitemap

import (
	"encoding/xml"
	"log"
)

// URLSet は <urlset> の構造定義です.
type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	XMLNS   string   `xml:"xmlns,attr"`
	URLs    []URL    `xml:"url"`
	Limit   int      `xml:"-"`
	Prefix  string   `xml:"-"`
	Ugly    bool     `xml:"-"`
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
	err := writeXML(p, *us, us.Ugly)
	if err != nil {
		log.Fatal(err)
	}
}
