package sitemap

import (
	"encoding/xml"
	"log"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"time"
)

// Index はサイトマップの index の構造定義です.
type Index struct {
	urlsets   map[string]URLSet
	Basedir   string
	Filename  string
	URLPrefix *url.URL
}

type sitemapindex struct {
	XMLName  xml.Name  `xml:"sitemapindex"`
	XMLNS    string    `xml:"xmlns,attr"`
	Sitemaps []sitemap `xml:"sitemap"`
}

type sitemap struct {
	XMLName xml.Name `xml:"sitemap"`
	Loc     string   `xml:"loc"`
	LastMod string   `xml:"lastmod"`
}

// NewIndex は Index を初期化してそのポインタを返します.
func NewIndex() *Index {
	m := map[string]URLSet{}
	idx := &Index{urlsets: m, Filename: "sitemap_index"}
	return idx
}

// Configure は Index のプロパティを書き換えるために使われます.
func (idx *Index) Configure(options ...func(*Index)) *Index {
	for _, option := range options {
		option(idx)
	}
	return idx
}

// Add は Index.urlsets に URLSet を追加します.
func (idx *Index) Add(filename string, us URLSet) *Index {
	idx.urlsets[filename] = us
	return idx
}

// Output はサイトマップの index ファイルを生成します.
func (idx *Index) Output() {
	var file *os.File
	var err error
	name := addExt(filepath.Join(idx.Basedir, idx.Filename))
	file, err = os.OpenFile(name, os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	b := []byte(xml.Header)
	_, err = file.Write(b)
	if err != nil {
		log.Fatal(err)
	}
	smi := &sitemapindex{XMLNS: xmlNS}
	for p, us := range idx.urlsets {
		for i := 0; i <= len(us.URLs)/us.limit; i++ {
			name := addExt(addNum(p, i))
			loc := idx.URLPrefix.Scheme + "://" + idx.URLPrefix.Hostname() + path.Join("/", idx.URLPrefix.Path, name)
			lastMod := time.Now().Format("2006-01-02")
			sm := sitemap{Loc: loc, LastMod: lastMod}
			smi.Sitemaps = append(smi.Sitemaps, sm)
		}
	}
	b, err = xml.MarshalIndent(smi, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Write(b)
	if err != nil {
		log.Fatal(err)
	}
}

// Generate はサイトマップの index と各ファイルを生成します.
func (idx *Index) Generate() {
	idx.Output()
	for p, us := range idx.urlsets {
		us.Output(filepath.Join(idx.Basedir, p))
	}
}
