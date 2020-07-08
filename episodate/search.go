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
func SearchTvShow(showName string) {
	show := strings.ReplaceAll(showName, " ", "+")
	showInfoUrl := SearchUrl + show

	getSearchData(showInfoUrl, 1, 0)
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

// getSearchData is used a recursive function in order to retrieve every search result page
func getSearchData(url string, page int, totalPages int) {

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	// show the HTML code as a string %s
	search := SearchJson{}
	json.Unmarshal([]byte(html), &search)

	if totalPages == 0 {
		totalPages = search.Pages
	}

	if page <= totalPages {

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
		getSearchData(pageUrl, page, totalPages)
	}

}
