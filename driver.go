package sitemap

import (
	"encoding/xml"
	"os"
	"path/filepath"
)

type driver interface {
	writeXML(string, interface{}, bool) error
}

// FileDriver はXMLファイルを書き出すためのドライバです.
type FileDriver struct {
	BasePath string
	Ugly     bool
}

func (d *FileDriver) writeXML(name string, xs interface{}, ugly bool) (err error) {
	var f *os.File
	f, err = newXMLFile(filepath.Join(d.BasePath, name))
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
