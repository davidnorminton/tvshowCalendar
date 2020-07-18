package web

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
	"tvshowCalendar/calendar"
	"tvshowCalendar/showlist"
	"tvshowCalendar/utils"
)

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
	formatLatestList(latest)
	if err != nil {
		fmt.Println("Error getting latest episodes!")
	}

	data := IndexData{
		ShowList: showsInList,
		Latest:   latest,
	}
	err = t.Execute(w, data)
	if err != nil {
		panic(err)
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

func getEpisode(episode string) (string, string) {
	split := strings.Split(episode, ",")
	return split[0], strings.ReplaceAll(split[1], "released", "")
}

func createDateStr(date string) string {
	return date[:4] + "-" + date[4:6] + "-" + date[6:8]
}
