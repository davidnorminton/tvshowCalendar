package web

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"tvshowCalendar/episodate"
	"tvshowCalendar/showlist"
)

type Show struct {
	TvShow ShowDetails `json:"TvShow"`
}

type ShowDetails struct {
	Name        string `json:"name"`
	Permalink   string `json:"permalink"`
	Description string `json:"description"`
	StartDate   string `json:"start_date"`
	Country     string `json:"country"`
	Status      string `json:"status"`
	Network     string `json:"network"`
	Rating      int    `json:"rating"`
	YouTube     string `json:"youtube_link"`
	Image       string `json:"image_path"`
}

type DetailsView struct {
	DetailsData ShowDetails
	Query       string
	IsAdded     bool
}

func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	query := r.URL.Query().Get("q")

	t, err := template.ParseFiles("web/html/details.html", "web/html/header.html", "web/html/footer.html")
	if err != nil {
		panic(err)
	}
	url := episodate.GetApiShowUrl(query)
	htmlData, err := episodate.GetShowData(url)

	if err != nil {
		fmt.Println("show not found")
	}

	var data Show
	json.Unmarshal([]byte(htmlData), &data)
	isAdded := false
	if err := showlist.CheckIfShowInList(data.TvShow.Permalink); err != nil {
		isAdded = true
	}
	showData := ShowDetails{
		data.TvShow.Name,
		data.TvShow.Permalink,
		data.TvShow.Description,
		data.TvShow.StartDate,
		data.TvShow.Country,
		data.TvShow.Status,
		data.TvShow.Network,
		data.TvShow.Rating,
		data.TvShow.YouTube,
		data.TvShow.Image,
	}

	detailsData := DetailsView{
		DetailsData: showData,
		Query:       query,
		IsAdded:     isAdded,
	}
	err = t.Execute(w, detailsData)
	if err != nil {
		panic(err)
	}
}
