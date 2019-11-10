package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/text/encoding/htmlindex"
	"golang.org/x/text/transform"
)

type (
	Root struct {
		Items []string `xml:"-"`
	}
)

func main() {
	var (
		r       Root
		rd, err = os.Open("./test.xml")
	)

	if err != nil {
		panic(err)
	}

	xr := xml.NewDecoder(rd)
	buf := new(bytes.Buffer)
	var start, end, offset int64

	xr.CharsetReader = func(label string, input io.Reader) (io.Reader, error) {
		charset, err := htmlindex.Get(label)
		if err != nil {
			return nil, err
		}

		offset = xr.InputOffset()
		decoder := transform.NewReader(input, charset.NewDecoder())

		return io.TeeReader(decoder, buf), nil
	}

loop:
	for {
		tok, err := xr.RawToken()
		if err != nil {
			if err == io.EOF {
				break loop
			}

			panic(err)
		}

		switch t := tok.(type) {
		case xml.StartElement:
			switch t.Name.Local {
			case "r":
				start = xr.InputOffset()
			}

		case xml.EndElement:
			switch t.Name.Local {
			case "a":
				end = xr.InputOffset()
				r.Items = append(r.Items, strings.TrimSpace(buf.String()[start-offset:end-offset]))
				start = end
			case "r":
				break loop
			}
		}
	}

	for i := range r.Items {
		fmt.Printf("%d: %s\n", i+1, r.Items[i])
	}
}
