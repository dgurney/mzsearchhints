package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func (t trends) catch(err error) string {
	var e []Trend
	for i := 0; i < 6; i++ {
		e = append(e, Trend{"https://search.itunes.apple.com/WebObjects/MZStore.woa/wa/search?src=trending&term=Halsey", fmt.Sprintf("%s", err)})
	}
	eb := new(bytes.Buffer)
	enc := json.NewEncoder(eb)
	enc.SetEscapeHTML(false)
	err = enc.Encode(e)
	return eb.String()
}

func (h hints) catch(err error) string {
	var e []Hint
	for i := 0; i < 6; i++ {
		e = append(e, Hint{fmt.Sprintf("%s", err), fmt.Sprintf("%s", err), "1"})
	}
	eb := new(bytes.Buffer)
	enc := json.NewEncoder(eb)
	enc.SetEscapeHTML(false)
	err = enc.Encode(e)
	return eb.String()
}
