package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// PornhubSuggestion contains all the fields that make up a Pornhub suggestion.
type PornhubSuggestion struct {
	Num0            string   `json:"0"`
	Num1            string   `json:"1"`
	Num2            string   `json:"2"`
	Num3            string   `json:"3"`
	Num4            string   `json:"4"`
	Num5            string   `json:"5"`
	Num6            string   `json:"6"`
	PopularSearches []string `json:"popularSearches"`
}

func getPornhubResults(typed string) (PornhubSuggestion, error) {
	u := "https://www.pornhub.com/video/search_autocomplete?orientation=straight&q=" + url.QueryEscape(typed) + "&alt=0"
	phClient := http.Client{
		Timeout: time.Second * 10,
	}
	request, err := http.NewRequest(http.MethodGet, u, nil)
	request.Header.Set("User-Agent", "Making Steve Jobs turn in his grave")
	if err != nil {
		return PornhubSuggestion{}, err
	}
	response, err := phClient.Do(request)
	if err != nil {
		return PornhubSuggestion{}, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return PornhubSuggestion{}, err
	}
	var result PornhubSuggestion
	err = json.Unmarshal(body, &result)
	return result, nil
}
