package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Trend contains all the fields that make up a trending search.
type Trend struct {
	URL   string `json:"url"` // iTunes cares about where it points, so unfortunately you can't open YouTube or whatever by changing this.
	Label string `json:"label"`
}
type trends []Trend

func trendsHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Trends request from ", r.RemoteAddr, " (user agent: ", r.UserAgent(), ")")
	w.Header().Set("Content-Type", "application/json")
	ph, err := getPornhubResults("ass")
	var pronTrends trends
	if err != nil {
		fmt.Fprintf(w, pronTrends.catch(err))
		return
	}
	for i := 0; i < len(ph.PopularSearches); i++ {
		trend := Trend{
			Label: ph.PopularSearches[i],
			URL:   "https://search.itunes.apple.com/WebObjects/MZStore.woa/wa/search?src=trending&term=Ariana%20Grande",
		}
		pronTrends = append(pronTrends, trend)
	}

	// Emulate the real service by not escaping ampersands
	trends := new(bytes.Buffer)
	enc := json.NewEncoder(trends)
	enc.SetEscapeHTML(false)
	err = enc.Encode(pronTrends)
	if err != nil {
		fmt.Fprint(w, pronTrends.catch(err))
		return
	}
	fmt.Fprint(w, trends.String())
}
