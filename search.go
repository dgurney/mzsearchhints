// Quick and dirty iTunes search suggestion engine
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

type Hint struct {
	Term         string `json:"term"`
	CensoredTerm string `json:"censoredTerm"`
	Score        string `json:"score"`
}

type Trend struct {
	URL   string `json:"url"`
	Label string `json:"label"`
}

type Trends []Trend
type Hints []Hint

func hintsHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Suggestion request from ", r.RemoteAddr, " (user agent: ", r.UserAgent(), ")")
	w.Header().Set("Content-Type", "application/json")
	randomTerms := []string{"This is pointless", "Most Palone", "Saylor Twift", "Kim Jong-un", "Lukas Marsik", "Daniel Gurney", "💩"}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(randomTerms), func(i, j int) { randomTerms[i], randomTerms[j] = randomTerms[j], randomTerms[i] })
	var testHints Hints
	for i := 0; i < len(randomTerms); i++ {
		hint := Hint{
			Term:         randomTerms[i],
			CensoredTerm: randomTerms[i],
			Score:        strconv.Itoa(rand.Intn(10000)),
		}
		testHints = append(testHints, hint)
	}
	hints, err := json.Marshal(testHints)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%s", hints)
}

func trendsHandler(w http.ResponseWriter, r *http.Request) {
	log.Print("Trends request from ", r.RemoteAddr, " (user agent: ", r.UserAgent(), ")")
	w.Header().Set("Content-Type", "application/json")
	randomLabels := []string{"🍕🥓", "Vladimir Putin", "Your mom", "PewDiePie", "👉👌"}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(randomLabels), func(i, j int) { randomLabels[i], randomLabels[j] = randomLabels[j], randomLabels[i] })
	var testTrends Trends
	for i := 0; i < len(randomLabels); i++ {
		trend := Trend{
			Label: randomLabels[i],
			URL:   "https://search.itunes.apple.com/WebObjects/MZStore.woa/wa/search?src=trending&term=Ariana%20Grande",
		}
		testTrends = append(testTrends, trend)
	}
	trends, err := json.Marshal(testTrends)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, "%s", trends)
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
