package main

import (
	"flag"
	"fmt"

	"tvshowCalendar/calendar"
	"tvshowCalendar/episodate"
	"tvshowCalendar/showlist"
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
	default:
		fmt.Println("Help instructions here\n")
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
	calendar.UpdateCalendar()
}
