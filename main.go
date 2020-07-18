package main

import (
	"flag"
	"fmt"

	"tvshowCalendar/calendar"
	"tvshowCalendar/episodate"
	"tvshowCalendar/showlist"
	"tvshowCalendar/web"
)

// cmd line - search show, add to list
// get api data
// update calendar
func main() {

	searchShowPtr := flag.String("s", "", "TV Show to search for")
	addShowPtr := flag.String("a", "", "Name of show to add")
	listShowsPtr := flag.Bool("l", false, "List of TV Shows to scan")
	showDetailsPtr := flag.String("d", "", "TV Show Details")
	updateCalendarPtr := flag.Bool("u", false, "Update calendar with TV show air dates")
	webServerPtr := flag.Bool("S", false, "Start web server")
	removeShowPtr := flag.String("r", "", "Remove show from list")
	lastestEpisodesPtr := flag.Bool("e", false, "Shows latest episode releases")

	flag.Parse()

	switch {
	case len(*searchShowPtr) > 0:
		searchForShow(*searchShowPtr)
	case len(*addShowPtr) > 0:
		addShow(*addShowPtr)
	case *listShowsPtr:
		listShows()
	case len(*showDetailsPtr) > 0:
		getShowDetails(*showDetailsPtr)
	case *updateCalendarPtr:
		updateCalendarWithShows()
	case *webServerPtr:
		startWebServer()
	case len(*removeShowPtr) > 0:
		removeShow(*removeShowPtr)
	case *lastestEpisodesPtr:
		latestEpisodes()
	default:
		fmt.Println("Help instructions here")
	}

}

func searchForShow(show string) {
	fmt.Printf("Searching for %s\n", show)
	episodate.SearchTvShow(show)
}

func addShow(show string) {
	result, err := showlist.AddTvShow(show)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}

func removeShow(show string) {
	result, err := showlist.RemoveShowFromFile(show)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func listShows() {
	showlist.ListShows()
}

func getShowDetails(show string) {
	err := episodate.ShowDetails(show)
	if err != nil {
		fmt.Println(err)
	}
}

func updateCalendarWithShows() {
	if err := calendar.UpdateCalendar(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ICS file has been updated")
	}
}

func startWebServer() {
	web.StartWebServer()
}

func latestEpisodes() {
	latest, err := calendar.GetLatestEpisodes()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(latest)
}
