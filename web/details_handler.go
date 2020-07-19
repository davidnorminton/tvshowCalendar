// This component handles the data that is then rendered to the /details page
package web

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/davidnorminton/tvshowCalendar/episodate"
	"github.com/davidnorminton/tvshowCalendar/showlist"
	"github.com/davidnorminton/tvshowCalendar/utils"
)

type Show struct {
	TvShow ShowDetails `json:"TvShow"`
}

// ShowDetails is the structure of the show json response data from the api
type ShowDetails struct {
	Name        string         `json:"name"`
	Permalink   string         `json:"permalink"`
	Description string         `json:"description"`
	StartDate   string         `json:"start_date"`
	Country     string         `json:"country"`
	Status      string         `json:"status"`
	Network     string         `json:"network"`
	Rating      int            `json:"rating"`
	YouTube     string         `json:"youtube_link"`
	Image       string         `json:"image_path"`
	Episodes    []episodesList `json:"episodes"`
}

// DetailsView is the structure of the data used on the details page
type DetailsView struct {
	DetailsData ShowDetails
	Query       string
	IsAdded     bool
	Desc        template.HTML
}

// episodesList is the structure of the data needed for the calendar
type episodesList struct {
	Season  int    `json:"season"`
	Episode int    `json:"episode"`
	Name    string `json:"name"`
	AirDate string `json:"air_date"`
}

// DetailsHandler handles the information displayed on the /details page
func DetailsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	query := r.URL.Query().Get("q")

	t, err := template.ParseFiles("web/html/details.html", "web/html/header.html", "web/html/footer.html")
	if err != nil {
		panic(err)
	}
	url := episodate.GetApiShowUrl(query)
	jsonFromApi, err := episodate.GetShowData(url)
	if err != nil {
		fmt.Println("show not found")
	}

	var data Show
	json.Unmarshal([]byte(jsonFromApi), &data)
	isAdded := false
	if err := showlist.CheckIfShowInList(data.TvShow.Permalink); err != nil {
		isAdded = true
	}
	Desc := template.HTML(data.TvShow.Description)
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
		fmtEpisodeList(data.TvShow.Episodes),
	}

	detailsData := DetailsView{
		DetailsData: showData,
		Query:       query,
		IsAdded:     isAdded,
		Desc:        Desc,
	}
	err = t.Execute(w, detailsData)
	if err != nil {
		panic(err)
	}
}

// fmtEpisodeList changes some of the datas format before rendering
func fmtEpisodeList(episodes []episodesList) []episodesList {
	newList := []episodesList{}
	for _, val := range episodes {
		temp := val
		temp.AirDate = utils.FmtDate(val.AirDate)
		newList = append(newList, temp)
	}
	return newList
}
