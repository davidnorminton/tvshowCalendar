package web

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"tvshowCalendar/episodate"
	"tvshowCalendar/showlist"
)

type SearchData struct {
	Name        string
	Reference   string
	StartDate   string
	Status      string
	StatusClass string
	Country     string
	Network     string
	IsAdded     bool
}

type SearchView struct {
	SearchData []SearchData
	Query      string
	Pages      []Pagination
}

type Pagination struct {
	Url    string
	Page   int
	Active bool
}

func getStatusClass(status string) string {
	return strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(status, " ", "-"), "/", "-"))
}

func isActive(i int, page int) bool {
	return i == page
}

// SearchHandler hadles the search results page
func SearchHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/html")
	query := r.URL.Query().Get("q")
	pageNumber := r.URL.Query().Get("page")
	if len(pageNumber) == 0 {
		pageNumber = "1"
	}
	t, err := template.ParseFiles("web/html/search.html", "web/html/header.html", "web/html/footer.html")
	if err != nil {
		panic(err)
	}

	showUrl := episodate.GetSearchUrl(query) + "&page=" + pageNumber
	html, _, _ := episodate.GetSearchData(showUrl, 0)
	search := episodate.SearchJson{}
	json.Unmarshal([]byte(html), &search)

	showData := []SearchData{}

	for i := 0; i < len(search.TvShows); i++ {
		isAdded := false
		if err := showlist.CheckIfShowInList(search.TvShows[i].Permalink); err != nil {
			isAdded = true
		}
		show := SearchData{
			search.TvShows[i].Name,
			search.TvShows[i].Permalink,
			search.TvShows[i].StartDate,
			search.TvShows[i].Status,
			getStatusClass(search.TvShows[i].Status),
			search.TvShows[i].Country,
			search.TvShows[i].Network,
			isAdded,
		}

		showData = append(showData, show)
	}

	pages := []Pagination{}
	if search.Pages > 1 {
		for i := 1; i <= search.Pages; i++ {
			numb, _ := strconv.Atoi(pageNumber)

			page := Pagination{
				Url:    "/search?q=" + episodate.FormatShowName(query) + "&page=" + strconv.Itoa(i),
				Page:   i,
				Active: isActive(i, numb),
			}
			pages = append(pages, page)
		}
	} else {
		pages = []Pagination{}
	}

	data := SearchView{
		SearchData: showData,
		Query:      query,
		Pages:      pages,
	}

	err = t.Execute(w, data)
	if err != nil {
		panic(err)
	}

}
