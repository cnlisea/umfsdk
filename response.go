package umfsdk

import (
	"errors"
	"golang.org/x/net/html"
	"io"
	"strings"
	"unsafe"
)

func ResponseParse(data io.Reader) ([]byte, error) {
	doc, err := html.Parse(data)
	if err != nil {
		return nil, err
	}

	var search func(*html.Node) *html.Node
	search = func(n *html.Node) *html.Node {
		if n.Type == html.ElementNode && n.Data == "meta" {
			return n
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if r := search(c); r != nil {
				return r
			}
		}
		return nil
	}

	if doc = search(doc); doc == nil {
		return nil, errors.New("Invalid response data")
	}

	var ret []byte
	for i := range doc.Attr {
		if doc.Attr[i].Key == "CONTENT" || strings.Index(doc.Attr[i].Val, "sign") >= 0 {
			ret = *(*[]byte)(unsafe.Pointer(&doc.Attr[i].Val))
			break
		}
	}

	return ret, nil
}
