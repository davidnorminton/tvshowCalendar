// Copyright 2020 David Norminton. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

// Package episodate, this file contains functions used to search
// episodate's api for terms related to TV Shows
package episodate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

// SearchUrl is the main url part of the search api
const (
	SearchUrl = "https://www.episodate.com/api/search?q="
)

// SearchTvShow searchs for TV Shows that contain the term
func SearchTvShow(show string) {
	outputConsole(GetSearchData(GetSearchUrl(show), 1))
}

// FormatShowName removes whitespace and replaces with a +
func FormatShowName(show string) string {
	return strings.ReplaceAll(show, " ", "+")
}

// GetSearchUrl creates the search url
func GetSearchUrl(show string) string {
	return SearchUrl + FormatShowName(show)
}

// SearchShowArray is the structure of the json show data
type SearchShowArray struct {
	Id                 int    `json:"id"`
	Name               string `json:"name"`
	Permalink          string `json:"permalink"`
	StartDate          string `json:"start_date"`
	EndDate            string `json:"end_date"`
	Country            string `json:"country"`
	Network            string `json:"network"`
	Status             string `json:"status"`
	ImageThumbnailPath string `json:"image_thumbnail_path"`
}

// SearchJson is the top level of the json tree
type SearchJson struct {
	Total   string            `json:"total"`
	Page    int               `json:"page"`
	Pages   int               `json:"pages"`
	TvShows []SearchShowArray `json:"tv_shows"`
}

// GetSearchData retrieves the json response from the api
func GetSearchData(url string, page int) ([]byte, string, int) {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	return html, url, page

}

func outputConsole(html []byte, url string, page int) {
	// show the HTML code as a string %s
	search := SearchJson{}
	json.Unmarshal([]byte(html), &search)

	if page <= search.Pages {

		fmt.Println("--------------------------")
		fmt.Printf("Page %d of results", page)
		fmt.Println("--------------------------")

		for i := 0; i < len(search.TvShows); i++ {
			fmt.Printf("Show name: %s\n", search.TvShows[i].Name)
			fmt.Printf("Reference: %s\n", search.TvShows[i].Permalink)
			fmt.Printf("Start Date: %s\n", search.TvShows[i].StartDate)
			fmt.Printf("Status: %s\n", search.TvShows[i].Status)
			fmt.Println("--------------------------")
		}

		page = page + 1
		pageUrl := url + "&page=" + strconv.Itoa(page)
		outputConsole(GetSearchData(pageUrl, page))
	}
}
