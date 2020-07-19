// Copyright 2020 David Norminton. All rights reserved.
// Use of this source code is governed by a MIT License
// license that can be found in the LICENSE file.

// Package episodate uses the api provided by https://www.episodate.com to
// retrieve TV Show data. The user can view show data, add shows to a list,
// and view the created list
package episodate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// ApiShowUrl is the link to the show-details page of episodate minus the name of the show
const ApiShowUrl = "https://www.episodate.com/api/show-details?q="

// ShowDetails starts the routine to extract the data of the TV Show
func ShowDetails(show string) error {
	data, err := GetShowData(GetApiShowUrl(show))
	if err != nil {
		return fmt.Errorf("Error: There was a problem getting show data! %v", err)
	}
	outputShowData(data)
	return nil
}

// GetApiShowUrl generates the full url of the chosen show
func GetApiShowUrl(show string) string {
	return ApiShowUrl + strings.ReplaceAll(show, " ", "-")
}

// Show is the structure of the api at the show-details page
type Show struct {
	TvShow ShowData `json:"TvShow"`
}

// ShowData is the structure of the show-details, show data
type ShowData struct {
	Id          int            `json:"id"`
	Name        string         `json:"name"`
	Permalink   string         `json:"paermalink"`
	Url         string         `json:"url"`
	Description string         `json:"description"`
	StartDate   string         `json:"start_date"`
	Country     string         `json:"country"`
	Status      string         `json:"status"`
	Network     string         `json:"network"`
	Rating      string         `json:"rating"`
	Episodes    []EpisodesList `json:"episodes"`
}

// EpisodesList is the structure of the data needed for the calendar
type EpisodesList struct {
	Season  int    `json:"season"`
	Episode int    `json:"episode"`
	Name    string `json:"name"`
	AirDate string `json:"air_date"`
}

// GetShowData gets the TV show json data from api
func GetShowData(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte("Error"), err
	}

	defer resp.Body.Close()

	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("Error"), err
	}
	return html, nil

}

// output to terminal api response of TV Show
func outputShowData(html []byte) {
	var data Show
	json.Unmarshal([]byte(html), &data)
	fmt.Printf("Name: %s\n", data.TvShow.Name)
	fmt.Printf("Permalink: %s\n", data.TvShow.Permalink)
	fmt.Printf("Url: %s\n", data.TvShow.Url)
	fmt.Printf("Description: %s\n", data.TvShow.Description)
	fmt.Printf("StartDate: %s\n", data.TvShow.StartDate)
	fmt.Printf("Country: %s\n", data.TvShow.Country)
	fmt.Printf("Status: %s\n", data.TvShow.Status)
	fmt.Printf("Network: %s\n", data.TvShow.Network)
	fmt.Printf("Rating: %s\n", data.TvShow.Rating)
	fmt.Println("------------------------------------------------")

	for i := 0; i < len(data.TvShow.Episodes); i++ {
		fmt.Printf("S%d E%d\n", data.TvShow.Episodes[i].Season, data.TvShow.Episodes[i].Episode)
		fmt.Printf("%s \n", data.TvShow.Episodes[i].Name)
		fmt.Printf("Air Date %s\n", data.TvShow.Episodes[i].AirDate)
		fmt.Println("------------------------------------------------")
	}
}
