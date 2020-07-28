package web

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/davidnorminton/tvshowCalendar/calendar"
	"github.com/davidnorminton/tvshowCalendar/showlist"
	"github.com/davidnorminton/tvshowCalendar/utils"
)

// IndexData is the structure of the data to be used on the index
type IndexData struct {
	ShowList []string
	Latest   []map[string]string
}

// IndexHandler handles the index page
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	t, err := template.ParseFiles("web/html/index.html", "web/html/header.html", "web/html/footer.html")
	if err != nil {
		panic(err)
	}

	showListFile, err := showlist.GetSavelistFileLocation()
	showsInList, err := readLines(showListFile)
	latest, err := calendar.GetLatestEpisodes()
	if err != nil {
		fmt.Println("Error getting latest episodes!")
	}

	formatLatestList(latest)

	data := IndexData{
		ShowList: showsInList,
		Latest:   latest,
	}
	if err = t.Execute(w, data); err != nil {
		fmt.Fprintf(w, "Error building page")
	}

}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// formatLatestList takes the current episodes list from the ics file and transforms
// it into a list suitable for the web interface
func formatLatestList(list []map[string]string) []map[string]string {
	newList := []map[string]string{}
	for _, val := range list {
		temp := val
		name, episode := getEpisode(val["Summary"])
		temp["Name"] = name
		temp["Summary"] = episode
		temp["Date"] = utils.FmtDate(createDateStr(val["Date"]))
		check := calendar.CheckAirDate(temp["Date"])
		if check == nil {
			newList = append(newList, temp)
		}
	}

	return newList
}

// getEpisode takes the summary from the ics event which is a string in the form
// show name, s1 e3 the show name and episode number are separated by a comma
func getEpisode(episode string) (string, string) {
	split := strings.Split(episode, ",")
	if len(split) > 1 {
		return split[0], split[1]
	}
	return "", episode
}

// createDateStr takes a date string in the format YMD ie 20201002
// and separates the digits wit a minus such as 2020-10-02
func createDateStr(date string) string {
	return date[:4] + "-" + date[4:6] + "-" + date[6:8]
}
