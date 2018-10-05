# sitemap

## Example

*main_go:*

```go
package main

import (
	"github.com/KNJ/sitemap"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	idx := sitemap.NewIndex().Configure(urlPrefix("https://wazly.com"))
	us1 := sitemap.NewURLSet().Configure(prefix("https://wazly.com/works/"))
	for i := 0; i < 50; i++ {
		loc := strconv.Itoa(i)
		us1.AddURL(sitemap.URL{loc, 0.8, "monthly"})
	}
	us2 := sitemap.NewURLSet().Configure(prefix("https://wazly.com/todo?page="), limit(200))
	for i := 0; i < 1000; i++ {
		loc := strconv.Itoa(i)
		us2.AddURL(sitemap.URL{loc, 0.5, "daily"})
	}
	idx.Add("works", *us1)
	idx.Add("todo", *us2)
	idx.Generate(&sitemap.FileDriver{BasePath: filepath.Join(wd, "xml")})
}

func urlPrefix(s string) func(*sitemap.Index) {
	u, err := url.Parse(s)
	if err != nil {
		log.Fatal(err)
	}
	return func(idx *sitemap.Index) {
		idx.URLPrefix = u
	}
}

func prefix(s string) func(*sitemap.URLSet) {
	return func(us *sitemap.URLSet) {
		us.Prefix = s
	}
}

func limit(i int) func(*sitemap.URLSet) {
	return func(us *sitemap.URLSet) {
		us.Limit = i
	}
}
```

## Output

*./xml/sitemap_index.xml:*

```xml
<?xml version="1.0" encoding="UTF-8"?>
<sitemapindex xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
    <sitemap>
        <loc>https://wazly.com/works_0.xml</loc>
        <lastmod>2018-10-06</lastmod>
    </sitemap>
    <sitemap>
        <loc>https://wazly.com/todo_0.xml</loc>
        <lastmod>2018-10-06</lastmod>
    </sitemap>
    <sitemap>
        <loc>https://wazly.com/todo_1.xml</loc>
        <lastmod>2018-10-06</lastmod>
    </sitemap>
    <sitemap>
        <loc>https://wazly.com/todo_2.xml</loc>
        <lastmod>2018-10-06</lastmod>
    </sitemap>
    <sitemap>
        <loc>https://wazly.com/todo_3.xml</loc>
        <lastmod>2018-10-06</lastmod>
    </sitemap>
    <sitemap>
        <loc>https://wazly.com/todo_4.xml</loc>
        <lastmod>2018-10-06</lastmod>
    </sitemap>
</sitemapindex>
```

*./xml/todo_0.xml:*

```xml
<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
    <url>
        <loc>https://wazly.com/todo?page=0</loc>
        <priority>0.5</priority>
        <changefreq>daily</changefreq>
    </url>
    <url>
        <loc>https://wazly.com/todo?page=1</loc>
        <priority>0.5</priority>
        <changefreq>daily</changefreq>
    </url>
    
...

    <url>
        <loc>https://wazly.com/todo?page=198</loc>
        <priority>0.5</priority>
        <changefreq>daily</changefreq>
    </url>
    <url>
        <loc>https://wazly.com/todo?page=199</loc>
        <priority>0.5</priority>
        <changefreq>daily</changefreq>
    </url>
</urlset>    
```