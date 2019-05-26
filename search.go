// Quick and dirty (literally) iTunes search suggestion engine
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// Hint contains all the fields that make up a hint/suggestion.
type Hint struct {
	Term         string `json:"term"`
	CensoredTerm string `json:"censoredTerm"`
	Score        string `json:"score"`
}

type hints []Hint

func hintsHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Suggestion request from ", r.RemoteAddr, " (user agent: ", r.UserAgent(), ")")
	w.Header().Set("Content-Type", "application/json")
	rand.Seed(time.Now().UnixNano())
	var hints hints
	term := ""
	switch {
	default:
		term = r.URL.Query()["term"][0]
	case len(r.URL.Query()["term"][0]) == 0:
		term = "test"
	}
	ph, err := getPornhubResults(term)
	if err != nil {
		fmt.Fprint(w, hints.catch(err))
		return
	}
	hint0 := Hint{
		Term:         ph.Num0,
		CensoredTerm: ph.Num0,
		Score:        strconv.Itoa(rand.Intn(10000)),
	}
	hint1 := Hint{
		Term:         ph.Num1,
		CensoredTerm: ph.Num1,
		Score:        strconv.Itoa(rand.Intn(10000)),
	}
	hint2 := Hint{
		Term:         ph.Num2,
		CensoredTerm: ph.Num2,
		Score:        strconv.Itoa(rand.Intn(10000)),
	}
	hint3 := Hint{
		Term:         ph.Num3,
		CensoredTerm: ph.Num3,
		Score:        strconv.Itoa(rand.Intn(10000)),
	}
	hint4 := Hint{
		Term:         ph.Num4,
		CensoredTerm: ph.Num4,
		Score:        strconv.Itoa(rand.Intn(10000)),
	}
	hint5 := Hint{
		Term:         ph.Num5,
		CensoredTerm: ph.Num5,
		Score:        strconv.Itoa(rand.Intn(10000)),
	}
	hint6 := Hint{
		Term:         ph.Num6,
		CensoredTerm: ph.Num6,
		Score:        strconv.Itoa(rand.Intn(10000)),
	}
	hints = append(hints, hint0)
	hints = append(hints, hint1)
	hints = append(hints, hint2)
	hints = append(hints, hint3)
	hints = append(hints, hint4)
	hints = append(hints, hint5)
	hints = append(hints, hint6)
	phHints, err := json.Marshal(hints)
	if err != nil {
		fmt.Fprint(w, hints.catch(err))
		return
	}
	fmt.Fprint(w, string(phHints))
}

func main() {
	port := flag.Int("p", 80, "Port to run on")
	flag.Parse()
	if *port == 0 {
		*port = 80
	}
	http.HandleFunc("/WebObjects/MZSearchHints.woa/wa/hints/", hintsHandler)
	http.HandleFunc("/WebObjects/MZSearchHints.woa/wa/trends/", trendsHandler)
	s := http.Server{
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
		Addr:         ":" + strconv.Itoa(*port),
	}
	log.Print("Starting MZSearchHints on port ", *port)
	if err := s.ListenAndServe(); err != nil {
		fmt.Println("Unable to start server:", err)
	}
}
